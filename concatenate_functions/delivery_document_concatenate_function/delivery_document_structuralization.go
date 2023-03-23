package delivery_document_concatenate_function

import (
	dpfm_api_output_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Output_Formatter"
	dpfm_api_processing_data_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Processing_Data_Formatter"
	"strings"
)

func (c *DeliveryDocumentContraller) DeliveryDocumentStructuralization(deliveryDocumentSDC dpfm_api_processing_data_formatter.DeliveryDocumentSDC) (dpfm_api_output_formatter.DeliveryDocumentSDC, error) {
	output := dpfm_api_output_formatter.ConvertToDeliveryDocumentSDC(c.msg.Raw())

	// Header
	deliveryDocumentHeader, err := structuralizeHeader(deliveryDocumentSDC)
	if err != nil {
		return dpfm_api_output_formatter.DeliveryDocumentSDC{}, err
	}

	// Item
	deliveryDocumentItems, err := structuralizeItem(deliveryDocumentSDC)
	if err != nil {
		return dpfm_api_output_formatter.DeliveryDocumentSDC{}, err
	}

	// Partner
	deliveryDocumentPartners, err := structuralizePartner(deliveryDocumentSDC)
	if err != nil {
		return dpfm_api_output_formatter.DeliveryDocumentSDC{}, err
	}

	// Address
	deliveryDocumentAddresses, err := structuralizeAddress(deliveryDocumentSDC)
	if err != nil {
		return dpfm_api_output_formatter.DeliveryDocumentSDC{}, err
	}

	deliveryDocumentHeader.Item = deliveryDocumentItems
	deliveryDocumentHeader.Partner = deliveryDocumentPartners
	deliveryDocumentHeader.Address = deliveryDocumentAddresses

	output.Header = deliveryDocumentHeader
	output.ServiceLabel = "DELIVERY_DOCUMENT"

	switch strings.ToLower(output.APIType) {
	case "creates":
		output.APISchema = "DPFMDeliveryDocumentCreates"
	case "updates":
		output.APISchema = "DPFMDeliveryDocumentUpdates"
	case "function":
		output.APISchema = "DPFMDeliveryDocumentCreates"
	default:
		c.log.Error("unknown apitype %s", output.APIType)
		output.APIType = "creates"
		output.APISchema = "DPFMDeliveryDocumentCreates"
	}

	return output, nil
}

func structuralizeHeader(deliveryDocumentSDC dpfm_api_processing_data_formatter.DeliveryDocumentSDC) (dpfm_api_output_formatter.DeliveryDocumentHeader, error) {
	deliveryDocumentHeader, err := dpfm_api_output_formatter.TypeConverter[dpfm_api_output_formatter.DeliveryDocumentHeader](deliveryDocumentSDC.HeaderAndItem.Header)
	if err != nil {
		return dpfm_api_output_formatter.DeliveryDocumentHeader{}, err
	}

	return deliveryDocumentHeader, nil
}

func structuralizeItem(deliveryDocumentSDC dpfm_api_processing_data_formatter.DeliveryDocumentSDC) ([]dpfm_api_output_formatter.DeliveryDocumentItem, error) {
	deliveryDocumentItems := make([]dpfm_api_output_formatter.DeliveryDocumentItem, 0)
	for _, item := range deliveryDocumentSDC.HeaderAndItem.Item {
		deliveryDocumentItem, err := dpfm_api_output_formatter.TypeConverter[dpfm_api_output_formatter.DeliveryDocumentItem](item)
		if err != nil {
			return []dpfm_api_output_formatter.DeliveryDocumentItem{}, err
		}

		deliveryDocumentItems = append(deliveryDocumentItems, deliveryDocumentItem)
	}

	return deliveryDocumentItems, nil
}

func structuralizePartner(deliveryDocumentSDC dpfm_api_processing_data_formatter.DeliveryDocumentSDC) ([]dpfm_api_output_formatter.DeliveryDocumentPartner, error) {
	deliveryDocumentPartners := make([]dpfm_api_output_formatter.DeliveryDocumentPartner, 0)
	for _, partner := range deliveryDocumentSDC.HeaderAndPartner.Partner {
		deliveryDocumentPartner, err := dpfm_api_output_formatter.TypeConverter[dpfm_api_output_formatter.DeliveryDocumentPartner](partner)
		if err != nil {
			return []dpfm_api_output_formatter.DeliveryDocumentPartner{}, err
		}
		deliveryDocumentPartners = append(deliveryDocumentPartners, deliveryDocumentPartner)
	}

	return deliveryDocumentPartners, nil
}

func structuralizeAddress(deliveryDocumentSDC dpfm_api_processing_data_formatter.DeliveryDocumentSDC) ([]dpfm_api_output_formatter.DeliveryDocumentAddress, error) {
	deliveryDocumentAddresses := make([]dpfm_api_output_formatter.DeliveryDocumentAddress, 0)
	for _, address := range deliveryDocumentSDC.HeaderAndAddress.Address {
		deliveryDocumentAddress, err := dpfm_api_output_formatter.TypeConverter[dpfm_api_output_formatter.DeliveryDocumentAddress](address)
		if err != nil {
			return []dpfm_api_output_formatter.DeliveryDocumentAddress{}, err
		}
		deliveryDocumentAddresses = append(deliveryDocumentAddresses, deliveryDocumentAddress)
	}

	return deliveryDocumentAddresses, nil
}
