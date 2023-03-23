package orders_edi_for_smes_concatenate_function

import (
	dpfm_api_input_reader "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-data-concatenation-rmq-kube/config"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type OrdersEDIForSMEsContraller struct {
	conf *config.Conf
	rmq  *rabbitmq.RabbitmqClient
	msg  rabbitmq.RabbitmqMessage
	log  *logger.Logger
}

func NewOrdersEDIForSMEsContraller(
	conf *config.Conf,
	rmq *rabbitmq.RabbitmqClient,
	msg rabbitmq.RabbitmqMessage,
	log *logger.Logger,
) *OrdersEDIForSMEsContraller {
	return &OrdersEDIForSMEsContraller{
		conf: conf,
		rmq:  rmq,
		msg:  msg,
		log:  log,
	}
}

func (c *OrdersEDIForSMEsContraller) OrdersEDIForSMEsProcess(concatenateMapper []dpfm_api_input_reader.ConcatenateMapper, queueName string) (dpfm_api_output_formatter.OrdersEDIForSMEsSDC, error) {
	ordersEDIForSMEsConcatenated, err := c.OrdersEDIForSMEsConcatenation(concatenateMapper)
	if err != nil {
		return dpfm_api_output_formatter.OrdersEDIForSMEsSDC{}, xerrors.Errorf("Concatenate Error: %w", err)
	}

	ordersEDIForSMEsOutput, err := c.OrdersEDIForSMEsStructuralization(ordersEDIForSMEsConcatenated)
	if err != nil {
		return dpfm_api_output_formatter.OrdersEDIForSMEsSDC{}, xerrors.Errorf("Structuralize Error: %w", err)
	}

	c.rmq.Send(queueName, ordersEDIForSMEsOutput)

	return ordersEDIForSMEsOutput, nil
}
