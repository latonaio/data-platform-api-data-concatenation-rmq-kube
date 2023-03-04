package orders_concatenate_function

import (
	dpfm_api_input_reader "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_data_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Processing_Data_Formatter"

	"golang.org/x/xerrors"
)

func newHeaderAndAddress(sdc dpfm_api_input_reader.OrdersSDC) (dpfm_api_processing_data_formatter.OrdersHeaderAndAddress, error) {
	processingHeader, err := dpfm_api_processing_data_formatter.TypeConverter[dpfm_api_processing_data_formatter.OrdersHeader](sdc.DataConcatenation.Header)
	if err != nil {
		return dpfm_api_processing_data_formatter.OrdersHeaderAndAddress{}, xerrors.Errorf("TypeConverter Error: %w", err)
	}

	orderID := processingHeader.OrderID
	processingAddressses, err := extractAddress(orderID, sdc.DataConcatenation.Address)
	if err != nil {
		return dpfm_api_processing_data_formatter.OrdersHeaderAndAddress{}, xerrors.Errorf("extractAddress Error: %w", err)
	}

	headerAndAddress := dpfm_api_processing_data_formatter.OrdersHeaderAndAddress{
		Header:  processingHeader,
		Address: processingAddressses,
	}

	return headerAndAddress, err
}

func extractAddress(orderID int, addresses []dpfm_api_input_reader.OrdersAddress) ([]dpfm_api_processing_data_formatter.OrdersAddress, error) {
	processingAddresses := make([]dpfm_api_processing_data_formatter.OrdersAddress, 0)

	for _, address := range addresses {
		if address.OrderID != orderID {
			continue
		}
		processingAddress, err := dpfm_api_processing_data_formatter.TypeConverter[dpfm_api_processing_data_formatter.OrdersAddress](address)
		if err != nil {
			return []dpfm_api_processing_data_formatter.OrdersAddress{}, xerrors.Errorf("TypeConverter Error: %w", err)
		}
		processingAddresses = append(processingAddresses, processingAddress)
	}

	return processingAddresses, nil
}
