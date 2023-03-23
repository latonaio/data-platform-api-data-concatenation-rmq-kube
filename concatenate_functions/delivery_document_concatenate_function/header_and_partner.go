package delivery_document_concatenate_function

import (
	dpfm_api_input_reader "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_data_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Processing_Data_Formatter"

	"golang.org/x/xerrors"
)

func newHeaderAndPartner(sdc dpfm_api_input_reader.DeliveryDocumentSDC) (dpfm_api_processing_data_formatter.DeliveryDocumentHeaderAndPartner, error) {
	processingHeader, err := dpfm_api_processing_data_formatter.TypeConverter[dpfm_api_processing_data_formatter.DeliveryDocumentHeader](sdc.DataConcatenation.Header)
	if err != nil {
		return dpfm_api_processing_data_formatter.DeliveryDocumentHeaderAndPartner{}, xerrors.Errorf("TypeConverter Error: %w", err)
	}

	deliveryDocument := processingHeader.DeliveryDocument
	processingPartners, err := extractPartner(deliveryDocument, sdc.DataConcatenation.Partner)
	if err != nil {
		return dpfm_api_processing_data_formatter.DeliveryDocumentHeaderAndPartner{}, xerrors.Errorf("extractPartner Error: %w", err)
	}

	headerAndPartner := dpfm_api_processing_data_formatter.DeliveryDocumentHeaderAndPartner{
		Header:  processingHeader,
		Partner: processingPartners,
	}

	return headerAndPartner, err
}

func extractPartner(deliveryDocument int, partners []dpfm_api_input_reader.DeliveryDocumentPartner) ([]dpfm_api_processing_data_formatter.DeliveryDocumentPartner, error) {
	processingPartners := make([]dpfm_api_processing_data_formatter.DeliveryDocumentPartner, 0)

	for _, partner := range partners {
		if partner.DeliveryDocument != deliveryDocument {
			continue
		}
		processingPartner, err := dpfm_api_processing_data_formatter.TypeConverter[dpfm_api_processing_data_formatter.DeliveryDocumentPartner](partner)
		if err != nil {
			return []dpfm_api_processing_data_formatter.DeliveryDocumentPartner{}, xerrors.Errorf("TypeConverter Error: %w", err)
		}
		processingPartners = append(processingPartners, processingPartner)
	}

	return processingPartners, nil
}
