// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: subscriptions/v1/private/subscriptions.proto

package subscriptionsv1privateconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	private "github.com/keegancraigcruickshank/volley-api-connect/subscriptions/v1/private"
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
	// PrivateSubscriptionsServiceName is the fully-qualified name of the PrivateSubscriptionsService
	// service.
	PrivateSubscriptionsServiceName = "subscriptions.v1.private.PrivateSubscriptionsService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PrivateSubscriptionsServiceGetOrganisationPlanProcedure is the fully-qualified name of the
	// PrivateSubscriptionsService's GetOrganisationPlan RPC.
	PrivateSubscriptionsServiceGetOrganisationPlanProcedure = "/subscriptions.v1.private.PrivateSubscriptionsService/GetOrganisationPlan"
)

// PrivateSubscriptionsServiceClient is a client for the
// subscriptions.v1.private.PrivateSubscriptionsService service.
type PrivateSubscriptionsServiceClient interface {
	GetOrganisationPlan(context.Context, *connect_go.Request[private.GetOrganisationPlanRequest]) (*connect_go.Response[private.GetOrganisationPlanResponse], error)
}

// NewPrivateSubscriptionsServiceClient constructs a client for the
// subscriptions.v1.private.PrivateSubscriptionsService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPrivateSubscriptionsServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) PrivateSubscriptionsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &privateSubscriptionsServiceClient{
		getOrganisationPlan: connect_go.NewClient[private.GetOrganisationPlanRequest, private.GetOrganisationPlanResponse](
			httpClient,
			baseURL+PrivateSubscriptionsServiceGetOrganisationPlanProcedure,
			opts...,
		),
	}
}

// privateSubscriptionsServiceClient implements PrivateSubscriptionsServiceClient.
type privateSubscriptionsServiceClient struct {
	getOrganisationPlan *connect_go.Client[private.GetOrganisationPlanRequest, private.GetOrganisationPlanResponse]
}

// GetOrganisationPlan calls
// subscriptions.v1.private.PrivateSubscriptionsService.GetOrganisationPlan.
func (c *privateSubscriptionsServiceClient) GetOrganisationPlan(ctx context.Context, req *connect_go.Request[private.GetOrganisationPlanRequest]) (*connect_go.Response[private.GetOrganisationPlanResponse], error) {
	return c.getOrganisationPlan.CallUnary(ctx, req)
}

// PrivateSubscriptionsServiceHandler is an implementation of the
// subscriptions.v1.private.PrivateSubscriptionsService service.
type PrivateSubscriptionsServiceHandler interface {
	GetOrganisationPlan(context.Context, *connect_go.Request[private.GetOrganisationPlanRequest]) (*connect_go.Response[private.GetOrganisationPlanResponse], error)
}

// NewPrivateSubscriptionsServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPrivateSubscriptionsServiceHandler(svc PrivateSubscriptionsServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	privateSubscriptionsServiceGetOrganisationPlanHandler := connect_go.NewUnaryHandler(
		PrivateSubscriptionsServiceGetOrganisationPlanProcedure,
		svc.GetOrganisationPlan,
		opts...,
	)
	return "/subscriptions.v1.private.PrivateSubscriptionsService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PrivateSubscriptionsServiceGetOrganisationPlanProcedure:
			privateSubscriptionsServiceGetOrganisationPlanHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPrivateSubscriptionsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPrivateSubscriptionsServiceHandler struct{}

func (UnimplementedPrivateSubscriptionsServiceHandler) GetOrganisationPlan(context.Context, *connect_go.Request[private.GetOrganisationPlanRequest]) (*connect_go.Response[private.GetOrganisationPlanResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("subscriptions.v1.private.PrivateSubscriptionsService.GetOrganisationPlan is not implemented"))
}
