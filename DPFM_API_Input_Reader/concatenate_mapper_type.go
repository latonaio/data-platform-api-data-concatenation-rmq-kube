package dpfm_api_input_reader

type ConcatenateMapper struct {
	ConcatenateMapperID int     `json:"ConcatenateMapperID"`
	ServiceLabel        string  `json:"ServiceLabel"`
	APIQueueName        string  `json:"APIQueueName"`
	BaseAPIName         string  `json:"BaseAPIName"`
	ConcatenateAPIName1 string  `json:"ConcatenateAPIName1"`
	ConcatenateAPIName2 string  `json:"ConcatenateAPIName2"`
	ConcatenateAPIName3 string  `json:"ConcatenateAPIName3"`
	ConcatenateAPIName4 string  `json:"ConcatenateAPIName4"`
	ConcatenateAPIName5 string  `json:"ConcatenateAPIName5"`
	ConcatenateKey1     *string `json:"ConcatenateKey1"`
	ConcatenateKey2     *string `json:"ConcatenateKey2"`
	ConcatenateKey3     *string `json:"ConcatenateKey3"`
	ConcatenateKey4     *string `json:"ConcatenateKey4"`
	ConcatenateKey5     *string `json:"ConcatenateKey5"`
}
