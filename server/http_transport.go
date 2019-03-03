package server

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

func MakeHTTPHandler(endpoints Endpoints) http.Handler {
	m := http.NewServeMux()
	m.Handle("/sum",
		httptransport.NewServer(
			endpoints.SumEndpoint,
			DecodeHTTPSumRequest,
			EncodeHTTPResponse,
		))

	m.Handle("/concat",
		httptransport.NewServer(
			endpoints.ConcatEndpoint,
			DecodeHTTPConcatRequest,
			EncodeHTTPResponse,
		))

	return m
}

func DecodeHTTPSumRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request sumRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeHTTPConcatRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request concatRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeHTTPResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
