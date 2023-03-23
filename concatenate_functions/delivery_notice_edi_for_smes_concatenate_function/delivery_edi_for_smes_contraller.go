package delivery_notice_edi_for_smes_concatenate_function

import (
	dpfm_api_input_reader "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-data-concatenation-rmq-kube/config"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type DeliveryNoticeEDIForSMEsContraller struct {
	conf *config.Conf
	rmq  *rabbitmq.RabbitmqClient
	msg  rabbitmq.RabbitmqMessage
	log  *logger.Logger
}

func NewDeliveryNoticeEDIForSMEsContraller(
	conf *config.Conf,
	rmq *rabbitmq.RabbitmqClient,
	msg rabbitmq.RabbitmqMessage,
	log *logger.Logger,
) *DeliveryNoticeEDIForSMEsContraller {
	return &DeliveryNoticeEDIForSMEsContraller{
		conf: conf,
		rmq:  rmq,
		msg:  msg,
		log:  log,
	}
}

func (c *DeliveryNoticeEDIForSMEsContraller) DeliveryNoticeEDIForSMEsProcess(concatenateMapper []dpfm_api_input_reader.ConcatenateMapper, queueName string) (dpfm_api_output_formatter.DeliveryNoticeEDIForSMEsSDC, error) {
	deliveryNoticeEDIForSMEsConcatenated, err := c.DeliveryNoticeEDIForSMEsConcatenation(concatenateMapper)
	if err != nil {
		return dpfm_api_output_formatter.DeliveryNoticeEDIForSMEsSDC{}, xerrors.Errorf("Concatenate Error: %w", err)
	}

	deliveryNoticeEDIForSMEsOutput, err := c.DeliveryNoticeEDIForSMEsStructuralization(deliveryNoticeEDIForSMEsConcatenated)
	if err != nil {
		return dpfm_api_output_formatter.DeliveryNoticeEDIForSMEsSDC{}, xerrors.Errorf("Structuralize Error: %w", err)
	}

	c.rmq.Send(queueName, deliveryNoticeEDIForSMEsOutput)

	return deliveryNoticeEDIForSMEsOutput, nil
}
