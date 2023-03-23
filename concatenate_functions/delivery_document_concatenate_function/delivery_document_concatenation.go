package delivery_document_concatenate_function

import (
	dpfm_api_input_reader "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_data_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Processing_Data_Formatter"
	"sync"
)

func (c *DeliveryDocumentContraller) DeliveryDocumentConcatenation(concatenateMapper []dpfm_api_input_reader.ConcatenateMapper) (dpfm_api_processing_data_formatter.DeliveryDocumentSDC, error) {
	var err error
	var e error

	sdc := dpfm_api_input_reader.ConvertToDeliveryDocumentSDC(c.msg.Raw())
	headerAndItem := dpfm_api_processing_data_formatter.DeliveryDocumentHeaderAndItem{}
	headerAndPartner := dpfm_api_processing_data_formatter.DeliveryDocumentHeaderAndPartner{}
	headerAndAddress := dpfm_api_processing_data_formatter.DeliveryDocumentHeaderAndAddress{}

	wg := sync.WaitGroup{}
	for _, v := range concatenateMapper {
		wg.Add(1)
		baseAPIName := v.BaseAPIName
		concatenateAPIName1 := v.ConcatenateAPIName1
		concatenateAPIName2 := v.ConcatenateAPIName2

		if baseAPIName == "A_Header" && concatenateAPIName1 == "A_Item" && concatenateAPIName2 == "" {
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				headerAndItem, e = newHeaderAndItem(sdc)
				if e != nil {
					err = e
					return
				}
			}(&wg)
		} else if baseAPIName == "A_Header" && concatenateAPIName1 == "A_Partner" {
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				headerAndPartner, e = newHeaderAndPartner(sdc)
				if e != nil {
					err = e
					return
				}
			}(&wg)
		} else if baseAPIName == "A_Header" && concatenateAPIName1 == "A_Address" {
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				headerAndAddress, e = newHeaderAndAddress(sdc)
				if e != nil {
					err = e
					return
				}
			}(&wg)
		} else {
			wg.Done()
		}
	}

	wg.Wait()
	if err != nil {
		return dpfm_api_processing_data_formatter.DeliveryDocumentSDC{}, err
	}

	deliveryDocumentConcatenated := dpfm_api_processing_data_formatter.DeliveryDocumentSDC{
		HeaderAndItem:    headerAndItem,
		HeaderAndPartner: headerAndPartner,
		HeaderAndAddress: headerAndAddress,
	}

	return deliveryDocumentConcatenated, nil
}
