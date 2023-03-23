package dpfm_api_output_formatter

import (
	"encoding/json"
	"fmt"
	"os"
)

func ConvertToOrdersSDC(data []byte) OrdersSDC {
	sdc := OrdersSDC{}
	err := json.Unmarshal(data, &sdc)
	if err != nil {
		fmt.Printf("input data marshal error :%#v", err.Error())
		os.Exit(1)
	}

	return sdc
}

func ConvertToDeliveryDocumentSDC(data []byte) DeliveryDocumentSDC {
	sdc := DeliveryDocumentSDC{}
	err := json.Unmarshal(data, &sdc)
	if err != nil {
		fmt.Printf("input data marshal error :%#v", err.Error())
		os.Exit(1)
	}

	return sdc
}

func ConvertToOrdersEDIForSMEsSDC(data []byte) OrdersEDIForSMEsSDC {
	sdc := OrdersEDIForSMEsSDC{}
	err := json.Unmarshal(data, &sdc)
	if err != nil {
		fmt.Printf("input data marshal error :%#v", err.Error())
		os.Exit(1)
	}

	return sdc
}

func ConvertToOrdersEDIForVoluntaryChainSMEsSDC(data []byte) OrdersEDIForVoluntaryChainSMEsSDC {
	sdc := OrdersEDIForVoluntaryChainSMEsSDC{}
	err := json.Unmarshal(data, &sdc)
	if err != nil {
		fmt.Printf("input data marshal error :%#v", err.Error())
		os.Exit(1)
	}

	return sdc
}

func ConvertToDeliveryNoticeEDIForSMEsSDC(data []byte) DeliveryNoticeEDIForSMEsSDC {
	sdc := DeliveryNoticeEDIForSMEsSDC{}
	err := json.Unmarshal(data, &sdc)
	if err != nil {
		fmt.Printf("input data marshal error :%#v", err.Error())
		os.Exit(1)
	}

	return sdc
}

func ConvertToDeliveryNoticeEDIForVoluntaryChainSMEsSDC(data []byte) DeliveryNoticeEDIForVoluntaryChainSMEsSDC {
	sdc := DeliveryNoticeEDIForVoluntaryChainSMEsSDC{}
	err := json.Unmarshal(data, &sdc)
	if err != nil {
		fmt.Printf("input data marshal error :%#v", err.Error())
		os.Exit(1)
	}

	return sdc
}
