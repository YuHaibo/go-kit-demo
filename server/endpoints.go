package server

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

//all endpoints required by AddService.
type Endpoints struct {
	SumEndpoint		endpoint.Endpoint
	ConcatEndpoint	endpoint.Endpoint
}

type sumRequest struct {
	A int
	B int
}

type sumResponse struct {
	V int
}

// MakeSumEndpoint returns an endpoint that invokes Sum on the AddService
// for server
func MakeSumEndpoint(svc AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(sumRequest)
		v := svc.Sum(ctx, req.A, req.B)
		return sumResponse{v}, nil
	}
}

// Sum implements AddService
//for client
func (e Endpoints) Sum(ctx context.Context, a, b int) int {
	req := sumRequest{A:a, B:b}
	res, err := e.SumEndpoint(ctx, req)
	if err != nil {
		return sumResponse{0}.V
	}
	return res.(sumResponse).V
}

type concatRequest struct {
	A string
	B string
}

type concatResponse struct {
	V string
}

// MakeConcatEndpoint returns an endpoint that invokes Sum on the AddService
// for server
func MakeConcatEndpoint(svc AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(concatRequest)
		v := svc.Concat(ctx, req.A, req.B)
		return concatResponse{v}, nil
	}
}

// Concat implements AddService
//for client
func (e Endpoints) Concat(ctx context.Context, a, b string) string {
	req := concatRequest{A:a, B:b}
	res, err := e.ConcatEndpoint(ctx, req)
	if err != nil {
		return concatResponse{"error"}.V
	}
	return res.(concatResponse).V
}
