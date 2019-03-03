package main

import (
	"context"
	"flag"
	grpcclient "go-kit-demo/client/grpc"
	"go-kit-demo/server"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	gRPCAddr := flag.String("gRPC", ":8891", "gRPC client")
	flag.Parse()

	conn, err := grpc.Dial(
		*gRPCAddr, grpc.WithInsecure(),
		grpc.WithTimeout(time.Second),
	)

	if err != nil {
		log.Fatalln("gRPC dial error:", err)
	}
	defer conn.Close()

	addService := grpcclient.New(conn)

	println("Sum Response:")
	sum(context.Background(), addService, 11111, 22222)

	println("Concat Response:")
	concat(context.Background(), addService, "11111", "22222")
}

func sum(ctx context.Context, svc server.AddService, a, b int) {
	output := svc.Sum(ctx, a, b)
	println(output)
}

func concat(ctx context.Context, svc server.AddService, a, b string) {
	output := svc.Concat(ctx, a, b)
	println(output)
}
