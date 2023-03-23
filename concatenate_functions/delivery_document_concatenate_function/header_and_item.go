package delivery_document_concatenate_function

import (
	dpfm_api_input_reader "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_data_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Processing_Data_Formatter"

	"golang.org/x/xerrors"
)

func newHeaderAndItem(sdc dpfm_api_input_reader.DeliveryDocumentSDC) (dpfm_api_processing_data_formatter.DeliveryDocumentHeaderAndItem, error) {
	processingHeader, err := dpfm_api_processing_data_formatter.TypeConverter[dpfm_api_processing_data_formatter.DeliveryDocumentHeader](sdc.DataConcatenation.Header)
	if err != nil {
		return dpfm_api_processing_data_formatter.DeliveryDocumentHeaderAndItem{}, xerrors.Errorf("TypeConverter Error: %w", err)
	}

	deliveryDocument := processingHeader.DeliveryDocument
	processingItems, err := extractItem(deliveryDocument, sdc.DataConcatenation.Item)
	if err != nil {
		return dpfm_api_processing_data_formatter.DeliveryDocumentHeaderAndItem{}, xerrors.Errorf("extractItem Error: %w", err)
	}

	headerAndItem := dpfm_api_processing_data_formatter.DeliveryDocumentHeaderAndItem{
		Header: processingHeader,
		Item:   processingItems,
	}

	return headerAndItem, err
}

func extractItem(deliveryDocument int, items []dpfm_api_input_reader.DeliveryDocumentItem) ([]dpfm_api_processing_data_formatter.DeliveryDocumentItem, error) {
	processingItems := make([]dpfm_api_processing_data_formatter.DeliveryDocumentItem, 0)

	for _, item := range items {
		if item.DeliveryDocument != deliveryDocument {
			continue
		}
		processingItem, err := dpfm_api_processing_data_formatter.TypeConverter[dpfm_api_processing_data_formatter.DeliveryDocumentItem](item)
		if err != nil {
			return []dpfm_api_processing_data_formatter.DeliveryDocumentItem{}, xerrors.Errorf("TypeConverter Error: %w", err)
		}
		processingItems = append(processingItems, processingItem)
	}

	return processingItems, nil
}
