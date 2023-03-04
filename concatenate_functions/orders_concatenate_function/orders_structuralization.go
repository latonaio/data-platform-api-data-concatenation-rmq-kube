package orders_concatenate_function

import (
	dpfm_api_output_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Output_Formatter"
	dpfm_api_processing_data_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Processing_Data_Formatter"
)

func (c *OrdersContraller) OrdersStructuralization(ordersSDC dpfm_api_processing_data_formatter.OrdersSDC) (dpfm_api_output_formatter.OrdersSDC, error) {
	output := dpfm_api_output_formatter.ConvertToOrdersSDC(c.msg.Raw())

	// Header
	ordersHeader, err := structuralizeHeader(ordersSDC)
	if err != nil {
		return dpfm_api_output_formatter.OrdersSDC{}, err
	}

	// Item
	ordersItems, err := structuralizeItem(ordersSDC)
	if err != nil {
		return dpfm_api_output_formatter.OrdersSDC{}, err
	}

	// Partner
	ordersPartners, err := structuralizePartner(ordersSDC)
	if err != nil {
		return dpfm_api_output_formatter.OrdersSDC{}, err
	}

	// Address
	ordersAddresses, err := structuralizeAddress(ordersSDC)
	if err != nil {
		return dpfm_api_output_formatter.OrdersSDC{}, err
	}

	ordersHeader.Item = ordersItems
	ordersHeader.Partner = ordersPartners
	ordersHeader.Address = ordersAddresses

	output.Header = ordersHeader
	output.ServiceLabel = "ORDERS"
	output.APIType = "creates"
	output.APISchema = "DPFMOrdersCreates"

	return output, nil
}

func structuralizeHeader(ordersSDC dpfm_api_processing_data_formatter.OrdersSDC) (dpfm_api_output_formatter.OrdersHeader, error) {
	ordersHeader, err := dpfm_api_output_formatter.TypeConverter[dpfm_api_output_formatter.OrdersHeader](ordersSDC.HeaderAndItem.Header)
	if err != nil {
		return dpfm_api_output_formatter.OrdersHeader{}, err
	}

	return ordersHeader, nil
}

func structuralizeItem(ordersSDC dpfm_api_processing_data_formatter.OrdersSDC) ([]dpfm_api_output_formatter.OrdersItem, error) {
	ordersItems := make([]dpfm_api_output_formatter.OrdersItem, 0)
	for i, item := range ordersSDC.HeaderAndItem.Item {
		ordersItem, err := dpfm_api_output_formatter.TypeConverter[dpfm_api_output_formatter.OrdersItem](item)
		if err != nil {
			return []dpfm_api_output_formatter.OrdersItem{}, err
		}

		// ItemPricingElement
		ordersItemPricingElements, err := structuralizeItemPricingElemnt(ordersSDC, i)
		if err != nil {
			return []dpfm_api_output_formatter.OrdersItem{}, err
		}

		// ItemScheduleLine
		ordersItemScheduleLines, err := structuralizeItemScheduleLine(ordersSDC, i)
		if err != nil {
			return []dpfm_api_output_formatter.OrdersItem{}, err
		}

		ordersItem.ItemPricingElement = ordersItemPricingElements
		ordersItem.ItemScheduleLine = ordersItemScheduleLines
		ordersItems = append(ordersItems, ordersItem)
	}

	return ordersItems, nil
}

func structuralizeItemPricingElemnt(ordersSDC dpfm_api_processing_data_formatter.OrdersSDC, i int) ([]dpfm_api_output_formatter.OrdersItemPricingElement, error) {
	ordersItemPricingElements := make([]dpfm_api_output_formatter.OrdersItemPricingElement, 0)
	for _, itemPricingElement := range ordersSDC.HeaderAndItemAndPricingElement[i].ItemPricingElement {
		ordersItemPricingElement, err := dpfm_api_output_formatter.TypeConverter[dpfm_api_output_formatter.OrdersItemPricingElement](itemPricingElement)
		if err != nil {
			return []dpfm_api_output_formatter.OrdersItemPricingElement{}, err
		}
		ordersItemPricingElements = append(ordersItemPricingElements, ordersItemPricingElement)
	}

	return ordersItemPricingElements, nil
}

func structuralizeItemScheduleLine(ordersSDC dpfm_api_processing_data_formatter.OrdersSDC, i int) ([]dpfm_api_output_formatter.OrdersItemScheduleLine, error) {
	ordersItemScheduleLines := make([]dpfm_api_output_formatter.OrdersItemScheduleLine, 0)
	for _, itemScheduleLine := range ordersSDC.HeaderAndItemAndScheduleLine[i].ItemScheduleLine {
		ordersItemScheduleLine, err := dpfm_api_output_formatter.TypeConverter[dpfm_api_output_formatter.OrdersItemScheduleLine](itemScheduleLine)
		if err != nil {
			return []dpfm_api_output_formatter.OrdersItemScheduleLine{}, err
		}
		ordersItemScheduleLines = append(ordersItemScheduleLines, ordersItemScheduleLine)
	}

	return ordersItemScheduleLines, nil
}

func structuralizePartner(ordersSDC dpfm_api_processing_data_formatter.OrdersSDC) ([]dpfm_api_output_formatter.OrdersPartner, error) {
	ordersPartners := make([]dpfm_api_output_formatter.OrdersPartner, 0)
	for _, partner := range ordersSDC.HeaderAndPartner.Partner {
		ordersPartner, err := dpfm_api_output_formatter.TypeConverter[dpfm_api_output_formatter.OrdersPartner](partner)
		if err != nil {
			return []dpfm_api_output_formatter.OrdersPartner{}, err
		}
		ordersPartners = append(ordersPartners, ordersPartner)
	}

	return ordersPartners, nil
}

func structuralizeAddress(ordersSDC dpfm_api_processing_data_formatter.OrdersSDC) ([]dpfm_api_output_formatter.OrdersAddress, error) {
	ordersAddresses := make([]dpfm_api_output_formatter.OrdersAddress, 0)
	for _, address := range ordersSDC.HeaderAndAddress.Address {
		ordersAddress, err := dpfm_api_output_formatter.TypeConverter[dpfm_api_output_formatter.OrdersAddress](address)
		if err != nil {
			return []dpfm_api_output_formatter.OrdersAddress{}, err
		}
		ordersAddresses = append(ordersAddresses, ordersAddress)
	}

	return ordersAddresses, nil
}
