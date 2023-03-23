package dpfm_api_input_reader

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

func ConvertToSDC[T any](data []byte) T {
	var sdc T
	err := json.Unmarshal(data, &sdc)
	if err != nil {
		fmt.Printf("input data marshal error :%#v", err.Error())
		os.Exit(1)
	}

	return sdc
}
