package services

import (
	"context"

	dt "../datastruct"
	lib "../lib"
	"../logging"
)

// CorporateRequiredInfoServices provides operations for endpoint
type CorporateRequiredInfoServices interface {
	CorporateRequiredInfo(context.Context, dt.CorporateRequiredInfoJSONRequest, lib.DbConnection) dt.CorporateRequiredInfoJSONResponse
}

// CorporateRequiredInfoService is a concrete implementation of QueueServices
type CorporateRequiredInfoService struct{}

//Get Category list
func GetRequiredInfo(dbConn lib.DbConnection, standardCode string) ([]dt.RequiredInfo, error) {
	var key string
	var value string
	var tipe string

	var requiredInfo dt.RequiredInfo
	var aRequiredInfo []dt.RequiredInfo

	var retVal, e = dbConn.Query("GetRequiredInfo", standardCode)
	if e != nil {
		logging.Logf("query GetRequiredInfo error : %s", e)
		return nil, e
	}
	defer retVal.Close()
	for retVal.Next() {
		if e = retVal.Scan(&key, &value, &tipe); e != nil {
			logging.Logf("scan GetRequiredInfoLevel1 error : %s", e)
			return nil, e
		}
		requiredInfo.Key = key
		requiredInfo.Value = value
		requiredInfo.Type = tipe
		aRequiredInfo = append(aRequiredInfo, requiredInfo)
	}

	return aRequiredInfo, nil
}

// CorporateRequiredInfo service
func (CorporateRequiredInfoService) CorporateRequiredInfo(ctx context.Context, req dt.CorporateRequiredInfoJSONRequest, dbConn lib.DbConnection) dt.CorporateRequiredInfoJSONResponse {

	RequiredInfo, qlErr := GetRequiredInfo(dbConn, req.StandardCode)
	if qlErr != nil {
		logging.Logf("%s Failed to get Requied info", qlErr)
		return dt.CorporateRequiredInfoJSONResponse{
			ResponseCode: dt.ErrGetRequiredInfoFailed,
			ResponseDesc: dt.DescGetRequiredInfoFailed,
		}
	}
	return dt.CorporateRequiredInfoJSONResponse{
		RequiredInfo: RequiredInfo,
		ResponseCode: dt.ErrSuccess,
		ResponseDesc: dt.DescSuccess,
	}

}
