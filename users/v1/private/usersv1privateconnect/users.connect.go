// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: users/v1/private/users.proto

package usersv1privateconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	private "github.com/keegancraigcruickshank/volley-api-connect/users/v1/private"
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
	// PrivateUsersServiceName is the fully-qualified name of the PrivateUsersService service.
	PrivateUsersServiceName = "users.v1.private.PrivateUsersService"
)

// PrivateUsersServiceClient is a client for the users.v1.private.PrivateUsersService service.
type PrivateUsersServiceClient interface {
	GetMe(context.Context, *connect_go.Request[private.GetMeRequest]) (*connect_go.Response[private.GetMeResponse], error)
}

// NewPrivateUsersServiceClient constructs a client for the users.v1.private.PrivateUsersService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPrivateUsersServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) PrivateUsersServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &privateUsersServiceClient{
		getMe: connect_go.NewClient[private.GetMeRequest, private.GetMeResponse](
			httpClient,
			baseURL+"/users.v1.private.PrivateUsersService/GetMe",
			opts...,
		),
	}
}

// privateUsersServiceClient implements PrivateUsersServiceClient.
type privateUsersServiceClient struct {
	getMe *connect_go.Client[private.GetMeRequest, private.GetMeResponse]
}

// GetMe calls users.v1.private.PrivateUsersService.GetMe.
func (c *privateUsersServiceClient) GetMe(ctx context.Context, req *connect_go.Request[private.GetMeRequest]) (*connect_go.Response[private.GetMeResponse], error) {
	return c.getMe.CallUnary(ctx, req)
}

// PrivateUsersServiceHandler is an implementation of the users.v1.private.PrivateUsersService
// service.
type PrivateUsersServiceHandler interface {
	GetMe(context.Context, *connect_go.Request[private.GetMeRequest]) (*connect_go.Response[private.GetMeResponse], error)
}

// NewPrivateUsersServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPrivateUsersServiceHandler(svc PrivateUsersServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/users.v1.private.PrivateUsersService/GetMe", connect_go.NewUnaryHandler(
		"/users.v1.private.PrivateUsersService/GetMe",
		svc.GetMe,
		opts...,
	))
	return "/users.v1.private.PrivateUsersService/", mux
}

// UnimplementedPrivateUsersServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPrivateUsersServiceHandler struct{}

func (UnimplementedPrivateUsersServiceHandler) GetMe(context.Context, *connect_go.Request[private.GetMeRequest]) (*connect_go.Response[private.GetMeResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("users.v1.private.PrivateUsersService.GetMe is not implemented"))
}