package datastruct

//CorporateCreateTicketJSONRequest is structure for service Request for Queue Action
type CorporateCreateTicketJSONRequest struct {
	StandardCode string     `json:"standard_code,omitempty"`
	Mdn          string     `json:"mdn,omitempty"`
	RequiredInfo []Required `json:"required_info,omitempty"`
	Remark       string     `json:"reamrk,omitempty"`
}

//CorporateCreateTicketJSONResponse is structure for service Response for Queue Action
type CorporateCreateTicketJSONResponse struct {
	ResponseCode int    `json:"responseCode"`
	ResponseDesc string `json:"responseDesc,omitempty"`
}

type Required struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}
