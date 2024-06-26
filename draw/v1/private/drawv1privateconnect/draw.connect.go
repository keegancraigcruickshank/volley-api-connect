// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: draw/v1/private/draw.proto

package drawv1privateconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	private "github.com/keegancraigcruickshank/volley-api-connect/draw/v1/private"
	flexible_round_robin "github.com/keegancraigcruickshank/volley-api-connect/draw/v1/private/flexible-round-robin"
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
	// PrivateDrawServiceName is the fully-qualified name of the PrivateDrawService service.
	PrivateDrawServiceName = "draw.v1.private.PrivateDrawService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PrivateDrawServiceGetDrawProcedure is the fully-qualified name of the PrivateDrawService's
	// GetDraw RPC.
	PrivateDrawServiceGetDrawProcedure = "/draw.v1.private.PrivateDrawService/GetDraw"
	// PrivateDrawServiceDeleteDrawProcedure is the fully-qualified name of the PrivateDrawService's
	// DeleteDraw RPC.
	PrivateDrawServiceDeleteDrawProcedure = "/draw.v1.private.PrivateDrawService/DeleteDraw"
	// PrivateDrawServiceListDrawsProcedure is the fully-qualified name of the PrivateDrawService's
	// ListDraws RPC.
	PrivateDrawServiceListDrawsProcedure = "/draw.v1.private.PrivateDrawService/ListDraws"
	// PrivateDrawServiceCreateFlexibleRoundRobinDrawProcedure is the fully-qualified name of the
	// PrivateDrawService's CreateFlexibleRoundRobinDraw RPC.
	PrivateDrawServiceCreateFlexibleRoundRobinDrawProcedure = "/draw.v1.private.PrivateDrawService/CreateFlexibleRoundRobinDraw"
	// PrivateDrawServiceUpdateFlexibleRoundRobinDrawProcedure is the fully-qualified name of the
	// PrivateDrawService's UpdateFlexibleRoundRobinDraw RPC.
	PrivateDrawServiceUpdateFlexibleRoundRobinDrawProcedure = "/draw.v1.private.PrivateDrawService/UpdateFlexibleRoundRobinDraw"
	// PrivateDrawServiceGetFlexibleRoundRobinDrawRoundProcedure is the fully-qualified name of the
	// PrivateDrawService's GetFlexibleRoundRobinDrawRound RPC.
	PrivateDrawServiceGetFlexibleRoundRobinDrawRoundProcedure = "/draw.v1.private.PrivateDrawService/GetFlexibleRoundRobinDrawRound"
)

// PrivateDrawServiceClient is a client for the draw.v1.private.PrivateDrawService service.
type PrivateDrawServiceClient interface {
	// Common endpoints
	GetDraw(context.Context, *connect_go.Request[private.GetDrawRequest]) (*connect_go.Response[private.GetDrawResponse], error)
	DeleteDraw(context.Context, *connect_go.Request[private.DeleteDrawRequest]) (*connect_go.Response[private.DeleteDrawResponse], error)
	ListDraws(context.Context, *connect_go.Request[private.ListDrawsRequest]) (*connect_go.Response[private.ListDrawsResponse], error)
	// Flexible round robin
	CreateFlexibleRoundRobinDraw(context.Context, *connect_go.Request[flexible_round_robin.CreateFlexibleRoundRobinDrawRequest]) (*connect_go.Response[flexible_round_robin.CreateFlexibleRoundRobinDrawResponse], error)
	UpdateFlexibleRoundRobinDraw(context.Context, *connect_go.Request[flexible_round_robin.UpdateFlexibleRoundRobinDrawRequest]) (*connect_go.Response[flexible_round_robin.UpdateFlexibleRoundRobinDrawResponse], error)
	GetFlexibleRoundRobinDrawRound(context.Context, *connect_go.Request[flexible_round_robin.GetFlexibleRoundRobinDrawRoundRequest]) (*connect_go.Response[flexible_round_robin.GetFlexibleRoundRobinDrawRoundResponse], error)
}

// NewPrivateDrawServiceClient constructs a client for the draw.v1.private.PrivateDrawService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPrivateDrawServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) PrivateDrawServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &privateDrawServiceClient{
		getDraw: connect_go.NewClient[private.GetDrawRequest, private.GetDrawResponse](
			httpClient,
			baseURL+PrivateDrawServiceGetDrawProcedure,
			opts...,
		),
		deleteDraw: connect_go.NewClient[private.DeleteDrawRequest, private.DeleteDrawResponse](
			httpClient,
			baseURL+PrivateDrawServiceDeleteDrawProcedure,
			opts...,
		),
		listDraws: connect_go.NewClient[private.ListDrawsRequest, private.ListDrawsResponse](
			httpClient,
			baseURL+PrivateDrawServiceListDrawsProcedure,
			opts...,
		),
		createFlexibleRoundRobinDraw: connect_go.NewClient[flexible_round_robin.CreateFlexibleRoundRobinDrawRequest, flexible_round_robin.CreateFlexibleRoundRobinDrawResponse](
			httpClient,
			baseURL+PrivateDrawServiceCreateFlexibleRoundRobinDrawProcedure,
			opts...,
		),
		updateFlexibleRoundRobinDraw: connect_go.NewClient[flexible_round_robin.UpdateFlexibleRoundRobinDrawRequest, flexible_round_robin.UpdateFlexibleRoundRobinDrawResponse](
			httpClient,
			baseURL+PrivateDrawServiceUpdateFlexibleRoundRobinDrawProcedure,
			opts...,
		),
		getFlexibleRoundRobinDrawRound: connect_go.NewClient[flexible_round_robin.GetFlexibleRoundRobinDrawRoundRequest, flexible_round_robin.GetFlexibleRoundRobinDrawRoundResponse](
			httpClient,
			baseURL+PrivateDrawServiceGetFlexibleRoundRobinDrawRoundProcedure,
			opts...,
		),
	}
}

// privateDrawServiceClient implements PrivateDrawServiceClient.
type privateDrawServiceClient struct {
	getDraw                        *connect_go.Client[private.GetDrawRequest, private.GetDrawResponse]
	deleteDraw                     *connect_go.Client[private.DeleteDrawRequest, private.DeleteDrawResponse]
	listDraws                      *connect_go.Client[private.ListDrawsRequest, private.ListDrawsResponse]
	createFlexibleRoundRobinDraw   *connect_go.Client[flexible_round_robin.CreateFlexibleRoundRobinDrawRequest, flexible_round_robin.CreateFlexibleRoundRobinDrawResponse]
	updateFlexibleRoundRobinDraw   *connect_go.Client[flexible_round_robin.UpdateFlexibleRoundRobinDrawRequest, flexible_round_robin.UpdateFlexibleRoundRobinDrawResponse]
	getFlexibleRoundRobinDrawRound *connect_go.Client[flexible_round_robin.GetFlexibleRoundRobinDrawRoundRequest, flexible_round_robin.GetFlexibleRoundRobinDrawRoundResponse]
}

// GetDraw calls draw.v1.private.PrivateDrawService.GetDraw.
func (c *privateDrawServiceClient) GetDraw(ctx context.Context, req *connect_go.Request[private.GetDrawRequest]) (*connect_go.Response[private.GetDrawResponse], error) {
	return c.getDraw.CallUnary(ctx, req)
}

// DeleteDraw calls draw.v1.private.PrivateDrawService.DeleteDraw.
func (c *privateDrawServiceClient) DeleteDraw(ctx context.Context, req *connect_go.Request[private.DeleteDrawRequest]) (*connect_go.Response[private.DeleteDrawResponse], error) {
	return c.deleteDraw.CallUnary(ctx, req)
}

// ListDraws calls draw.v1.private.PrivateDrawService.ListDraws.
func (c *privateDrawServiceClient) ListDraws(ctx context.Context, req *connect_go.Request[private.ListDrawsRequest]) (*connect_go.Response[private.ListDrawsResponse], error) {
	return c.listDraws.CallUnary(ctx, req)
}

// CreateFlexibleRoundRobinDraw calls
// draw.v1.private.PrivateDrawService.CreateFlexibleRoundRobinDraw.
func (c *privateDrawServiceClient) CreateFlexibleRoundRobinDraw(ctx context.Context, req *connect_go.Request[flexible_round_robin.CreateFlexibleRoundRobinDrawRequest]) (*connect_go.Response[flexible_round_robin.CreateFlexibleRoundRobinDrawResponse], error) {
	return c.createFlexibleRoundRobinDraw.CallUnary(ctx, req)
}

// UpdateFlexibleRoundRobinDraw calls
// draw.v1.private.PrivateDrawService.UpdateFlexibleRoundRobinDraw.
func (c *privateDrawServiceClient) UpdateFlexibleRoundRobinDraw(ctx context.Context, req *connect_go.Request[flexible_round_robin.UpdateFlexibleRoundRobinDrawRequest]) (*connect_go.Response[flexible_round_robin.UpdateFlexibleRoundRobinDrawResponse], error) {
	return c.updateFlexibleRoundRobinDraw.CallUnary(ctx, req)
}

// GetFlexibleRoundRobinDrawRound calls
// draw.v1.private.PrivateDrawService.GetFlexibleRoundRobinDrawRound.
func (c *privateDrawServiceClient) GetFlexibleRoundRobinDrawRound(ctx context.Context, req *connect_go.Request[flexible_round_robin.GetFlexibleRoundRobinDrawRoundRequest]) (*connect_go.Response[flexible_round_robin.GetFlexibleRoundRobinDrawRoundResponse], error) {
	return c.getFlexibleRoundRobinDrawRound.CallUnary(ctx, req)
}

// PrivateDrawServiceHandler is an implementation of the draw.v1.private.PrivateDrawService service.
type PrivateDrawServiceHandler interface {
	// Common endpoints
	GetDraw(context.Context, *connect_go.Request[private.GetDrawRequest]) (*connect_go.Response[private.GetDrawResponse], error)
	DeleteDraw(context.Context, *connect_go.Request[private.DeleteDrawRequest]) (*connect_go.Response[private.DeleteDrawResponse], error)
	ListDraws(context.Context, *connect_go.Request[private.ListDrawsRequest]) (*connect_go.Response[private.ListDrawsResponse], error)
	// Flexible round robin
	CreateFlexibleRoundRobinDraw(context.Context, *connect_go.Request[flexible_round_robin.CreateFlexibleRoundRobinDrawRequest]) (*connect_go.Response[flexible_round_robin.CreateFlexibleRoundRobinDrawResponse], error)
	UpdateFlexibleRoundRobinDraw(context.Context, *connect_go.Request[flexible_round_robin.UpdateFlexibleRoundRobinDrawRequest]) (*connect_go.Response[flexible_round_robin.UpdateFlexibleRoundRobinDrawResponse], error)
	GetFlexibleRoundRobinDrawRound(context.Context, *connect_go.Request[flexible_round_robin.GetFlexibleRoundRobinDrawRoundRequest]) (*connect_go.Response[flexible_round_robin.GetFlexibleRoundRobinDrawRoundResponse], error)
}

// NewPrivateDrawServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPrivateDrawServiceHandler(svc PrivateDrawServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	privateDrawServiceGetDrawHandler := connect_go.NewUnaryHandler(
		PrivateDrawServiceGetDrawProcedure,
		svc.GetDraw,
		opts...,
	)
	privateDrawServiceDeleteDrawHandler := connect_go.NewUnaryHandler(
		PrivateDrawServiceDeleteDrawProcedure,
		svc.DeleteDraw,
		opts...,
	)
	privateDrawServiceListDrawsHandler := connect_go.NewUnaryHandler(
		PrivateDrawServiceListDrawsProcedure,
		svc.ListDraws,
		opts...,
	)
	privateDrawServiceCreateFlexibleRoundRobinDrawHandler := connect_go.NewUnaryHandler(
		PrivateDrawServiceCreateFlexibleRoundRobinDrawProcedure,
		svc.CreateFlexibleRoundRobinDraw,
		opts...,
	)
	privateDrawServiceUpdateFlexibleRoundRobinDrawHandler := connect_go.NewUnaryHandler(
		PrivateDrawServiceUpdateFlexibleRoundRobinDrawProcedure,
		svc.UpdateFlexibleRoundRobinDraw,
		opts...,
	)
	privateDrawServiceGetFlexibleRoundRobinDrawRoundHandler := connect_go.NewUnaryHandler(
		PrivateDrawServiceGetFlexibleRoundRobinDrawRoundProcedure,
		svc.GetFlexibleRoundRobinDrawRound,
		opts...,
	)
	return "/draw.v1.private.PrivateDrawService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PrivateDrawServiceGetDrawProcedure:
			privateDrawServiceGetDrawHandler.ServeHTTP(w, r)
		case PrivateDrawServiceDeleteDrawProcedure:
			privateDrawServiceDeleteDrawHandler.ServeHTTP(w, r)
		case PrivateDrawServiceListDrawsProcedure:
			privateDrawServiceListDrawsHandler.ServeHTTP(w, r)
		case PrivateDrawServiceCreateFlexibleRoundRobinDrawProcedure:
			privateDrawServiceCreateFlexibleRoundRobinDrawHandler.ServeHTTP(w, r)
		case PrivateDrawServiceUpdateFlexibleRoundRobinDrawProcedure:
			privateDrawServiceUpdateFlexibleRoundRobinDrawHandler.ServeHTTP(w, r)
		case PrivateDrawServiceGetFlexibleRoundRobinDrawRoundProcedure:
			privateDrawServiceGetFlexibleRoundRobinDrawRoundHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPrivateDrawServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPrivateDrawServiceHandler struct{}

func (UnimplementedPrivateDrawServiceHandler) GetDraw(context.Context, *connect_go.Request[private.GetDrawRequest]) (*connect_go.Response[private.GetDrawResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("draw.v1.private.PrivateDrawService.GetDraw is not implemented"))
}

func (UnimplementedPrivateDrawServiceHandler) DeleteDraw(context.Context, *connect_go.Request[private.DeleteDrawRequest]) (*connect_go.Response[private.DeleteDrawResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("draw.v1.private.PrivateDrawService.DeleteDraw is not implemented"))
}

func (UnimplementedPrivateDrawServiceHandler) ListDraws(context.Context, *connect_go.Request[private.ListDrawsRequest]) (*connect_go.Response[private.ListDrawsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("draw.v1.private.PrivateDrawService.ListDraws is not implemented"))
}

func (UnimplementedPrivateDrawServiceHandler) CreateFlexibleRoundRobinDraw(context.Context, *connect_go.Request[flexible_round_robin.CreateFlexibleRoundRobinDrawRequest]) (*connect_go.Response[flexible_round_robin.CreateFlexibleRoundRobinDrawResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("draw.v1.private.PrivateDrawService.CreateFlexibleRoundRobinDraw is not implemented"))
}

func (UnimplementedPrivateDrawServiceHandler) UpdateFlexibleRoundRobinDraw(context.Context, *connect_go.Request[flexible_round_robin.UpdateFlexibleRoundRobinDrawRequest]) (*connect_go.Response[flexible_round_robin.UpdateFlexibleRoundRobinDrawResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("draw.v1.private.PrivateDrawService.UpdateFlexibleRoundRobinDraw is not implemented"))
}

func (UnimplementedPrivateDrawServiceHandler) GetFlexibleRoundRobinDrawRound(context.Context, *connect_go.Request[flexible_round_robin.GetFlexibleRoundRobinDrawRoundRequest]) (*connect_go.Response[flexible_round_robin.GetFlexibleRoundRobinDrawRoundResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("draw.v1.private.PrivateDrawService.GetFlexibleRoundRobinDrawRound is not implemented"))
}
