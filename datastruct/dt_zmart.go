package datastruct

//ZSmartPayloadSON is structure for ticket on QueueListJSONResponse
type ZSmartPayloadSON struct {
	PersonalInfo ZSmartProfileJSON `json:"personalInfo,omitempty"`
}

//ZSmartProfileJSON is structure for ticket on QueueListJSONResponse
type ZSmartProfileJSON struct {
	CustomerName string `json:"customerName,omitempty"`
}

//ZSmartJSONResponse is structure for service Response for Queue List
type ZSmartCustomerJSONResponse struct {
	IsSuccess  string           `json:"isSuccess,omitempty"`
	StatusCode string           `json:"statusCode,omitempty"`
	StatusMsg  string           ` json:"statusMsg,omitempty"`
	PayLoad    ZSmartPayloadSON `json:"payload,omitempty"`
}
