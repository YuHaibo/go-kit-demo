package server

//Server-side bindings for the gRPC transport

import (
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"go-kit-demo/pb"
)

//returns a set of handlers available as a gRPC AddServer.
func MakeGRPCServer(endpoints Endpoints) pb.AddServer {
	return &grpcServer{
		sum: grpctransport.NewServer(
			endpoints.SumEndpoint,
			DecodeGRPCSumRequest,
			EncodeGRPCSumResponse,
		),
		concat: grpctransport.NewServer(
			endpoints.ConcatEndpoint,
			DecodeGRPCConcatRequest,
			EncodeGRPCConcatResponse,
		),
	}
}

type grpcServer struct {
	sum 	grpctransport.Handler
	concat 	grpctransport.Handler
}

func (s *grpcServer) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumReply, error) {
	_, resp, err := s.sum.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.SumReply), nil
}

func (s *grpcServer) Concat(ctx context.Context, req *pb.ConcatRequest) (*pb.ConcatReply, error) {
	_, resp, err := s.concat.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ConcatReply), nil
}

//Encode & Decode Func
//https://github.com/go-kit/kit/blob/master/transport/grpc/encode_decode.go

// DecodeGRPCSumRequest is a transport/grpc.DecodeRequestFunc
// for server
func DecodeGRPCSumRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.SumRequest)
	return sumRequest{int(req.A), int(req.B)}, nil
}

// EncodeGRPCSumRequest is a transport/grpc.EncodeRequestFunc
// for client
func EncodeGRPCSumRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(sumRequest)
	return &pb.SumRequest{A: int64(req.A), B: int64(req.B)}, nil
}

// EncodeGRPCSumResponse is a transport/grpc.EncodeResponseFunc
// for server
func EncodeGRPCSumResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(sumResponse)
	return &pb.SumReply{V: int64(resp.V)}, nil
}

// DecodeGRPCSumResponse is a transport/grpc.DecodeResponseFunc
// for client
func DecodeGRPCSumResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.SumReply)
	return sumResponse{int(resp.V)}, nil
}

func DecodeGRPCConcatRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ConcatRequest)
	return concatRequest{req.A, req.B}, nil
}

func EncodeGRPCConcatRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(concatRequest)
	return &pb.ConcatRequest{A: req.A, B: req.B}, nil
}

func EncodeGRPCConcatResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(concatResponse)
	return &pb.ConcatReply{V: resp.V}, nil
}

func DecodeGRPCConcatResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ConcatReply)
	return concatResponse{resp.V}, nil
}