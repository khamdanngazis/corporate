package transport

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	conf "../config"
	dt "../datastruct"
	ex "../error"
	lib "../lib"
	logger "../logging"
	"../services"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/go-kit/kit/endpoint"
)

// CorporateCategoryList : request param for queue list using JSON format place in body
func CorporateCategoryListDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dt.CorporateCategoryListJSONRequest
	var tokenValidation dt.TokenValidation

	//token Authorization
	var mySigningKey = []byte(conf.Param.TokenAuth)
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		})
	if err != nil {
		return tokenValidation, nil
	}
	if !token.Valid {
		return tokenValidation, nil
	}
	var body []byte

	//decode request body
	body, err = ioutil.ReadAll(r.Body)
	logger.Logf("CorporateCategoryList : %s", string(body[:]))
	if err != nil {
		return ex.Errorc(dt.ErrInvalidFormat).Rem("Unable to read request body"), nil
	}

	if err = json.Unmarshal(body, &req); err != nil {
		return ex.Error(err, dt.ErrInvalidFormat).Rem("Failed decoding json message"), nil
	}

	return req, nil
}

// CorporateCategoryListEncodeResponse : response using JSON format
func CorporateCategoryListEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	var body []byte
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
	body, err := json.Marshal(&response)
	logger.Logf("CorporateCategoryListEncodeResponse : %s", string(body[:]))

	if err != nil {
		return err
	}

	//w.Header().Set("X-Checksum", cm.Cksum(body))

	var e = response.(dt.CorporateCategoryListJSONResponse).ResponseCode

	if e <= dt.HeaderStatusOk {
		w.WriteHeader(http.StatusOK)
	} else if e <= dt.StatusBadRequest {
		w.WriteHeader(http.StatusBadRequest)
	} else if e <= 998 {
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	_, err = w.Write(body)

	return err
}

// CorporateCategoryListEndpoint call Queue List
func CorporateCategoryListEndpoint(svc services.CorporateCategoryListServices, dbConn lib.DbConnection) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if req, ok := request.(dt.CorporateCategoryListJSONRequest); ok {

			return svc.CorporateCategoryList(ctx, req, dbConn), nil

		}
		if _, ok := request.(dt.TokenValidation); ok {
			logger.Error("Invalid Token")
			return dt.CorporateCategoryListJSONResponse{ResponseCode: dt.ErrInvalidToken, ResponseDesc: dt.DescInvalidToken}, nil
		}
		logger.Error("Unhandled error occured: request is in unknown format")
		return dt.CorporateCategoryListJSONResponse{ResponseCode: dt.ErrOthers}, nil
	}
}
