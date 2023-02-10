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
