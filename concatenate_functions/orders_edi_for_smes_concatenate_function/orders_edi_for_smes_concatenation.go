package orders_edi_for_smes_concatenate_function

import (
	dpfm_api_input_reader "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_data_formatter "data-platform-api-data-concatenation-rmq-kube/DPFM_API_Processing_Data_Formatter"
	"sync"
)

func (c *OrdersEDIForSMEsContraller) OrdersEDIForSMEsConcatenation(concatenateMapper []dpfm_api_input_reader.ConcatenateMapper) (dpfm_api_processing_data_formatter.OrdersEDIForSMEsSDC, error) {
	var err error
	var e error

	sdc := dpfm_api_input_reader.ConvertToSDC[dpfm_api_input_reader.OrdersEDIForSMEsSDC](c.msg.Raw())
	headerAndItem := dpfm_api_processing_data_formatter.OrdersEDIForSMEsHeaderAndItem{}

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
		} else {
			wg.Done()
		}
	}

	wg.Wait()
	if err != nil {
		return dpfm_api_processing_data_formatter.OrdersEDIForSMEsSDC{}, err
	}

	ordersEDIForSMEsConcatenated := dpfm_api_processing_data_formatter.OrdersEDIForSMEsSDC{
		HeaderAndItem: headerAndItem,
	}

	return ordersEDIForSMEsConcatenated, nil
}
