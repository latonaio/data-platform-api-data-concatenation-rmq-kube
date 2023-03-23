package delivery_notice_edi_for_voluntary_chain_smes_concatenate_function

import (
	dpfm_api_input_reader "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_data_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Processing_Data_Formatter"

	"golang.org/x/xerrors"
)

func newHeaderAndItem(sdc dpfm_api_input_reader.DeliveryNoticeEDIForVoluntaryChainSMEsSDC) (dpfm_api_processing_data_formatter.DeliveryNoticeEDIForVoluntaryChainSMEsHeaderAndItem, error) {
	processingHeader, err := dpfm_api_processing_data_formatter.TypeConverter[dpfm_api_processing_data_formatter.DeliveryNoticeEDIForVoluntaryChainSMEsHeader](sdc.DataConcatenation.Header)
	if err != nil {
		return dpfm_api_processing_data_formatter.DeliveryNoticeEDIForVoluntaryChainSMEsHeaderAndItem{}, xerrors.Errorf("TypeConverter Error: %w", err)
	}

	exchangedDeliveryNoticeDocumentIdentifier := processingHeader.ExchangedDeliveryNoticeDocumentIdentifier
	processingItems, err := extractItem(exchangedDeliveryNoticeDocumentIdentifier, sdc.DataConcatenation.Item)
	if err != nil {
		return dpfm_api_processing_data_formatter.DeliveryNoticeEDIForVoluntaryChainSMEsHeaderAndItem{}, xerrors.Errorf("extractItem Error: %w", err)
	}

	headerAndItem := dpfm_api_processing_data_formatter.DeliveryNoticeEDIForVoluntaryChainSMEsHeaderAndItem{
		Header: processingHeader,
		Item:   processingItems,
	}

	return headerAndItem, err
}

func extractItem(exchangedDeliveryNoticeDocumentIdentifier string, items []dpfm_api_input_reader.DeliveryNoticeEDIForVoluntaryChainSMEsItem) ([]dpfm_api_processing_data_formatter.DeliveryNoticeEDIForVoluntaryChainSMEsItem, error) {
	processingItems := make([]dpfm_api_processing_data_formatter.DeliveryNoticeEDIForVoluntaryChainSMEsItem, 0)

	for _, item := range items {
		if item.ExchangedDeliveryNoticeDocumentIdentifier != exchangedDeliveryNoticeDocumentIdentifier {
			continue
		}
		processingItem, err := dpfm_api_processing_data_formatter.TypeConverter[dpfm_api_processing_data_formatter.DeliveryNoticeEDIForVoluntaryChainSMEsItem](item)
		if err != nil {
			return []dpfm_api_processing_data_formatter.DeliveryNoticeEDIForVoluntaryChainSMEsItem{}, xerrors.Errorf("TypeConverter Error: %w", err)
		}
		processingItems = append(processingItems, processingItem)
	}

	return processingItems, nil
}
