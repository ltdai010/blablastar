package data


type ResponseBool struct {
	Data bool `json:"data" xml:"data"`
	*Err
}

type ResponseCommonSingle struct {
	Data interface{} `json:"data" xml:"data"`
	*Err
}

type ResponseCommonArray struct {
	Data       interface{} `json:"data" xml:"data"`
	TotalCount int       `json:"total_count" xml:"total_count"`
	*Err
}
