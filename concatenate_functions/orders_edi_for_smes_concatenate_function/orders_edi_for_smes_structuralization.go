package orders_edi_for_smes_concatenate_function

import (
	dpfm_api_output_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Output_Formatter"
	dpfm_api_processing_data_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Processing_Data_Formatter"
)

func (c *OrdersEDIForSMEsContraller) OrdersEDIForSMEsStructuralization(ordersEDIForSMEsSDC dpfm_api_processing_data_formatter.OrdersEDIForSMEsSDC) (dpfm_api_output_formatter.OrdersEDIForSMEsSDC, error) {
	output := dpfm_api_output_formatter.ConvertToOrdersEDIForSMEsSDC(c.msg.Raw())

	// Header
	ordersEDIForSMEsHeader, err := structuralizeHeader(ordersEDIForSMEsSDC)
	if err != nil {
		return dpfm_api_output_formatter.OrdersEDIForSMEsSDC{}, err
	}

	// Item
	ordersEDIForSMEsItems, err := structuralizeItem(ordersEDIForSMEsSDC)
	if err != nil {
		return dpfm_api_output_formatter.OrdersEDIForSMEsSDC{}, err
	}

	ordersEDIForSMEsHeader.Item = ordersEDIForSMEsItems

	output.Header = ordersEDIForSMEsHeader
	output.ServiceLabel = "FUNCTION_CONVERT_TO_DPFM_ORDERS_FROM_ORDERS_EDI_FOR_SMES"
	output.APIType = "function"
	output.APISchema = "OrdersEDIForSMEsCreates"

	return output, nil
}

func structuralizeHeader(ordersEDIForSMEsSDC dpfm_api_processing_data_formatter.OrdersEDIForSMEsSDC) (dpfm_api_output_formatter.OrdersEDIForSMEsHeader, error) {
	ordersEDIForSMEsHeader, err := dpfm_api_output_formatter.TypeConverter[dpfm_api_output_formatter.OrdersEDIForSMEsHeader](ordersEDIForSMEsSDC.HeaderAndItem.Header)
	if err != nil {
		return dpfm_api_output_formatter.OrdersEDIForSMEsHeader{}, err
	}

	return ordersEDIForSMEsHeader, nil
}

func structuralizeItem(ordersEDIForSMEsSDC dpfm_api_processing_data_formatter.OrdersEDIForSMEsSDC) ([]dpfm_api_output_formatter.OrdersEDIForSMEsItem, error) {
	ordersEDIForSMEsItems := make([]dpfm_api_output_formatter.OrdersEDIForSMEsItem, 0)
	for _, item := range ordersEDIForSMEsSDC.HeaderAndItem.Item {
		ordersEDIForSMEsItem, err := dpfm_api_output_formatter.TypeConverter[dpfm_api_output_formatter.OrdersEDIForSMEsItem](item)
		if err != nil {
			return []dpfm_api_output_formatter.OrdersEDIForSMEsItem{}, err
		}

		ordersEDIForSMEsItems = append(ordersEDIForSMEsItems, ordersEDIForSMEsItem)
	}

	return ordersEDIForSMEsItems, nil
}
