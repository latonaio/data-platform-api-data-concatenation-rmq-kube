package main

import (
	dpfm_api_input_reader "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-data-concatenation-rmq-kube/concatenate_functions/orders_concatenate_function"
	"data-platform-api-data-concatenation-rmq-kube/concatenate_functions/orders_edi_for_smes_concatenate_function"
	"data-platform-api-data-concatenation-rmq-kube/config"
	"encoding/json"
	"fmt"
	"time"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

func main() {
	l := logger.NewLogger()
	conf := config.NewConf()
	db, err := database.NewMySQL(conf.DB)
	if err != nil {
		l.Fatal(err.Error())
	}
	defer db.Close()

	rmq, err := rabbitmq.NewRabbitmqClient(conf.RMQ.URL(), conf.RMQ.QueueFrom(), "", nil, 0)
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Close()
	iter, err := rmq.Iterator()
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Stop()

	l.Info("READY!")
	for msg := range iter {
		go func(msg rabbitmq.RabbitmqMessage) {
			start := time.Now()
			err = callProcess(conf, rmq, msg, db)
			if err != nil {
				l.Error(err)
			}
			msg.Success()
			l.Info("process time %v\n", time.Since(start).Milliseconds())
		}(msg)
	}
}

func callProcess(conf *config.Conf, rmq *rabbitmq.RabbitmqClient, msg rabbitmq.RabbitmqMessage, db *database.Mysql) (err error) {
	l := logger.NewLogger()
	l.AddHeaderInfo(map[string]interface{}{"runtime_session_id": getSessionID(msg.Data())})
	defer recovery(l, &err)
	input := map[string]interface{}{}
	err = json.Unmarshal(msg.Raw(), &input)
	if err != nil {
		return err
	}

	if !input["result"].(bool) {
		rmq.Send(conf.RMQ.QueueToErrResponse(), input)
		return xerrors.New("result is false")
	}

	serviceLabel, ok := input["service_label"].(string)
	if !ok {
		return xerrors.Errorf("service_label is not string")
	}
	concatenateMapper, err := getConcatenateMapper(serviceLabel, db)
	if err != nil {
		return err
	}

	var output interface{}
	switch serviceLabel {
	case "FUNCTION_ORDERS_DATA_CONCATENATION":
		c := orders_concatenate_function.NewOrdersContraller(conf, rmq, msg, l)
		output, err = c.OrdersProcess(concatenateMapper)
	case "FUNCTION_ORDERS_EDI_FOR_SMES_DATA_CONCATENATION":
		c := orders_edi_for_smes_concatenate_function.NewOrdersEDIForSMEsContraller(conf, rmq, msg, l)
		output, err = c.OrdersEDIForSMEsProcess(concatenateMapper)
	default:
		l.Info("Unknown service_label %v", input["service_label"])
	}
	if err != nil {
		input["result"] = false
		input["api_processing_result"] = false
		input["api_processing_error"] = err.Error()
		rmq.Send(conf.RMQ.QueueToErrResponse(), input)
		return err
	}

	l.JsonParseOut(output)

	return err
}

func getConcatenateMapper(serviceLabel string, db *database.Mysql) ([]dpfm_api_input_reader.ConcatenateMapper, error) {
	rows, err := db.Query(
		`SELECT *
		FROM DataPlatformCommonSettingsMysqlKube.data_platform_data_concatenate_mapper_data
		WHERE ServiceLabel = ?;`, serviceLabel,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res := make([]dpfm_api_input_reader.ConcatenateMapper, 0)

	i := 0
	for rows.Next() {
		i++
		data := dpfm_api_input_reader.ConcatenateMapper{}

		err := rows.Scan(
			&data.ConcatenateMapperID,
			&data.ServiceLabel,
			&data.BaseAPIName,
			&data.ConcatenateAPIName1,
			&data.ConcatenateAPIName2,
			&data.ConcatenateAPIName3,
			&data.ConcatenateAPIName4,
			&data.ConcatenateAPIName5,
			&data.ConcatenateKey1,
			&data.ConcatenateKey2,
			&data.ConcatenateKey3,
			&data.ConcatenateKey4,
			&data.ConcatenateKey5,
		)
		if err != nil {
			return nil, err
		}

		res = append(res, data)
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_concatenate_mapper_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func getSessionID(data map[string]interface{}) string {
	id := fmt.Sprintf("%v", data["runtime_session_id"])
	return id
}

func recovery(l *logger.Logger, err *error) {
	if e := recover(); e != nil {
		*err = xerrors.Errorf("error occurred: %w", e)
		l.Error("%+v", *err)
		return
	}
}
