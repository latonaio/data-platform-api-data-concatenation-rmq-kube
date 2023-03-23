package delivery_notice_edi_for_smes_concatenate_function

import (
	dpfm_api_output_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Output_Formatter"
	dpfm_api_processing_data_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Processing_Data_Formatter"
	"strings"
)

func (c *DeliveryNoticeEDIForSMEsContraller) DeliveryNoticeEDIForSMEsStructuralization(deliveryNoticeEDIForSMEsSDC dpfm_api_processing_data_formatter.DeliveryNoticeEDIForSMEsSDC) (dpfm_api_output_formatter.DeliveryNoticeEDIForSMEsSDC, error) {
	output := dpfm_api_output_formatter.ConvertToDeliveryNoticeEDIForSMEsSDC(c.msg.Raw())

	// Header
	deliveryNoticeEDIForSMEsHeader, err := structuralizeHeader(deliveryNoticeEDIForSMEsSDC)
	if err != nil {
		return dpfm_api_output_formatter.DeliveryNoticeEDIForSMEsSDC{}, err
	}

	// Item
	deliveryNoticeEDIForSMEsItems, err := structuralizeItem(deliveryNoticeEDIForSMEsSDC)
	if err != nil {
		return dpfm_api_output_formatter.DeliveryNoticeEDIForSMEsSDC{}, err
	}

	deliveryNoticeEDIForSMEsHeader.Item = deliveryNoticeEDIForSMEsItems

	output.Header = deliveryNoticeEDIForSMEsHeader
	output.ServiceLabel = "FUNCTION_CONVERT_TO_DPFM_DELIVERY_DOCUMENT_FROM_DELIVERY_NOTICE_EDI_FOR_SMES"

	switch strings.ToLower(output.APIType) {
	case "creates":
		output.APISchema = "DeliveryNoticeEDIForSMEsCreates"
	case "updates":
		output.APISchema = "DeliveryNoticeEDIForSMEsUpdates"
	case "function":
		output.APISchema = "DeliveryNoticeEDIForSMEsCreates"
	default:
		c.log.Error("unknown apitype %s", output.APIType)
		output.APIType = "creates"
		output.APISchema = "DeliveryNoticeEDIForSMEsCreates"
	}

	return output, nil
}

func structuralizeHeader(deliveryNoticeEDIForSMEsSDC dpfm_api_processing_data_formatter.DeliveryNoticeEDIForSMEsSDC) (dpfm_api_output_formatter.DeliveryNoticeEDIForSMEsHeader, error) {
	deliveryNoticeEDIForSMEsHeader, err := dpfm_api_output_formatter.TypeConverter[dpfm_api_output_formatter.DeliveryNoticeEDIForSMEsHeader](deliveryNoticeEDIForSMEsSDC.HeaderAndItem.Header)
	if err != nil {
		return dpfm_api_output_formatter.DeliveryNoticeEDIForSMEsHeader{}, err
	}

	return deliveryNoticeEDIForSMEsHeader, nil
}

func structuralizeItem(deliveryNoticeEDIForSMEsSDC dpfm_api_processing_data_formatter.DeliveryNoticeEDIForSMEsSDC) ([]dpfm_api_output_formatter.DeliveryNoticeEDIForSMEsItem, error) {
	deliveryNoticeEDIForSMEsItems := make([]dpfm_api_output_formatter.DeliveryNoticeEDIForSMEsItem, 0)
	for _, item := range deliveryNoticeEDIForSMEsSDC.HeaderAndItem.Item {
		deliveryNoticeEDIForSMEsItem, err := dpfm_api_output_formatter.TypeConverter[dpfm_api_output_formatter.DeliveryNoticeEDIForSMEsItem](item)
		if err != nil {
			return []dpfm_api_output_formatter.DeliveryNoticeEDIForSMEsItem{}, err
		}

		deliveryNoticeEDIForSMEsItems = append(deliveryNoticeEDIForSMEsItems, deliveryNoticeEDIForSMEsItem)
	}

	return deliveryNoticeEDIForSMEsItems, nil
}
