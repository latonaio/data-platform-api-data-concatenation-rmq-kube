package main

import (
	"context"
	dpfm_api_input_reader "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Input_Reader"
	api_processing_data_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Processing_Data_Formatter"
	"data-platform-api-data-concatenation-rmq-kube/concatenate_functions"
	"data-platform-api-data-concatenation-rmq-kube/config"
	"encoding/json"
	"fmt"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

func main() {
	ctx := context.Background()
	l := logger.NewLogger()
	conf := config.NewConf()
	db, err := database.NewMySQL(conf.DB)
	if err != nil {
		l.Fatal(err.Error())
	}
	defer db.Close()

	_ = ctx
	_ = db
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
			err = callProcess(msg, conf, db)
			if err != nil {
				l.Error(err)
			}
			msg.Success()
		}(msg)
	}
}

func callProcess(msg rabbitmq.RabbitmqMessage, conf *config.Conf, db *database.Mysql) (err error) {
	l := logger.NewLogger()
	l.AddHeaderInfo(map[string]interface{}{"runtime_session_id": getSessionID(msg.Data())})
	defer recovery(l, &err)
	input := map[string]interface{}{}
	err = json.Unmarshal(msg.Raw(), &input)
	if err != nil {
		return err
	}

	serviceLabel, ok := input["service_label"].(string)
	if !ok {
		return xerrors.Errorf("service_label is not string")
	}
	concatenateMapper, err := getConcatenateMapper(serviceLabel, db)
	if err != nil {
		return err
	}

	ordersSDC := api_processing_data_formatter.OrdersSDC{}
	switch serviceLabel {
	case "ORDERS":
		ordersSDC, err = concatenate_functions.OrdersConcatenate(msg, concatenateMapper)
		if err != nil {
			return err
		}
	default:
		l.Info("Unknown service_label %v", input["service_label"])
	}

	l.Info(ordersSDC)

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
