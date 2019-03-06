package services

import (
	"context"

	dt "../datastruct"
	lib "../lib"
	"../logging"
)

// JWTServices provides operations for endpoint
type JWTServices interface {
	JWT(context.Context, dt.JWTJSONRequest) dt.JWTJSONResponse
}

// JWTService is a concrete implementation of QueueServices
type JWTService struct{}

// JWT service for Call queue
func (JWTService) JWT(ctx context.Context, req dt.JWTJSONRequest) dt.JWTJSONResponse {
	logging.Log("Os Type " + req.OsType)
	Auth := lib.TokenAuth(req.Mdn, req.KeyAuth, req.TimeStamp)
	if Auth == false {
		return dt.JWTJSONResponse{
			ResponseCode: dt.ErrGetToken,
			ResponseDesc: dt.DescGetToken,
		}
	}

	token, err := lib.GetToken()
	if err != nil {
		return dt.JWTJSONResponse{
			ResponseCode: dt.ErrGetToken,
			ResponseDesc: dt.DescGetToken,
		}
	}

	custMDN := lib.MDNNormalisation(req.Mdn)
	custName, _ := lib.GetCustomerNameGrade(custMDN)

	return dt.JWTJSONResponse{
		ResponseCode: dt.ErrSuccess,
		ResponseDesc: dt.DescSuccess,
		Token:        token,
		CustName:     custName,
	}
}
