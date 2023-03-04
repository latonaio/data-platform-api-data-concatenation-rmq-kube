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

func ConvertToOrdersEDIForSMEsSDC(data []byte) OrdersEDIForSMEsSDC {
	sdc := OrdersEDIForSMEsSDC{}
	err := json.Unmarshal(data, &sdc)
	if err != nil {
		fmt.Printf("input data marshal error :%#v", err.Error())
		os.Exit(1)
	}

	return sdc
}
