// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: users/v1/public/users.proto

package usersv1publicconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	public "github.com/keegancraigcruickshank/volley-api-connect/users/v1/public"
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
	PrivateUsersServiceName = "users.v1.public.PrivateUsersService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PrivateUsersServiceGetTokenProcedure is the fully-qualified name of the PrivateUsersService's
	// GetToken RPC.
	PrivateUsersServiceGetTokenProcedure = "/users.v1.public.PrivateUsersService/GetToken"
)

// PrivateUsersServiceClient is a client for the users.v1.public.PrivateUsersService service.
type PrivateUsersServiceClient interface {
	GetToken(context.Context, *connect_go.Request[public.GetTokenRequest]) (*connect_go.Response[public.GetTokenResponse], error)
}

// NewPrivateUsersServiceClient constructs a client for the users.v1.public.PrivateUsersService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPrivateUsersServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) PrivateUsersServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &privateUsersServiceClient{
		getToken: connect_go.NewClient[public.GetTokenRequest, public.GetTokenResponse](
			httpClient,
			baseURL+PrivateUsersServiceGetTokenProcedure,
			opts...,
		),
	}
}

// privateUsersServiceClient implements PrivateUsersServiceClient.
type privateUsersServiceClient struct {
	getToken *connect_go.Client[public.GetTokenRequest, public.GetTokenResponse]
}

// GetToken calls users.v1.public.PrivateUsersService.GetToken.
func (c *privateUsersServiceClient) GetToken(ctx context.Context, req *connect_go.Request[public.GetTokenRequest]) (*connect_go.Response[public.GetTokenResponse], error) {
	return c.getToken.CallUnary(ctx, req)
}

// PrivateUsersServiceHandler is an implementation of the users.v1.public.PrivateUsersService
// service.
type PrivateUsersServiceHandler interface {
	GetToken(context.Context, *connect_go.Request[public.GetTokenRequest]) (*connect_go.Response[public.GetTokenResponse], error)
}

// NewPrivateUsersServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPrivateUsersServiceHandler(svc PrivateUsersServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	privateUsersServiceGetTokenHandler := connect_go.NewUnaryHandler(
		PrivateUsersServiceGetTokenProcedure,
		svc.GetToken,
		opts...,
	)
	return "/users.v1.public.PrivateUsersService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PrivateUsersServiceGetTokenProcedure:
			privateUsersServiceGetTokenHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPrivateUsersServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPrivateUsersServiceHandler struct{}

func (UnimplementedPrivateUsersServiceHandler) GetToken(context.Context, *connect_go.Request[public.GetTokenRequest]) (*connect_go.Response[public.GetTokenResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("users.v1.public.PrivateUsersService.GetToken is not implemented"))
}
