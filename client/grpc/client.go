package grpc

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"go-kit-demo/pb"
	"go-kit-demo/server"
	"google.golang.org/grpc"
)

func New(conn *grpc.ClientConn) server.AddService {
	sumEndpoint := grpctransport.NewClient(
		conn, "pb.Add", "Sum",
		server.EncodeGRPCSumRequest,
		server.DecodeGRPCSumResponse,
		pb.SumReply{},
	).Endpoint()

	concatEndpoint := grpctransport.NewClient(
		conn, "pb.Add", "Concat",
		server.EncodeGRPCConcatRequest,
		server.DecodeGRPCConcatResponse,
		pb.ConcatReply{},
	).Endpoint()

	return server.Endpoints{
		SumEndpoint:    sumEndpoint,
		ConcatEndpoint: concatEndpoint,
	}
}
