package delivery_document_concatenate_function

import (
	dpfm_api_input_reader "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_data_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Processing_Data_Formatter"

	"golang.org/x/xerrors"
)

func newHeaderAndAddress(sdc dpfm_api_input_reader.DeliveryDocumentSDC) (dpfm_api_processing_data_formatter.DeliveryDocumentHeaderAndAddress, error) {
	processingHeader, err := dpfm_api_processing_data_formatter.TypeConverter[dpfm_api_processing_data_formatter.DeliveryDocumentHeader](sdc.DataConcatenation.Header)
	if err != nil {
		return dpfm_api_processing_data_formatter.DeliveryDocumentHeaderAndAddress{}, xerrors.Errorf("TypeConverter Error: %w", err)
	}

	deliveryDocument := processingHeader.DeliveryDocument
	processingAddressses, err := extractAddress(deliveryDocument, sdc.DataConcatenation.Address)
	if err != nil {
		return dpfm_api_processing_data_formatter.DeliveryDocumentHeaderAndAddress{}, xerrors.Errorf("extractAddress Error: %w", err)
	}

	headerAndAddress := dpfm_api_processing_data_formatter.DeliveryDocumentHeaderAndAddress{
		Header:  processingHeader,
		Address: processingAddressses,
	}

	return headerAndAddress, err
}

func extractAddress(deliveryDocument int, addresses []dpfm_api_input_reader.DeliveryDocumentAddress) ([]dpfm_api_processing_data_formatter.DeliveryDocumentAddress, error) {
	processingAddresses := make([]dpfm_api_processing_data_formatter.DeliveryDocumentAddress, 0)

	for _, address := range addresses {
		if address.DeliveryDocument != deliveryDocument {
			continue
		}
		processingAddress, err := dpfm_api_processing_data_formatter.TypeConverter[dpfm_api_processing_data_formatter.DeliveryDocumentAddress](address)
		if err != nil {
			return []dpfm_api_processing_data_formatter.DeliveryDocumentAddress{}, xerrors.Errorf("TypeConverter Error: %w", err)
		}
		processingAddresses = append(processingAddresses, processingAddress)
	}

	return processingAddresses, nil
}
