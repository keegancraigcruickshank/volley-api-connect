// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: busdriver/v1/busdriver.proto

package busdriverv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/keegancraigcruickshank/volley-api-connect/busdriver/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// BusDriverServiceName is the fully-qualified name of the BusDriverService service.
	BusDriverServiceName = "busdriver.v1.BusDriverService"
)

// BusDriverServiceClient is a client for the busdriver.v1.BusDriverService service.
type BusDriverServiceClient interface {
	GetMe(context.Context, *connect_go.Request[v1.GetMeRequest]) (*connect_go.Response[v1.GetMeResponse], error)
}

// NewBusDriverServiceClient constructs a client for the busdriver.v1.BusDriverService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewBusDriverServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) BusDriverServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &busDriverServiceClient{
		getMe: connect_go.NewClient[v1.GetMeRequest, v1.GetMeResponse](
			httpClient,
			baseURL+"/busdriver.v1.BusDriverService/GetMe",
			opts...,
		),
	}
}

// busDriverServiceClient implements BusDriverServiceClient.
type busDriverServiceClient struct {
	getMe *connect_go.Client[v1.GetMeRequest, v1.GetMeResponse]
}

// GetMe calls busdriver.v1.BusDriverService.GetMe.
func (c *busDriverServiceClient) GetMe(ctx context.Context, req *connect_go.Request[v1.GetMeRequest]) (*connect_go.Response[v1.GetMeResponse], error) {
	return c.getMe.CallUnary(ctx, req)
}

// BusDriverServiceHandler is an implementation of the busdriver.v1.BusDriverService service.
type BusDriverServiceHandler interface {
	GetMe(context.Context, *connect_go.Request[v1.GetMeRequest]) (*connect_go.Response[v1.GetMeResponse], error)
}

// NewBusDriverServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewBusDriverServiceHandler(svc BusDriverServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/busdriver.v1.BusDriverService/GetMe", connect_go.NewUnaryHandler(
		"/busdriver.v1.BusDriverService/GetMe",
		svc.GetMe,
		opts...,
	))
	return "/busdriver.v1.BusDriverService/", mux
}

// UnimplementedBusDriverServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedBusDriverServiceHandler struct{}

func (UnimplementedBusDriverServiceHandler) GetMe(context.Context, *connect_go.Request[v1.GetMeRequest]) (*connect_go.Response[v1.GetMeResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("busdriver.v1.BusDriverService.GetMe is not implemented"))
}
