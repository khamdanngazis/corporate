package transport

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	dt "../datastruct"
	ex "../error"
	logger "../logging"
	"../services"
	"github.com/go-kit/kit/endpoint"
)

// JWTDecodeRequest : request param for queue list using JSON format place in body
func JWTDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request dt.JWTJSONRequest

	var body []byte

	//decode request body
	body, err := ioutil.ReadAll(r.Body)
	logger.Logf("JWTDecodeRequest : %s", string(body[:]))
	if err != nil {
		return ex.Errorc(dt.ErrInvalidFormat).Rem("Unable to read request body"), nil
	}

	if err = json.Unmarshal(body, &request); err != nil {
		return ex.Error(err, dt.ErrInvalidFormat).Rem("Failed decoding json message"), nil
	}

	return request, nil
}

// JWTEncodeResponse : response using JSON format
func JWTEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	var body []byte
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	body, err := json.Marshal(&response)
	logger.Logf("JWTEncodeResponse : %s", string(body[:]))

	if err != nil {
		return err
	}

	//w.Header().Set("X-Checksum", cm.Cksum(body))

	var e = response.(dt.JWTJSONResponse).ResponseCode

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

// JWTEndpoint call Queue List
func JWTEndpoint(svc services.JWTServices) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if req, ok := request.(dt.JWTJSONRequest); ok {

			return svc.JWT(ctx, req), nil

		}
		logger.Error("Unhandled error occured: request is in unknown format")
		return dt.JWTJSONResponse{ResponseCode: dt.ErrOthers}, nil
	}
}
