package datastruct

//CorporateRequiredInfoJSONRequest is structure for service Request for Queue Action
type CorporateRequiredInfoJSONRequest struct {
	StandardCode string `json:"standard_code,omitempty"`
}

//CorporateRequiredInfoJSONResponse is structure for service Response for Queue Action
type CorporateRequiredInfoJSONResponse struct {
	RequiredInfo []RequiredInfo `json:"required_info,omitempty"`
	ResponseCode int            `json:"responseCode"`
	ResponseDesc string         `json:"responseDesc,omitempty"`
}

type RequiredInfo struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
	Type  string `json:"type,omitempty"`
}
