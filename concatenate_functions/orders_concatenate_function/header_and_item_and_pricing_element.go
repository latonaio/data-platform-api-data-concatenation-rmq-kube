package orders_concatenate_function

import (
	dpfm_api_input_reader "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_data_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Processing_Data_Formatter"

	"golang.org/x/xerrors"
)

func newHeaderAndItemAndScheduleLine(sdc dpfm_api_input_reader.OrdersSDC) ([]dpfm_api_processing_data_formatter.OrdersHeaderAndItemAndScheduleLine, error) {
	headerAndItemAndScheduleLines := make([]dpfm_api_processing_data_formatter.OrdersHeaderAndItemAndScheduleLine, 0)

	processingHeader, err := dpfm_api_processing_data_formatter.TypeConverter[dpfm_api_processing_data_formatter.OrdersHeader](sdc.DataConcatenation.Header)
	if err != nil {
		return []dpfm_api_processing_data_formatter.OrdersHeaderAndItemAndScheduleLine{}, xerrors.Errorf("TypeConverter Error: %w", err)
	}

	for _, item := range sdc.DataConcatenation.Item {
		orderID := item.OrderID
		orderItem := item.OrderItem
		if orderID != processingHeader.OrderID {
			continue
		}
		processingItem, err := dpfm_api_processing_data_formatter.TypeConverter[dpfm_api_processing_data_formatter.OrdersItem](item)
		if err != nil {
			return []dpfm_api_processing_data_formatter.OrdersHeaderAndItemAndScheduleLine{}, xerrors.Errorf("TypeConverter Error: %w", err)
		}

		processingItemScheduleLines, err := extractItemScheduleLine(orderID, orderItem, sdc.DataConcatenation.ItemScheduleLine)
		if err != nil {
			return []dpfm_api_processing_data_formatter.OrdersHeaderAndItemAndScheduleLine{}, xerrors.Errorf("extractItemScheduleLine Error: %w", err)
		}

		headerAndItemAndScheduleLines = append(headerAndItemAndScheduleLines, dpfm_api_processing_data_formatter.OrdersHeaderAndItemAndScheduleLine{
			Header:           processingHeader,
			Item:             processingItem,
			ItemScheduleLine: processingItemScheduleLines,
		})
	}

	return headerAndItemAndScheduleLines, err
}

func extractItemScheduleLine(orderID, orderItem int, itemScheduleLines []dpfm_api_input_reader.OrdersItemScheduleLine) ([]dpfm_api_processing_data_formatter.OrdersItemScheduleLine, error) {
	processingItemScheduleLines := make([]dpfm_api_processing_data_formatter.OrdersItemScheduleLine, 0)

	for _, itemScheduleLine := range itemScheduleLines {
		if itemScheduleLine.OrderID != orderID || itemScheduleLine.OrderItem != orderItem {
			continue
		}
		processingItemScheduleLine, err := dpfm_api_processing_data_formatter.TypeConverter[dpfm_api_processing_data_formatter.OrdersItemScheduleLine](itemScheduleLine)
		if err != nil {
			return []dpfm_api_processing_data_formatter.OrdersItemScheduleLine{}, err
		}

		processingItemScheduleLines = append(processingItemScheduleLines, processingItemScheduleLine)
	}

	return processingItemScheduleLines, nil
}
