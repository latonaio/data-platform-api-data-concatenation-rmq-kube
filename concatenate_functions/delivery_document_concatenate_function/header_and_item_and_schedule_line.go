package delivery_document_concatenate_function

import (
	dpfm_api_input_reader "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_data_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Processing_Data_Formatter"

	"golang.org/x/xerrors"
)

func newHeaderAndItemAndPricingElement(sdc dpfm_api_input_reader.OrdersSDC) ([]dpfm_api_processing_data_formatter.OrdersHeaderAndItemAndPricingElement, error) {
	headerAndItemAndPricingElements := make([]dpfm_api_processing_data_formatter.OrdersHeaderAndItemAndPricingElement, 0)

	processingHeader, err := dpfm_api_processing_data_formatter.TypeConverter[dpfm_api_processing_data_formatter.OrdersHeader](sdc.DataConcatenation.Header)
	if err != nil {
		return []dpfm_api_processing_data_formatter.OrdersHeaderAndItemAndPricingElement{}, xerrors.Errorf("TypeConverter Error: %w", err)
	}

	for _, item := range sdc.DataConcatenation.Item {
		orderID := item.OrderID
		orderItem := item.OrderItem
		if orderID != processingHeader.OrderID {
			continue
		}
		processingItem, err := dpfm_api_processing_data_formatter.TypeConverter[dpfm_api_processing_data_formatter.OrdersItem](item)
		if err != nil {
			return []dpfm_api_processing_data_formatter.OrdersHeaderAndItemAndPricingElement{}, xerrors.Errorf("TypeConverter Error: %w", err)
		}

		processingItemPricingElements, err := extractItemPricingElement(orderID, orderItem, sdc.DataConcatenation.ItemPricingElement)
		if err != nil {
			return []dpfm_api_processing_data_formatter.OrdersHeaderAndItemAndPricingElement{}, xerrors.Errorf("extractItemPricingElement Error: %w", err)
		}

		headerAndItemAndPricingElements = append(headerAndItemAndPricingElements, dpfm_api_processing_data_formatter.OrdersHeaderAndItemAndPricingElement{
			Header:             processingHeader,
			Item:               processingItem,
			ItemPricingElement: processingItemPricingElements,
		})
	}

	return headerAndItemAndPricingElements, err
}

func extractItemPricingElement(orderID, orderItem int, itemPricingElements []dpfm_api_input_reader.OrdersItemPricingElement) ([]dpfm_api_processing_data_formatter.OrdersItemPricingElement, error) {
	processingItemPricingElements := make([]dpfm_api_processing_data_formatter.OrdersItemPricingElement, 0)

	for _, itemPricingElement := range itemPricingElements {
		if itemPricingElement.OrderID != orderID || itemPricingElement.OrderItem != orderItem {
			continue
		}
		processingItemPricingElement, err := dpfm_api_processing_data_formatter.TypeConverter[dpfm_api_processing_data_formatter.OrdersItemPricingElement](itemPricingElement)
		if err != nil {
			return []dpfm_api_processing_data_formatter.OrdersItemPricingElement{}, err
		}

		processingItemPricingElements = append(processingItemPricingElements, processingItemPricingElement)
	}

	return processingItemPricingElements, nil
}
