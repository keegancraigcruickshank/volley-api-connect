// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: scoring/v1/private/scoring.proto

package scoringv1privateconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	private "github.com/keegancraigcruickshank/volley-api-connect/scoring/v1/private"
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
	// PrivateScoringServiceName is the fully-qualified name of the PrivateScoringService service.
	PrivateScoringServiceName = "scoring.v1.private.PrivateScoringService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PrivateScoringServiceUploadScoreProcedure is the fully-qualified name of the
	// PrivateScoringService's UploadScore RPC.
	PrivateScoringServiceUploadScoreProcedure = "/scoring.v1.private.PrivateScoringService/UploadScore"
)

// PrivateScoringServiceClient is a client for the scoring.v1.private.PrivateScoringService service.
type PrivateScoringServiceClient interface {
	// Common endpoints
	UploadScore(context.Context, *connect_go.Request[private.UploadScoreRequest]) (*connect_go.Response[private.UploadScoreResponse], error)
}

// NewPrivateScoringServiceClient constructs a client for the
// scoring.v1.private.PrivateScoringService service. By default, it uses the Connect protocol with
// the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed requests. To use
// the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPrivateScoringServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) PrivateScoringServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &privateScoringServiceClient{
		uploadScore: connect_go.NewClient[private.UploadScoreRequest, private.UploadScoreResponse](
			httpClient,
			baseURL+PrivateScoringServiceUploadScoreProcedure,
			opts...,
		),
	}
}

// privateScoringServiceClient implements PrivateScoringServiceClient.
type privateScoringServiceClient struct {
	uploadScore *connect_go.Client[private.UploadScoreRequest, private.UploadScoreResponse]
}

// UploadScore calls scoring.v1.private.PrivateScoringService.UploadScore.
func (c *privateScoringServiceClient) UploadScore(ctx context.Context, req *connect_go.Request[private.UploadScoreRequest]) (*connect_go.Response[private.UploadScoreResponse], error) {
	return c.uploadScore.CallUnary(ctx, req)
}

// PrivateScoringServiceHandler is an implementation of the scoring.v1.private.PrivateScoringService
// service.
type PrivateScoringServiceHandler interface {
	// Common endpoints
	UploadScore(context.Context, *connect_go.Request[private.UploadScoreRequest]) (*connect_go.Response[private.UploadScoreResponse], error)
}

// NewPrivateScoringServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPrivateScoringServiceHandler(svc PrivateScoringServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	privateScoringServiceUploadScoreHandler := connect_go.NewUnaryHandler(
		PrivateScoringServiceUploadScoreProcedure,
		svc.UploadScore,
		opts...,
	)
	return "/scoring.v1.private.PrivateScoringService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PrivateScoringServiceUploadScoreProcedure:
			privateScoringServiceUploadScoreHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPrivateScoringServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPrivateScoringServiceHandler struct{}

func (UnimplementedPrivateScoringServiceHandler) UploadScore(context.Context, *connect_go.Request[private.UploadScoreRequest]) (*connect_go.Response[private.UploadScoreResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("scoring.v1.private.PrivateScoringService.UploadScore is not implemented"))
}