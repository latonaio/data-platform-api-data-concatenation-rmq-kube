package orders_edi_for_voluntary_chain_smes_concatenate_function

import (
	dpfm_api_input_reader "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_data_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Processing_Data_Formatter"

	"golang.org/x/xerrors"
)

func newHeaderAndItem(sdc dpfm_api_input_reader.OrdersEDIForVoluntaryChainSMEsSDC) (dpfm_api_processing_data_formatter.OrdersEDIForVoluntaryChainSMEsHeaderAndItem, error) {
	processingHeader, err := dpfm_api_processing_data_formatter.TypeConverter[dpfm_api_processing_data_formatter.OrdersEDIForVoluntaryChainSMEsHeader](sdc.DataConcatenation.Header)
	if err != nil {
		return dpfm_api_processing_data_formatter.OrdersEDIForVoluntaryChainSMEsHeaderAndItem{}, xerrors.Errorf("TypeConverter Error: %w", err)
	}

	exchangedOrdersDocumentIdentifier := processingHeader.ExchangedOrdersDocumentIdentifier
	processingItems, err := extractItem(exchangedOrdersDocumentIdentifier, sdc.DataConcatenation.Item)
	if err != nil {
		return dpfm_api_processing_data_formatter.OrdersEDIForVoluntaryChainSMEsHeaderAndItem{}, xerrors.Errorf("extractItem Error: %w", err)
	}

	headerAndItem := dpfm_api_processing_data_formatter.OrdersEDIForVoluntaryChainSMEsHeaderAndItem{
		Header: processingHeader,
		Item:   processingItems,
	}

	return headerAndItem, err
}

func extractItem(exchangedOrdersDocumentIdentifier string, items []dpfm_api_input_reader.OrdersEDIForVoluntaryChainSMEsItem) ([]dpfm_api_processing_data_formatter.OrdersEDIForVoluntaryChainSMEsItem, error) {
	processingItems := make([]dpfm_api_processing_data_formatter.OrdersEDIForVoluntaryChainSMEsItem, 0)

	for _, item := range items {
		if item.ExchangedOrdersDocumentIdentifier != exchangedOrdersDocumentIdentifier {
			continue
		}
		processingItem, err := dpfm_api_processing_data_formatter.TypeConverter[dpfm_api_processing_data_formatter.OrdersEDIForVoluntaryChainSMEsItem](item)
		if err != nil {
			return []dpfm_api_processing_data_formatter.OrdersEDIForVoluntaryChainSMEsItem{}, xerrors.Errorf("TypeConverter Error: %w", err)
		}
		processingItems = append(processingItems, processingItem)
	}

	return processingItems, nil
}
