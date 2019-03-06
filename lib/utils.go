package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	conf "../config"
	logging "../logging"
)

//send request to another aplication
func SendRequestZsmart(method string, url string, jsonStr string) ([]byte, error) {
	payload := strings.NewReader(jsonStr)

	req, _ := http.NewRequest(method, url, payload)
	req.Header.Add("requestID", conf.Param.ZSmartRequestID)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("username", conf.Param.ZSmartUserName)
	req.Header.Add("password", conf.Param.ZSmartPassword)
	req.Header.Add("Cache-Control", "no-cache")

	res, e := http.DefaultClient.Do(req)
	if e != nil {
		logging.Logf("%s", e)
		logging.Logf("Failed to send request to zsmart")
		return nil, e
	}

	defer res.Body.Close()
	body, e := ioutil.ReadAll(res.Body)
	if e != nil {
		logging.Logf("%s", e)
		logging.Logf("Failed to send request to zsmart")
		return nil, e
	}

	return body, e
}

func MDNNormalisation(mdn string) string {
	var retVal string
	if mdn[0:3] == "628" {
		retVal = mdn
	} else if mdn[0:2] == "08" {
		retVal = strings.Replace(mdn, "08", "628", 1)
	} else if mdn[0:1] == "8" {
		retVal = strings.Replace(mdn, "8", "628", 1)
	} else if mdn[0:1] == "+" {
		retVal = strings.Replace(mdn, "+", "", 1)
		retVal = MDNNormalisation(retVal)

	}
	return retVal
}

//MDNCustomerFormat is to normalize mdn number to 088xxxxxxxx
func MDNCustomerFormat(mdn string) string {
	var retVal string
	if mdn[0:2] == "08" {
		retVal = mdn
	} else if mdn[0:3] == "628" {
		retVal = strings.Replace(mdn, "628", "08", 1)
	} else if mdn[0:1] == "8" {
		retVal = strings.Replace(mdn, "8", "08", 1)
	} else if mdn[0:1] == "+" {
		retVal = strings.Replace(mdn, "+", "", 1)
		retVal = MDNCustomerFormat(retVal)

	}
	return retVal
}

//SearchJSON for accessing raw json without marshall it
func SearchJSON(jsonStr []byte, searchStr string) string {
	sJSONStr := fmt.Sprintf(string(jsonStr[:]))
	regexStr := fmt.Sprintf("(?:\"%s\"\\s*:)(.*?)(?:[,|}])", searchStr)
	r, _ := regexp.Compile(regexStr)
	//fmt.Println(r.FindString("{\"isSuccess\":\"1\",\"statusCode\":\"00\",\"statusMsg\":\"Success\"}"))

	sJSONElement := fmt.Sprintf(r.FindString(sJSONStr))
	regexStr = "(?::)(.*?)(?:[,|}])"
	r, _ = regexp.Compile(regexStr)
	sJSONValue := r.FindString(sJSONElement)
	sJSONValue = strings.Replace(sJSONValue, "\"", "", -1)
	sJSONValue = strings.TrimLeft(sJSONValue, ":")
	sJSONValue = strings.TrimRight(sJSONValue, "},")
	return sJSONValue
}
