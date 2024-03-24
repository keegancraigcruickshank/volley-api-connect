// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: teams/v1/private/teams.proto

package teamsv1privateconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	private "github.com/keegancraigcruickshank/volley-api-connect/teams/v1/private"
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
	// PrivateTeamsServiceName is the fully-qualified name of the PrivateTeamsService service.
	PrivateTeamsServiceName = "teams.v1.private.PrivateTeamsService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PrivateTeamsServiceAddTeamProcedure is the fully-qualified name of the PrivateTeamsService's
	// AddTeam RPC.
	PrivateTeamsServiceAddTeamProcedure = "/teams.v1.private.PrivateTeamsService/AddTeam"
	// PrivateTeamsServiceListTeamsProcedure is the fully-qualified name of the PrivateTeamsService's
	// ListTeams RPC.
	PrivateTeamsServiceListTeamsProcedure = "/teams.v1.private.PrivateTeamsService/ListTeams"
	// PrivateTeamsServiceRemoveTeamsProcedure is the fully-qualified name of the PrivateTeamsService's
	// RemoveTeams RPC.
	PrivateTeamsServiceRemoveTeamsProcedure = "/teams.v1.private.PrivateTeamsService/RemoveTeams"
	// PrivateTeamsServiceUpdateTeamProcedure is the fully-qualified name of the PrivateTeamsService's
	// UpdateTeam RPC.
	PrivateTeamsServiceUpdateTeamProcedure = "/teams.v1.private.PrivateTeamsService/UpdateTeam"
)

// PrivateTeamsServiceClient is a client for the teams.v1.private.PrivateTeamsService service.
type PrivateTeamsServiceClient interface {
	AddTeam(context.Context, *connect_go.Request[private.AddTeamRequest]) (*connect_go.Response[private.AddTeamResponse], error)
	ListTeams(context.Context, *connect_go.Request[private.AddTeamRequest]) (*connect_go.Response[private.ListTeamsResponse], error)
	RemoveTeams(context.Context, *connect_go.Request[private.RemoveTeamsRequest]) (*connect_go.Response[private.RemoveTeamsResponse], error)
	UpdateTeam(context.Context, *connect_go.Request[private.UpdateTeamRequest]) (*connect_go.Response[private.UpdateTeamResponse], error)
}

// NewPrivateTeamsServiceClient constructs a client for the teams.v1.private.PrivateTeamsService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPrivateTeamsServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) PrivateTeamsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &privateTeamsServiceClient{
		addTeam: connect_go.NewClient[private.AddTeamRequest, private.AddTeamResponse](
			httpClient,
			baseURL+PrivateTeamsServiceAddTeamProcedure,
			opts...,
		),
		listTeams: connect_go.NewClient[private.AddTeamRequest, private.ListTeamsResponse](
			httpClient,
			baseURL+PrivateTeamsServiceListTeamsProcedure,
			opts...,
		),
		removeTeams: connect_go.NewClient[private.RemoveTeamsRequest, private.RemoveTeamsResponse](
			httpClient,
			baseURL+PrivateTeamsServiceRemoveTeamsProcedure,
			opts...,
		),
		updateTeam: connect_go.NewClient[private.UpdateTeamRequest, private.UpdateTeamResponse](
			httpClient,
			baseURL+PrivateTeamsServiceUpdateTeamProcedure,
			opts...,
		),
	}
}

// privateTeamsServiceClient implements PrivateTeamsServiceClient.
type privateTeamsServiceClient struct {
	addTeam     *connect_go.Client[private.AddTeamRequest, private.AddTeamResponse]
	listTeams   *connect_go.Client[private.AddTeamRequest, private.ListTeamsResponse]
	removeTeams *connect_go.Client[private.RemoveTeamsRequest, private.RemoveTeamsResponse]
	updateTeam  *connect_go.Client[private.UpdateTeamRequest, private.UpdateTeamResponse]
}

// AddTeam calls teams.v1.private.PrivateTeamsService.AddTeam.
func (c *privateTeamsServiceClient) AddTeam(ctx context.Context, req *connect_go.Request[private.AddTeamRequest]) (*connect_go.Response[private.AddTeamResponse], error) {
	return c.addTeam.CallUnary(ctx, req)
}

// ListTeams calls teams.v1.private.PrivateTeamsService.ListTeams.
func (c *privateTeamsServiceClient) ListTeams(ctx context.Context, req *connect_go.Request[private.AddTeamRequest]) (*connect_go.Response[private.ListTeamsResponse], error) {
	return c.listTeams.CallUnary(ctx, req)
}

// RemoveTeams calls teams.v1.private.PrivateTeamsService.RemoveTeams.
func (c *privateTeamsServiceClient) RemoveTeams(ctx context.Context, req *connect_go.Request[private.RemoveTeamsRequest]) (*connect_go.Response[private.RemoveTeamsResponse], error) {
	return c.removeTeams.CallUnary(ctx, req)
}

// UpdateTeam calls teams.v1.private.PrivateTeamsService.UpdateTeam.
func (c *privateTeamsServiceClient) UpdateTeam(ctx context.Context, req *connect_go.Request[private.UpdateTeamRequest]) (*connect_go.Response[private.UpdateTeamResponse], error) {
	return c.updateTeam.CallUnary(ctx, req)
}

// PrivateTeamsServiceHandler is an implementation of the teams.v1.private.PrivateTeamsService
// service.
type PrivateTeamsServiceHandler interface {
	AddTeam(context.Context, *connect_go.Request[private.AddTeamRequest]) (*connect_go.Response[private.AddTeamResponse], error)
	ListTeams(context.Context, *connect_go.Request[private.AddTeamRequest]) (*connect_go.Response[private.ListTeamsResponse], error)
	RemoveTeams(context.Context, *connect_go.Request[private.RemoveTeamsRequest]) (*connect_go.Response[private.RemoveTeamsResponse], error)
	UpdateTeam(context.Context, *connect_go.Request[private.UpdateTeamRequest]) (*connect_go.Response[private.UpdateTeamResponse], error)
}

// NewPrivateTeamsServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPrivateTeamsServiceHandler(svc PrivateTeamsServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	privateTeamsServiceAddTeamHandler := connect_go.NewUnaryHandler(
		PrivateTeamsServiceAddTeamProcedure,
		svc.AddTeam,
		opts...,
	)
	privateTeamsServiceListTeamsHandler := connect_go.NewUnaryHandler(
		PrivateTeamsServiceListTeamsProcedure,
		svc.ListTeams,
		opts...,
	)
	privateTeamsServiceRemoveTeamsHandler := connect_go.NewUnaryHandler(
		PrivateTeamsServiceRemoveTeamsProcedure,
		svc.RemoveTeams,
		opts...,
	)
	privateTeamsServiceUpdateTeamHandler := connect_go.NewUnaryHandler(
		PrivateTeamsServiceUpdateTeamProcedure,
		svc.UpdateTeam,
		opts...,
	)
	return "/teams.v1.private.PrivateTeamsService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PrivateTeamsServiceAddTeamProcedure:
			privateTeamsServiceAddTeamHandler.ServeHTTP(w, r)
		case PrivateTeamsServiceListTeamsProcedure:
			privateTeamsServiceListTeamsHandler.ServeHTTP(w, r)
		case PrivateTeamsServiceRemoveTeamsProcedure:
			privateTeamsServiceRemoveTeamsHandler.ServeHTTP(w, r)
		case PrivateTeamsServiceUpdateTeamProcedure:
			privateTeamsServiceUpdateTeamHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPrivateTeamsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPrivateTeamsServiceHandler struct{}

func (UnimplementedPrivateTeamsServiceHandler) AddTeam(context.Context, *connect_go.Request[private.AddTeamRequest]) (*connect_go.Response[private.AddTeamResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("teams.v1.private.PrivateTeamsService.AddTeam is not implemented"))
}

func (UnimplementedPrivateTeamsServiceHandler) ListTeams(context.Context, *connect_go.Request[private.AddTeamRequest]) (*connect_go.Response[private.ListTeamsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("teams.v1.private.PrivateTeamsService.ListTeams is not implemented"))
}

func (UnimplementedPrivateTeamsServiceHandler) RemoveTeams(context.Context, *connect_go.Request[private.RemoveTeamsRequest]) (*connect_go.Response[private.RemoveTeamsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("teams.v1.private.PrivateTeamsService.RemoveTeams is not implemented"))
}

func (UnimplementedPrivateTeamsServiceHandler) UpdateTeam(context.Context, *connect_go.Request[private.UpdateTeamRequest]) (*connect_go.Response[private.UpdateTeamResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("teams.v1.private.PrivateTeamsService.UpdateTeam is not implemented"))
}
