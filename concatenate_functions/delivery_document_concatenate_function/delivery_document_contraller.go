package delivery_document_concatenate_function

import (
	dpfm_api_input_reader "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-data-concatenation-rmq-kube/config"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type OrdersContraller struct {
	conf *config.Conf
	rmq  *rabbitmq.RabbitmqClient
	msg  rabbitmq.RabbitmqMessage
	log  *logger.Logger
}

func NewOrdersContraller(
	conf *config.Conf,
	rmq *rabbitmq.RabbitmqClient,
	msg rabbitmq.RabbitmqMessage,
	log *logger.Logger,
) *OrdersContraller {
	return &OrdersContraller{
		conf: conf,
		rmq:  rmq,
		msg:  msg,
		log:  log,
	}
}

func (c *OrdersContraller) OrdersProcess(concatenateMapper []dpfm_api_input_reader.ConcatenateMapper) (dpfm_api_output_formatter.OrdersSDC, error) {
	ordersConcatenated, err := c.OrdersConcatenation(concatenateMapper)
	if err != nil {
		return dpfm_api_output_formatter.OrdersSDC{}, xerrors.Errorf("Concatenate Error: %w", err)
	}

	ordersOutput, err := c.OrdersStructuralization(ordersConcatenated)
	if err != nil {
		return dpfm_api_output_formatter.OrdersSDC{}, xerrors.Errorf("Structuralize Error: %w", err)
	}

	c.rmq.Send(c.conf.RMQ.QueueTo()[0], ordersOutput)

	return ordersOutput, nil
}
