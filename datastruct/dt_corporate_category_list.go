package datastruct

//CorporateCategoryListJSONRequest is structure for service Request for Queue Action
type CorporateCategoryListJSONRequest struct {
	Level     int    `json:"level"`
	KeyLevel1 string `json:"keylevel1,omitempty"`
	KeyLevel2 string `json:"keylevel2,omitempty"`
	KeyLevel3 string `json:"keylevel3,omitempty"`
}

//CorporateCategoryListJSONResponse is structure for service Response for Queue Action
type CorporateCategoryListJSONResponse struct {
	CategoryList []CategoryList `json:"category_list,omitempty"`
	ResponseCode int            `json:"responseCode"`
	ResponseDesc string         `json:"responseDesc,omitempty"`
}

//Category is structure for DetailsCategoryList on CorporateCategoryListJSONResponse
type CategoryList struct {
	Key      string `json:"key,omitempty"`
	Value    string `json:"value,omitempty"`
	HasChild bool   `json:"HasChild"`
}
