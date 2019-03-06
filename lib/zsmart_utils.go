package lib

import (
	"encoding/json"
	"fmt"

	conf "../config"
	dt "../datastruct"
	logging "../logging"
)

//GetCustomerNameGrade is to get customer name and grade from mdn by accessing zsmart api
func GetCustomerNameGrade(mdn string) (string, error) {
	var zSmartData dt.ZSmartCustomerJSONResponse
	var custName string
	url := conf.Param.ZSmartPath + "/api/v1/inquiry/opGetSubscriberInfo"

	jsonStr := fmt.Sprintf("{\r\n\t\"mdn\": \"%s\"\r\n}\r\n", mdn)

	respJSON, e := SendRequestZsmart("POST", url, jsonStr)
	sRespJSON := fmt.Sprintf(string(respJSON[:]))
	if e != nil {
		logging.Logf("%s", e)
		logging.Logf(sRespJSON)
		logging.Logf("Failed to get customer name from zsmart, Send request Error")
		return "", e
	}
	e1 := json.Unmarshal(respJSON, &zSmartData)

	if e1 != nil {
		logging.Logf("%s", e)
		logging.Logf(sRespJSON)
		logging.Logf("Failed to get customer name from zsmart, Unmarshall Error")
		return "", e1
	}
	if zSmartData.StatusMsg == "This MDN Grade is VIP or VVIP" {
		custName = SearchJSON(respJSON, "customerName")
		if custName == "DUMMY" {
			custName = mdn
		}
	} else if zSmartData.IsSuccess == "0" {
		logging.Logf(sRespJSON)
		logging.Logf("Failed to get customer name from zsmart, No MDN")
		return "", e1
	} else {
		custName = SearchJSON(respJSON, "customerName")
		if custName == "DUMMY" {
			custName = mdn
		}
	}
	return custName, nil
}
