package concatenate_functions

import (
	dpfm_api_input_reader "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Input_Reader"
	api_processing_data_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Processing_Data_Formatter"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

func OrdersConcatenate(msg rabbitmq.RabbitmqMessage, concatenateMapper []dpfm_api_input_reader.ConcatenateMapper) (api_processing_data_formatter.OrdersSDC, error) {
	sdc := dpfm_api_input_reader.ConvertToOrdersSDC(msg.Raw())

	var err error
	headerAndItem := api_processing_data_formatter.OrdersHeaderAndItem{}
	headerAndItemAndPricingElement := make([]api_processing_data_formatter.OrdersHeaderAndItemAndPricingElement, 0)
	headerAndItemAndScheduleLine := make([]api_processing_data_formatter.OrdersHeaderAndItemAndScheduleLine, 0)
	headerAndPartner := api_processing_data_formatter.OrdersHeaderAndPartner{}
	headerAndAddress := api_processing_data_formatter.OrdersHeaderAndAddress{}

	for _, v := range concatenateMapper {
		baseAPIName := v.BaseAPIName
		concatenateAPIName1 := v.ConcatenateAPIName1
		concatenateAPIName2 := v.ConcatenateAPIName2

		if baseAPIName == "A_Header" && concatenateAPIName1 == "A_Item" {
			headerAndItem, err = newHeaderAndItem(sdc)
			if err != nil {
				return api_processing_data_formatter.OrdersSDC{}, xerrors.Errorf("newHeaderAndItem Error: %w", err)
			}
		}

		if baseAPIName == "A_Header" && concatenateAPIName1 == "A_Item" && concatenateAPIName2 == "A_ItemPricingElement" {
			headerAndItemAndPricingElement, err = newHeaderAndItemAndPricingElement(sdc)
			if err != nil {
				return api_processing_data_formatter.OrdersSDC{}, xerrors.Errorf("newHeaderAndItemAndPricingElement Error: %w", err)
			}
		}

		if baseAPIName == "A_Header" && concatenateAPIName1 == "A_Item" && concatenateAPIName2 == "A_ItemScheduleLine" {
			headerAndItemAndScheduleLine, err = newHeaderAndItemAndScheduleLine(sdc)
			if err != nil {
				return api_processing_data_formatter.OrdersSDC{}, xerrors.Errorf("newHeaderAndItemAndScheduleLine Error: %w", err)
			}
		}

		if baseAPIName == "A_Header" && concatenateAPIName1 == "A_Partner" {
			headerAndPartner, err = newHeaderAndPartner(sdc)
			if err != nil {
				return api_processing_data_formatter.OrdersSDC{}, xerrors.Errorf("newHeaderAndPartner Error: %w", err)
			}
		}

		if baseAPIName == "A_Header" && concatenateAPIName1 == "A_Address" {
			headerAndAddress, err = newHeaderAndAddress(sdc)
			if err != nil {
				return api_processing_data_formatter.OrdersSDC{}, xerrors.Errorf("newHeaderAndAddress Error: %w", err)
			}
		}
	}

	ordersSDC := api_processing_data_formatter.OrdersSDC{
		HeaderAndItem:                  headerAndItem,
		HeaderAndItemAndPricingElement: headerAndItemAndPricingElement,
		HeaderAndItemAndScheduleLine:   headerAndItemAndScheduleLine,
		HeaderAndPartner:               headerAndPartner,
		HeaderAndAddress:               headerAndAddress,
	}

	return ordersSDC, nil
}

func newHeaderAndItem(sdc dpfm_api_input_reader.OrdersSDC) (api_processing_data_formatter.OrdersHeaderAndItem, error) {
	processingHeader, err := api_processing_data_formatter.TypeConverter[api_processing_data_formatter.OrdersHeader](sdc.Message.Header)
	if err != nil {
		return api_processing_data_formatter.OrdersHeaderAndItem{}, xerrors.Errorf("TypeConverter Error: %w", err)
	}

	orderID := processingHeader.OrderID
	processingItems, err := extractItem(orderID, sdc.Message.Item)
	if err != nil {
		return api_processing_data_formatter.OrdersHeaderAndItem{}, xerrors.Errorf("extractItem Error: %w", err)
	}

	headerAndItem := api_processing_data_formatter.OrdersHeaderAndItem{
		Header: processingHeader,
		Item:   processingItems,
	}

	return headerAndItem, err
}

func extractItem(orderID int, items []dpfm_api_input_reader.OrdersItem) ([]api_processing_data_formatter.OrdersItem, error) {
	processingItems := make([]api_processing_data_formatter.OrdersItem, 0)

	for _, item := range items {
		if item.OrderID != orderID {
			continue
		}
		processingItem, err := api_processing_data_formatter.TypeConverter[api_processing_data_formatter.OrdersItem](item)
		if err != nil {
			return []api_processing_data_formatter.OrdersItem{}, xerrors.Errorf("TypeConverter Error: %w", err)
		}
		processingItems = append(processingItems, processingItem)
	}

	return processingItems, nil
}

func newHeaderAndItemAndPricingElement(sdc dpfm_api_input_reader.OrdersSDC) ([]api_processing_data_formatter.OrdersHeaderAndItemAndPricingElement, error) {
	headerAndItemAndPricingElements := make([]api_processing_data_formatter.OrdersHeaderAndItemAndPricingElement, 0)

	processingHeader, err := api_processing_data_formatter.TypeConverter[api_processing_data_formatter.OrdersHeader](sdc.Message.Header)
	if err != nil {
		return []api_processing_data_formatter.OrdersHeaderAndItemAndPricingElement{}, xerrors.Errorf("TypeConverter Error: %w", err)
	}

	for _, item := range sdc.Message.Item {
		orderID := item.OrderID
		orderItem := item.OrderItem
		if orderID != processingHeader.OrderID {
			continue
		}
		processingItem, err := api_processing_data_formatter.TypeConverter[api_processing_data_formatter.OrdersItem](item)
		if err != nil {
			return []api_processing_data_formatter.OrdersHeaderAndItemAndPricingElement{}, xerrors.Errorf("TypeConverter Error: %w", err)
		}

		processingItemPricingElements, err := extractItemPricingElement(orderID, orderItem, sdc.Message.ItemPricingElement)
		if err != nil {
			return []api_processing_data_formatter.OrdersHeaderAndItemAndPricingElement{}, xerrors.Errorf("extractItemPricingElement Error: %w", err)
		}

		headerAndItemAndPricingElements = append(headerAndItemAndPricingElements, api_processing_data_formatter.OrdersHeaderAndItemAndPricingElement{
			Header:             processingHeader,
			Item:               processingItem,
			ItemPricingElement: processingItemPricingElements,
		})
	}

	return headerAndItemAndPricingElements, err
}

func extractItemPricingElement(orderID, orderItem int, itemPricingElements []dpfm_api_input_reader.OrdersItemPricingElement) ([]api_processing_data_formatter.OrdersItemPricingElement, error) {
	processingItemPricingElements := make([]api_processing_data_formatter.OrdersItemPricingElement, 0)

	for _, itemPricingElement := range itemPricingElements {
		if itemPricingElement.OrderID != orderID || itemPricingElement.OrderItem != orderItem {
			continue
		}
		processingItemPricingElement, err := api_processing_data_formatter.TypeConverter[api_processing_data_formatter.OrdersItemPricingElement](itemPricingElement)
		if err != nil {
			return []api_processing_data_formatter.OrdersItemPricingElement{}, err
		}

		processingItemPricingElements = append(processingItemPricingElements, processingItemPricingElement)
	}

	return processingItemPricingElements, nil
}

func newHeaderAndItemAndScheduleLine(sdc dpfm_api_input_reader.OrdersSDC) ([]api_processing_data_formatter.OrdersHeaderAndItemAndScheduleLine, error) {
	headerAndItemAndScheduleLines := make([]api_processing_data_formatter.OrdersHeaderAndItemAndScheduleLine, 0)

	processingHeader, err := api_processing_data_formatter.TypeConverter[api_processing_data_formatter.OrdersHeader](sdc.Message.Header)
	if err != nil {
		return []api_processing_data_formatter.OrdersHeaderAndItemAndScheduleLine{}, xerrors.Errorf("TypeConverter Error: %w", err)
	}

	for _, item := range sdc.Message.Item {
		orderID := item.OrderID
		orderItem := item.OrderItem
		if orderID != processingHeader.OrderID {
			continue
		}
		processingItem, err := api_processing_data_formatter.TypeConverter[api_processing_data_formatter.OrdersItem](item)
		if err != nil {
			return []api_processing_data_formatter.OrdersHeaderAndItemAndScheduleLine{}, xerrors.Errorf("TypeConverter Error: %w", err)
		}

		processingItemScheduleLines, err := extractItemScheduleLine(orderID, orderItem, sdc.Message.ItemScheduleLine)
		if err != nil {
			return []api_processing_data_formatter.OrdersHeaderAndItemAndScheduleLine{}, xerrors.Errorf("extractItemScheduleLine Error: %w", err)
		}

		headerAndItemAndScheduleLines = append(headerAndItemAndScheduleLines, api_processing_data_formatter.OrdersHeaderAndItemAndScheduleLine{
			Header:           processingHeader,
			Item:             processingItem,
			ItemScheduleLine: processingItemScheduleLines,
		})
	}

	return headerAndItemAndScheduleLines, err
}

func extractItemScheduleLine(orderID, orderItem int, itemScheduleLines []dpfm_api_input_reader.OrdersItemScheduleLine) ([]api_processing_data_formatter.OrdersItemScheduleLine, error) {
	processingItemScheduleLines := make([]api_processing_data_formatter.OrdersItemScheduleLine, 0)

	for _, itemScheduleLine := range itemScheduleLines {
		if itemScheduleLine.OrderID != orderID || itemScheduleLine.OrderItem != orderItem {
			continue
		}
		processingItemScheduleLine, err := api_processing_data_formatter.TypeConverter[api_processing_data_formatter.OrdersItemScheduleLine](itemScheduleLine)
		if err != nil {
			return []api_processing_data_formatter.OrdersItemScheduleLine{}, err
		}

		processingItemScheduleLines = append(processingItemScheduleLines, processingItemScheduleLine)
	}

	return processingItemScheduleLines, nil
}

func newHeaderAndPartner(sdc dpfm_api_input_reader.OrdersSDC) (api_processing_data_formatter.OrdersHeaderAndPartner, error) {
	processingHeader, err := api_processing_data_formatter.TypeConverter[api_processing_data_formatter.OrdersHeader](sdc.Message.Header)
	if err != nil {
		return api_processing_data_formatter.OrdersHeaderAndPartner{}, xerrors.Errorf("TypeConverter Error: %w", err)
	}

	orderID := processingHeader.OrderID
	processingPartners, err := extractPartner(orderID, sdc.Message.Partner)
	if err != nil {
		return api_processing_data_formatter.OrdersHeaderAndPartner{}, xerrors.Errorf("extractPartner Error: %w", err)
	}

	headerAndPartner := api_processing_data_formatter.OrdersHeaderAndPartner{
		Header:  processingHeader,
		Partner: processingPartners,
	}

	return headerAndPartner, err
}

func extractPartner(orderID int, partners []dpfm_api_input_reader.OrdersPartner) ([]api_processing_data_formatter.OrdersPartner, error) {
	processingPartners := make([]api_processing_data_formatter.OrdersPartner, 0)

	for _, partner := range partners {
		if partner.OrderID != orderID {
			continue
		}
		processingPartner, err := api_processing_data_formatter.TypeConverter[api_processing_data_formatter.OrdersPartner](partner)
		if err != nil {
			return []api_processing_data_formatter.OrdersPartner{}, xerrors.Errorf("TypeConverter Error: %w", err)
		}
		processingPartners = append(processingPartners, processingPartner)
	}

	return processingPartners, nil
}

func newHeaderAndAddress(sdc dpfm_api_input_reader.OrdersSDC) (api_processing_data_formatter.OrdersHeaderAndAddress, error) {
	processingHeader, err := api_processing_data_formatter.TypeConverter[api_processing_data_formatter.OrdersHeader](sdc.Message.Header)
	if err != nil {
		return api_processing_data_formatter.OrdersHeaderAndAddress{}, xerrors.Errorf("TypeConverter Error: %w", err)
	}

	orderID := processingHeader.OrderID
	processingAddressses, err := extractAddress(orderID, sdc.Message.Address)
	if err != nil {
		return api_processing_data_formatter.OrdersHeaderAndAddress{}, xerrors.Errorf("extractAddress Error: %w", err)
	}

	headerAndAddress := api_processing_data_formatter.OrdersHeaderAndAddress{
		Header:  processingHeader,
		Address: processingAddressses,
	}

	return headerAndAddress, err
}

func extractAddress(orderID int, addresses []dpfm_api_input_reader.OrdersAddress) ([]api_processing_data_formatter.OrdersAddress, error) {
	processingAddresses := make([]api_processing_data_formatter.OrdersAddress, 0)

	for _, address := range addresses {
		if address.OrderID != orderID {
			continue
		}
		processingAddress, err := api_processing_data_formatter.TypeConverter[api_processing_data_formatter.OrdersAddress](address)
		if err != nil {
			return []api_processing_data_formatter.OrdersAddress{}, xerrors.Errorf("TypeConverter Error: %w", err)
		}
		processingAddresses = append(processingAddresses, processingAddress)
	}

	return processingAddresses, nil
}
