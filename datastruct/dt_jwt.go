package datastruct

//QueueBreakJSONRequest is structure for service Request for Queue Action
type JWTJSONRequest struct {
	KeyAuth   string `json:"key,omitempty"`
	Mdn       string `json:"mdn,omitempty"`
	TimeStamp string `json:"timeStamp,omitempty"`
	OsType    string `json:"os,omitempty"`
}

//QueueBreakJSONResponse is structure for service Response for Queue Action
type JWTJSONResponse struct {
	ResponseCode int    `json:"responseCode"`
	ResponseDesc string `json:"responseDesc,omitempty"`
	Token        string `json:"token,omitempty"`
	CustName     string `json:"custName,omitempty"`
}

type TokenValidation struct {
	Validation bool `json:"validation,omitempty"`
}
