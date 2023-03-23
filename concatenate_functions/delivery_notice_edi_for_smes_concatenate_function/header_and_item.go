package delivery_notice_edi_for_smes_concatenate_function

import (
	dpfm_api_input_reader "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_data_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Processing_Data_Formatter"

	"golang.org/x/xerrors"
)

func newHeaderAndItem(sdc dpfm_api_input_reader.DeliveryNoticeEDIForSMEsSDC) (dpfm_api_processing_data_formatter.DeliveryNoticeEDIForSMEsHeaderAndItem, error) {
	processingHeader, err := dpfm_api_processing_data_formatter.TypeConverter[dpfm_api_processing_data_formatter.DeliveryNoticeEDIForSMEsHeader](sdc.DataConcatenation.Header)
	if err != nil {
		return dpfm_api_processing_data_formatter.DeliveryNoticeEDIForSMEsHeaderAndItem{}, xerrors.Errorf("TypeConverter Error: %w", err)
	}

	exchangedDeliveryNoticeDocumentIdentifier := processingHeader.ExchangedDeliveryNoticeDocumentIdentifier
	processingItems, err := extractItem(exchangedDeliveryNoticeDocumentIdentifier, sdc.DataConcatenation.Item)
	if err != nil {
		return dpfm_api_processing_data_formatter.DeliveryNoticeEDIForSMEsHeaderAndItem{}, xerrors.Errorf("extractItem Error: %w", err)
	}

	headerAndItem := dpfm_api_processing_data_formatter.DeliveryNoticeEDIForSMEsHeaderAndItem{
		Header: processingHeader,
		Item:   processingItems,
	}

	return headerAndItem, err
}

func extractItem(exchangedDeliveryNoticeDocumentIdentifier string, items []dpfm_api_input_reader.DeliveryNoticeEDIForSMEsItem) ([]dpfm_api_processing_data_formatter.DeliveryNoticeEDIForSMEsItem, error) {
	processingItems := make([]dpfm_api_processing_data_formatter.DeliveryNoticeEDIForSMEsItem, 0)

	for _, item := range items {
		if item.ExchangedDeliveryNoticeDocumentIdentifier != exchangedDeliveryNoticeDocumentIdentifier {
			continue
		}
		processingItem, err := dpfm_api_processing_data_formatter.TypeConverter[dpfm_api_processing_data_formatter.DeliveryNoticeEDIForSMEsItem](item)
		if err != nil {
			return []dpfm_api_processing_data_formatter.DeliveryNoticeEDIForSMEsItem{}, xerrors.Errorf("TypeConverter Error: %w", err)
		}
		processingItems = append(processingItems, processingItem)
	}

	return processingItems, nil
}
