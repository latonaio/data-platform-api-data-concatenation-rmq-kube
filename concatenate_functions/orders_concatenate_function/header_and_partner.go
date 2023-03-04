package orders_concatenate_function

import (
	dpfm_api_input_reader "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_data_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Processing_Data_Formatter"

	"golang.org/x/xerrors"
)

func newHeaderAndPartner(sdc dpfm_api_input_reader.OrdersSDC) (dpfm_api_processing_data_formatter.OrdersHeaderAndPartner, error) {
	processingHeader, err := dpfm_api_processing_data_formatter.TypeConverter[dpfm_api_processing_data_formatter.OrdersHeader](sdc.DataConcatenation.Header)
	if err != nil {
		return dpfm_api_processing_data_formatter.OrdersHeaderAndPartner{}, xerrors.Errorf("TypeConverter Error: %w", err)
	}

	orderID := processingHeader.OrderID
	processingPartners, err := extractPartner(orderID, sdc.DataConcatenation.Partner)
	if err != nil {
		return dpfm_api_processing_data_formatter.OrdersHeaderAndPartner{}, xerrors.Errorf("extractPartner Error: %w", err)
	}

	headerAndPartner := dpfm_api_processing_data_formatter.OrdersHeaderAndPartner{
		Header:  processingHeader,
		Partner: processingPartners,
	}

	return headerAndPartner, err
}

func extractPartner(orderID int, partners []dpfm_api_input_reader.OrdersPartner) ([]dpfm_api_processing_data_formatter.OrdersPartner, error) {
	processingPartners := make([]dpfm_api_processing_data_formatter.OrdersPartner, 0)

	for _, partner := range partners {
		if partner.OrderID != orderID {
			continue
		}
		processingPartner, err := dpfm_api_processing_data_formatter.TypeConverter[dpfm_api_processing_data_formatter.OrdersPartner](partner)
		if err != nil {
			return []dpfm_api_processing_data_formatter.OrdersPartner{}, xerrors.Errorf("TypeConverter Error: %w", err)
		}
		processingPartners = append(processingPartners, processingPartner)
	}

	return processingPartners, nil
}
