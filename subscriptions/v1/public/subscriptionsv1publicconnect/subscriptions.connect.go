// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: subscriptions/v1/public/subscriptions.proto

package subscriptionsv1publicconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	public "github.com/keegancraigcruickshank/volley-api-connect/subscriptions/v1/public"
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
	// PublicSubscriptionsServiceName is the fully-qualified name of the PublicSubscriptionsService
	// service.
	PublicSubscriptionsServiceName = "subscriptions.v1.public.PublicSubscriptionsService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PublicSubscriptionsServiceGetSubscriptionPlansProcedure is the fully-qualified name of the
	// PublicSubscriptionsService's GetSubscriptionPlans RPC.
	PublicSubscriptionsServiceGetSubscriptionPlansProcedure = "/subscriptions.v1.public.PublicSubscriptionsService/GetSubscriptionPlans"
)

// PublicSubscriptionsServiceClient is a client for the
// subscriptions.v1.public.PublicSubscriptionsService service.
type PublicSubscriptionsServiceClient interface {
	GetSubscriptionPlans(context.Context, *connect_go.Request[public.GetSubscriptionPlansRequest]) (*connect_go.Response[public.GetSubscriptionPlansResponse], error)
}

// NewPublicSubscriptionsServiceClient constructs a client for the
// subscriptions.v1.public.PublicSubscriptionsService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPublicSubscriptionsServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) PublicSubscriptionsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &publicSubscriptionsServiceClient{
		getSubscriptionPlans: connect_go.NewClient[public.GetSubscriptionPlansRequest, public.GetSubscriptionPlansResponse](
			httpClient,
			baseURL+PublicSubscriptionsServiceGetSubscriptionPlansProcedure,
			opts...,
		),
	}
}

// publicSubscriptionsServiceClient implements PublicSubscriptionsServiceClient.
type publicSubscriptionsServiceClient struct {
	getSubscriptionPlans *connect_go.Client[public.GetSubscriptionPlansRequest, public.GetSubscriptionPlansResponse]
}

// GetSubscriptionPlans calls
// subscriptions.v1.public.PublicSubscriptionsService.GetSubscriptionPlans.
func (c *publicSubscriptionsServiceClient) GetSubscriptionPlans(ctx context.Context, req *connect_go.Request[public.GetSubscriptionPlansRequest]) (*connect_go.Response[public.GetSubscriptionPlansResponse], error) {
	return c.getSubscriptionPlans.CallUnary(ctx, req)
}

// PublicSubscriptionsServiceHandler is an implementation of the
// subscriptions.v1.public.PublicSubscriptionsService service.
type PublicSubscriptionsServiceHandler interface {
	GetSubscriptionPlans(context.Context, *connect_go.Request[public.GetSubscriptionPlansRequest]) (*connect_go.Response[public.GetSubscriptionPlansResponse], error)
}

// NewPublicSubscriptionsServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPublicSubscriptionsServiceHandler(svc PublicSubscriptionsServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	publicSubscriptionsServiceGetSubscriptionPlansHandler := connect_go.NewUnaryHandler(
		PublicSubscriptionsServiceGetSubscriptionPlansProcedure,
		svc.GetSubscriptionPlans,
		opts...,
	)
	return "/subscriptions.v1.public.PublicSubscriptionsService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PublicSubscriptionsServiceGetSubscriptionPlansProcedure:
			publicSubscriptionsServiceGetSubscriptionPlansHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPublicSubscriptionsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPublicSubscriptionsServiceHandler struct{}

func (UnimplementedPublicSubscriptionsServiceHandler) GetSubscriptionPlans(context.Context, *connect_go.Request[public.GetSubscriptionPlansRequest]) (*connect_go.Response[public.GetSubscriptionPlansResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("subscriptions.v1.public.PublicSubscriptionsService.GetSubscriptionPlans is not implemented"))
}