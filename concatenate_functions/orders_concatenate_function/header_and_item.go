package orders_concatenate_function

import (
	dpfm_api_input_reader "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_data_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Processing_Data_Formatter"

	"golang.org/x/xerrors"
)

func newHeaderAndItem(sdc dpfm_api_input_reader.OrdersSDC) (dpfm_api_processing_data_formatter.OrdersHeaderAndItem, error) {
	processingHeader, err := dpfm_api_processing_data_formatter.TypeConverter[dpfm_api_processing_data_formatter.OrdersHeader](sdc.DataConcatenation.Header)
	if err != nil {
		return dpfm_api_processing_data_formatter.OrdersHeaderAndItem{}, xerrors.Errorf("TypeConverter Error: %w", err)
	}

	orderID := processingHeader.OrderID
	processingItems, err := extractItem(orderID, sdc.DataConcatenation.Item)
	if err != nil {
		return dpfm_api_processing_data_formatter.OrdersHeaderAndItem{}, xerrors.Errorf("extractItem Error: %w", err)
	}

	headerAndItem := dpfm_api_processing_data_formatter.OrdersHeaderAndItem{
		Header: processingHeader,
		Item:   processingItems,
	}

	return headerAndItem, err
}

func extractItem(orderID int, items []dpfm_api_input_reader.OrdersItem) ([]dpfm_api_processing_data_formatter.OrdersItem, error) {
	processingItems := make([]dpfm_api_processing_data_formatter.OrdersItem, 0)

	for _, item := range items {
		if item.OrderID != orderID {
			continue
		}
		processingItem, err := dpfm_api_processing_data_formatter.TypeConverter[dpfm_api_processing_data_formatter.OrdersItem](item)
		if err != nil {
			return []dpfm_api_processing_data_formatter.OrdersItem{}, xerrors.Errorf("TypeConverter Error: %w", err)
		}
		processingItems = append(processingItems, processingItem)
	}

	return processingItems, nil
}
