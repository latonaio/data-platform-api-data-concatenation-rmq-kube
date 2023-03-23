package delivery_notice_edi_for_voluntary_chain_smes_concatenate_function

import (
	dpfm_api_output_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Output_Formatter"
	dpfm_api_processing_data_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Processing_Data_Formatter"
	"strings"
)

func (c *DeliveryNoticeEDIForVoluntaryChainSMEsContraller) DeliveryNoticeEDIForVoluntaryChainSMEsStructuralization(deliveryNoticeEDIForVoluntaryChainSMEsSDC dpfm_api_processing_data_formatter.DeliveryNoticeEDIForVoluntaryChainSMEsSDC) (dpfm_api_output_formatter.DeliveryNoticeEDIForVoluntaryChainSMEsSDC, error) {
	output := dpfm_api_output_formatter.ConvertToDeliveryNoticeEDIForVoluntaryChainSMEsSDC(c.msg.Raw())

	// Header
	deliveryNoticeEDIForVoluntaryChainSMEsHeader, err := structuralizeHeader(deliveryNoticeEDIForVoluntaryChainSMEsSDC)
	if err != nil {
		return dpfm_api_output_formatter.DeliveryNoticeEDIForVoluntaryChainSMEsSDC{}, err
	}

	// Item
	deliveryNoticeEDIForVoluntaryChainSMEsItems, err := structuralizeItem(deliveryNoticeEDIForVoluntaryChainSMEsSDC)
	if err != nil {
		return dpfm_api_output_formatter.DeliveryNoticeEDIForVoluntaryChainSMEsSDC{}, err
	}

	deliveryNoticeEDIForVoluntaryChainSMEsHeader.Item = deliveryNoticeEDIForVoluntaryChainSMEsItems

	output.Header = deliveryNoticeEDIForVoluntaryChainSMEsHeader
	output.ServiceLabel = "FUNCTION_CONVERT_TO_DPFM_DELIVERY_DOCUMENT_FROM_DELIVERY_NOTICE_EDI_FOR_VOLUNTARY_CHAIN_SMES"

	switch strings.ToLower(output.APIType) {
	case "creates":
		output.APISchema = "DeliveryNoticeEDIForVoluntaryChainSMEsCreates"
	case "updates":
		output.APISchema = "DeliveryNoticeEDIForVoluntaryChainSMEsUpdates"
	case "function":
		output.APISchema = "DeliveryNoticeEDIForVoluntaryChainSMEsCreates"
	default:
		c.log.Error("unknown apitype %s", output.APIType)
		output.APIType = "creates"
		output.APISchema = "DeliveryNoticeEDIForVoluntaryChainSMEsCreates"
	}

	return output, nil
}

func structuralizeHeader(deliveryNoticeEDIForVoluntaryChainSMEsSDC dpfm_api_processing_data_formatter.DeliveryNoticeEDIForVoluntaryChainSMEsSDC) (dpfm_api_output_formatter.DeliveryNoticeEDIForVoluntaryChainSMEsHeader, error) {
	deliveryNoticeEDIForVoluntaryChainSMEsHeader, err := dpfm_api_output_formatter.TypeConverter[dpfm_api_output_formatter.DeliveryNoticeEDIForVoluntaryChainSMEsHeader](deliveryNoticeEDIForVoluntaryChainSMEsSDC.HeaderAndItem.Header)
	if err != nil {
		return dpfm_api_output_formatter.DeliveryNoticeEDIForVoluntaryChainSMEsHeader{}, err
	}

	return deliveryNoticeEDIForVoluntaryChainSMEsHeader, nil
}

func structuralizeItem(deliveryNoticeEDIForVoluntaryChainSMEsSDC dpfm_api_processing_data_formatter.DeliveryNoticeEDIForVoluntaryChainSMEsSDC) ([]dpfm_api_output_formatter.DeliveryNoticeEDIForVoluntaryChainSMEsItem, error) {
	deliveryNoticeEDIForVoluntaryChainSMEsItems := make([]dpfm_api_output_formatter.DeliveryNoticeEDIForVoluntaryChainSMEsItem, 0)
	for _, item := range deliveryNoticeEDIForVoluntaryChainSMEsSDC.HeaderAndItem.Item {
		deliveryNoticeEDIForVoluntaryChainSMEsItem, err := dpfm_api_output_formatter.TypeConverter[dpfm_api_output_formatter.DeliveryNoticeEDIForVoluntaryChainSMEsItem](item)
		if err != nil {
			return []dpfm_api_output_formatter.DeliveryNoticeEDIForVoluntaryChainSMEsItem{}, err
		}

		deliveryNoticeEDIForVoluntaryChainSMEsItems = append(deliveryNoticeEDIForVoluntaryChainSMEsItems, deliveryNoticeEDIForVoluntaryChainSMEsItem)
	}

	return deliveryNoticeEDIForVoluntaryChainSMEsItems, nil
}
