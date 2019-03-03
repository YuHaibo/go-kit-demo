# Go-Kit-Demo: Add Service

Microservice demo using Go, go-kit, gRPC. Based on [go-kit example: addsvc](https://github.com/go-kit/kit/tree/master/examples/addsvc).

## Usage

Run server using `main.go` file:
```bash
$ go run main.go

//HTTP Port: 8890
//gPRC Port: 8891
```

HTTP requests:
```bash
$ ./client/http/curl.sh
```

gRPC client:
```bash
$ go run client/main.go
```
