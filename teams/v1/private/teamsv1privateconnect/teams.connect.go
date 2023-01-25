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

// PrivateTeamsServiceClient is a client for the teams.v1.private.PrivateTeamsService service.
type PrivateTeamsServiceClient interface {
	CreateTeam(context.Context, *connect_go.Request[private.CreateTeamRequest]) (*connect_go.Response[private.CreateTeamResponse], error)
	ListTeams(context.Context, *connect_go.Request[private.ListTeamsRequest]) (*connect_go.Response[private.ListTeamsResponse], error)
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
		createTeam: connect_go.NewClient[private.CreateTeamRequest, private.CreateTeamResponse](
			httpClient,
			baseURL+"/teams.v1.private.PrivateTeamsService/CreateTeam",
			opts...,
		),
		listTeams: connect_go.NewClient[private.ListTeamsRequest, private.ListTeamsResponse](
			httpClient,
			baseURL+"/teams.v1.private.PrivateTeamsService/ListTeams",
			opts...,
		),
	}
}

// privateTeamsServiceClient implements PrivateTeamsServiceClient.
type privateTeamsServiceClient struct {
	createTeam *connect_go.Client[private.CreateTeamRequest, private.CreateTeamResponse]
	listTeams  *connect_go.Client[private.ListTeamsRequest, private.ListTeamsResponse]
}

// CreateTeam calls teams.v1.private.PrivateTeamsService.CreateTeam.
func (c *privateTeamsServiceClient) CreateTeam(ctx context.Context, req *connect_go.Request[private.CreateTeamRequest]) (*connect_go.Response[private.CreateTeamResponse], error) {
	return c.createTeam.CallUnary(ctx, req)
}

// ListTeams calls teams.v1.private.PrivateTeamsService.ListTeams.
func (c *privateTeamsServiceClient) ListTeams(ctx context.Context, req *connect_go.Request[private.ListTeamsRequest]) (*connect_go.Response[private.ListTeamsResponse], error) {
	return c.listTeams.CallUnary(ctx, req)
}

// PrivateTeamsServiceHandler is an implementation of the teams.v1.private.PrivateTeamsService
// service.
type PrivateTeamsServiceHandler interface {
	CreateTeam(context.Context, *connect_go.Request[private.CreateTeamRequest]) (*connect_go.Response[private.CreateTeamResponse], error)
	ListTeams(context.Context, *connect_go.Request[private.ListTeamsRequest]) (*connect_go.Response[private.ListTeamsResponse], error)
}

// NewPrivateTeamsServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPrivateTeamsServiceHandler(svc PrivateTeamsServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/teams.v1.private.PrivateTeamsService/CreateTeam", connect_go.NewUnaryHandler(
		"/teams.v1.private.PrivateTeamsService/CreateTeam",
		svc.CreateTeam,
		opts...,
	))
	mux.Handle("/teams.v1.private.PrivateTeamsService/ListTeams", connect_go.NewUnaryHandler(
		"/teams.v1.private.PrivateTeamsService/ListTeams",
		svc.ListTeams,
		opts...,
	))
	return "/teams.v1.private.PrivateTeamsService/", mux
}

// UnimplementedPrivateTeamsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPrivateTeamsServiceHandler struct{}

func (UnimplementedPrivateTeamsServiceHandler) CreateTeam(context.Context, *connect_go.Request[private.CreateTeamRequest]) (*connect_go.Response[private.CreateTeamResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("teams.v1.private.PrivateTeamsService.CreateTeam is not implemented"))
}

func (UnimplementedPrivateTeamsServiceHandler) ListTeams(context.Context, *connect_go.Request[private.ListTeamsRequest]) (*connect_go.Response[private.ListTeamsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("teams.v1.private.PrivateTeamsService.ListTeams is not implemented"))
}
