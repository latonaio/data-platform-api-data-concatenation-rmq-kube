package orders_edi_for_voluntary_chain_smes_concatenate_function

import (
	dpfm_api_output_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Output_Formatter"
	dpfm_api_processing_data_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Processing_Data_Formatter"
	"strings"
)

func (c *OrdersEDIForVoluntaryChainSMEsContraller) OrdersEDIForVoluntaryChainSMEsStructuralization(ordersEDIForVoluntaryChainSMEsSDC dpfm_api_processing_data_formatter.OrdersEDIForVoluntaryChainSMEsSDC) (dpfm_api_output_formatter.OrdersEDIForVoluntaryChainSMEsSDC, error) {
	output := dpfm_api_output_formatter.ConvertToOrdersEDIForVoluntaryChainSMEsSDC(c.msg.Raw())

	// Header
	ordersEDIForVoluntaryChainSMEsHeader, err := structuralizeHeader(ordersEDIForVoluntaryChainSMEsSDC)
	if err != nil {
		return dpfm_api_output_formatter.OrdersEDIForVoluntaryChainSMEsSDC{}, err
	}

	// Item
	ordersEDIForVoluntaryChainSMEsItems, err := structuralizeItem(ordersEDIForVoluntaryChainSMEsSDC)
	if err != nil {
		return dpfm_api_output_formatter.OrdersEDIForVoluntaryChainSMEsSDC{}, err
	}

	ordersEDIForVoluntaryChainSMEsHeader.Item = ordersEDIForVoluntaryChainSMEsItems

	output.Header = ordersEDIForVoluntaryChainSMEsHeader
	output.ServiceLabel = "FUNCTION_CONVERT_TO_DPFM_ORDERS_FROM_ORDERS_EDI_FOR_VOLUNTARY_CHAIN_SMES"

	switch strings.ToLower(output.APIType) {
	case "creates":
		output.APISchema = "OrdersEDIForVoluntaryChainSMEsCreates"
	case "updates":
		output.APISchema = "OrdersEDIForVoluntaryChainSMEsUpdates"
	case "function":
		output.APISchema = "OrdersEDIForVoluntaryChainSMEsCreates"
	default:
		c.log.Error("unknown apitype %s", output.APIType)
		output.APIType = "creates"
		output.APISchema = "OrdersEDIForVoluntaryChainSMEsCreates"
	}

	return output, nil
}

func structuralizeHeader(ordersEDIForVoluntaryChainSMEsSDC dpfm_api_processing_data_formatter.OrdersEDIForVoluntaryChainSMEsSDC) (dpfm_api_output_formatter.OrdersEDIForVoluntaryChainSMEsHeader, error) {
	ordersEDIForVoluntaryChainSMEsHeader, err := dpfm_api_output_formatter.TypeConverter[dpfm_api_output_formatter.OrdersEDIForVoluntaryChainSMEsHeader](ordersEDIForVoluntaryChainSMEsSDC.HeaderAndItem.Header)
	if err != nil {
		return dpfm_api_output_formatter.OrdersEDIForVoluntaryChainSMEsHeader{}, err
	}

	return ordersEDIForVoluntaryChainSMEsHeader, nil
}

func structuralizeItem(ordersEDIForVoluntaryChainSMEsSDC dpfm_api_processing_data_formatter.OrdersEDIForVoluntaryChainSMEsSDC) ([]dpfm_api_output_formatter.OrdersEDIForVoluntaryChainSMEsItem, error) {
	ordersEDIForVoluntaryChainSMEsItems := make([]dpfm_api_output_formatter.OrdersEDIForVoluntaryChainSMEsItem, 0)
	for _, item := range ordersEDIForVoluntaryChainSMEsSDC.HeaderAndItem.Item {
		ordersEDIForVoluntaryChainSMEsItem, err := dpfm_api_output_formatter.TypeConverter[dpfm_api_output_formatter.OrdersEDIForVoluntaryChainSMEsItem](item)
		if err != nil {
			return []dpfm_api_output_formatter.OrdersEDIForVoluntaryChainSMEsItem{}, err
		}

		ordersEDIForVoluntaryChainSMEsItems = append(ordersEDIForVoluntaryChainSMEsItems, ordersEDIForVoluntaryChainSMEsItem)
	}

	return ordersEDIForVoluntaryChainSMEsItems, nil
}
