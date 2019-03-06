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

// CorporateCreateTicket : request param for queue list using JSON format place in body
func CorporateCreateTicketDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dt.CorporateCreateTicketJSONRequest
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
	logger.Logf("CorporateCreateTicket : %s", string(body[:]))
	if err != nil {
		return ex.Errorc(dt.ErrInvalidFormat).Rem("Unable to read request body"), nil
	}

	if err = json.Unmarshal(body, &req); err != nil {
		return ex.Error(err, dt.ErrInvalidFormat).Rem("Failed decoding json message"), nil
	}

	return req, nil
}

// CorporateCreateTicketEncodeResponse : response using JSON format
func CorporateCreateTicketEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	var body []byte
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
	body, err := json.Marshal(&response)
	logger.Logf("CorporateCreateTicketEncodeResponse : %s", string(body[:]))

	if err != nil {
		return err
	}

	//w.Header().Set("X-Checksum", cm.Cksum(body))

	var e = response.(dt.CorporateCreateTicketJSONResponse).ResponseCode

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

// CorporateCreateTicketEndpoint call Queue List
func CorporateCreateTicketEndpoint(svc services.CorporateCreateTicketServices, dbConn lib.DbConnection) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if req, ok := request.(dt.CorporateCreateTicketJSONRequest); ok {

			return svc.CorporateCreateTicket(ctx, req, dbConn), nil

		}
		if _, ok := request.(dt.TokenValidation); ok {
			logger.Error("Invalid Token")
			return dt.CorporateCreateTicketJSONResponse{ResponseCode: dt.ErrInvalidToken, ResponseDesc: dt.DescInvalidToken}, nil
		}
		logger.Error("Unhandled error occured: request is in unknown format")
		return dt.CorporateCreateTicketJSONResponse{ResponseCode: dt.ErrOthers}, nil
	}
}
