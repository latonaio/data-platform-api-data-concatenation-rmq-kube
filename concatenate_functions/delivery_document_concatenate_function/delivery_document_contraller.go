package delivery_document_concatenate_function

import (
	dpfm_api_input_reader "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-data-concatenation-rmq-kube/config"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type DeliveryDocumentContraller struct {
	conf *config.Conf
	rmq  *rabbitmq.RabbitmqClient
	msg  rabbitmq.RabbitmqMessage
	log  *logger.Logger
}

func NewDeliveryDocumentContraller(
	conf *config.Conf,
	rmq *rabbitmq.RabbitmqClient,
	msg rabbitmq.RabbitmqMessage,
	log *logger.Logger,
) *DeliveryDocumentContraller {
	return &DeliveryDocumentContraller{
		conf: conf,
		rmq:  rmq,
		msg:  msg,
		log:  log,
	}
}

func (c *DeliveryDocumentContraller) DeliveryDocumentProcess(concatenateMapper []dpfm_api_input_reader.ConcatenateMapper, queueName string) (dpfm_api_output_formatter.DeliveryDocumentSDC, error) {
	deliveryDocumentConcatenated, err := c.DeliveryDocumentConcatenation(concatenateMapper)
	if err != nil {
		return dpfm_api_output_formatter.DeliveryDocumentSDC{}, xerrors.Errorf("Concatenate Error: %w", err)
	}

	deliveryDocumentOutput, err := c.DeliveryDocumentStructuralization(deliveryDocumentConcatenated)
	if err != nil {
		return dpfm_api_output_formatter.DeliveryDocumentSDC{}, xerrors.Errorf("Structuralize Error: %w", err)
	}

	c.rmq.Send(queueName, deliveryDocumentOutput)

	return deliveryDocumentOutput, nil
}
