package orders_edi_for_voluntary_chain_smes_concatenate_function

import (
	dpfm_api_input_reader "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-data-concatenation-rmq-kube/config"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type OrdersEDIForVoluntaryChainSMEsContraller struct {
	conf *config.Conf
	rmq  *rabbitmq.RabbitmqClient
	msg  rabbitmq.RabbitmqMessage
	log  *logger.Logger
}

func NewOrdersEDIForVoluntaryChainSMEsContraller(
	conf *config.Conf,
	rmq *rabbitmq.RabbitmqClient,
	msg rabbitmq.RabbitmqMessage,
	log *logger.Logger,
) *OrdersEDIForVoluntaryChainSMEsContraller {
	return &OrdersEDIForVoluntaryChainSMEsContraller{
		conf: conf,
		rmq:  rmq,
		msg:  msg,
		log:  log,
	}
}

func (c *OrdersEDIForVoluntaryChainSMEsContraller) OrdersEDIForVoluntaryChainSMEsProcess(concatenateMapper []dpfm_api_input_reader.ConcatenateMapper, queueName string) (dpfm_api_output_formatter.OrdersEDIForVoluntaryChainSMEsSDC, error) {
	ordersEDIForVoluntaryChainSMEsConcatenated, err := c.OrdersEDIForVoluntaryChainSMEsConcatenation(concatenateMapper)
	if err != nil {
		return dpfm_api_output_formatter.OrdersEDIForVoluntaryChainSMEsSDC{}, xerrors.Errorf("Concatenate Error: %w", err)
	}

	ordersEDIForVoluntaryChainSMEsOutput, err := c.OrdersEDIForVoluntaryChainSMEsStructuralization(ordersEDIForVoluntaryChainSMEsConcatenated)
	if err != nil {
		return dpfm_api_output_formatter.OrdersEDIForVoluntaryChainSMEsSDC{}, xerrors.Errorf("Structuralize Error: %w", err)
	}

	c.rmq.Send(queueName, ordersEDIForVoluntaryChainSMEsOutput)

	return ordersEDIForVoluntaryChainSMEsOutput, nil
}
