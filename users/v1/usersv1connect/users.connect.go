// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: users/v1/users.proto

package usersv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/keegancraigcruickshank/volley-api-connect/users/v1"
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
	// UsersServiceName is the fully-qualified name of the UsersService service.
	UsersServiceName = "users.v1.UsersService"
)

// UsersServiceClient is a client for the users.v1.UsersService service.
type UsersServiceClient interface {
	GetMe(context.Context, *connect_go.Request[v1.GetMeRequest]) (*connect_go.Response[v1.GetMeResponse], error)
}

// NewUsersServiceClient constructs a client for the users.v1.UsersService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUsersServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) UsersServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &usersServiceClient{
		getMe: connect_go.NewClient[v1.GetMeRequest, v1.GetMeResponse](
			httpClient,
			baseURL+"/users.v1.UsersService/GetMe",
			opts...,
		),
	}
}

// usersServiceClient implements UsersServiceClient.
type usersServiceClient struct {
	getMe *connect_go.Client[v1.GetMeRequest, v1.GetMeResponse]
}

// GetMe calls users.v1.UsersService.GetMe.
func (c *usersServiceClient) GetMe(ctx context.Context, req *connect_go.Request[v1.GetMeRequest]) (*connect_go.Response[v1.GetMeResponse], error) {
	return c.getMe.CallUnary(ctx, req)
}

// UsersServiceHandler is an implementation of the users.v1.UsersService service.
type UsersServiceHandler interface {
	GetMe(context.Context, *connect_go.Request[v1.GetMeRequest]) (*connect_go.Response[v1.GetMeResponse], error)
}

// NewUsersServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUsersServiceHandler(svc UsersServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/users.v1.UsersService/GetMe", connect_go.NewUnaryHandler(
		"/users.v1.UsersService/GetMe",
		svc.GetMe,
		opts...,
	))
	return "/users.v1.UsersService/", mux
}

// UnimplementedUsersServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUsersServiceHandler struct{}

func (UnimplementedUsersServiceHandler) GetMe(context.Context, *connect_go.Request[v1.GetMeRequest]) (*connect_go.Response[v1.GetMeResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("users.v1.UsersService.GetMe is not implemented"))
}
