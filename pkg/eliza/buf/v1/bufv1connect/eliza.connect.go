// Copyright 2022-2023 The Connect Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: buf/v1/eliza.proto

package bufv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/junsazanami430u/examples-go/pkg/eliza/buf/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// ElizaServiceName is the fully-qualified name of the ElizaService service.
	ElizaServiceName = "buf.v1.ElizaService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ElizaServiceSayProcedure is the fully-qualified name of the ElizaService's Say RPC.
	ElizaServiceSayProcedure = "/buf.v1.ElizaService/Say"
	// ElizaServiceConverseProcedure is the fully-qualified name of the ElizaService's Converse RPC.
	ElizaServiceConverseProcedure = "/buf.v1.ElizaService/Converse"
	// ElizaServiceIntroduceProcedure is the fully-qualified name of the ElizaService's Introduce RPC.
	ElizaServiceIntroduceProcedure = "/buf.v1.ElizaService/Introduce"
	// ElizaServiceGoodByeProcedure is the fully-qualified name of the ElizaService's GoodBye RPC.
	ElizaServiceGoodByeProcedure = "/buf.v1.ElizaService/GoodBye"
)

// ElizaServiceClient is a client for the buf.v1.ElizaService service.
type ElizaServiceClient interface {
	// Say is a unary RPC. Eliza responds to the prompt with a single sentence.
	Say(context.Context, *connect.Request[v1.SayRequest]) (*connect.Response[v1.SayResponse], error)
	// Converse is a bidirectional RPC. The caller may exchange multiple
	// back-and-forth messages with Eliza over a long-lived connection. Eliza
	// responds to each ConverseRequest with a ConverseResponse.
	Converse(context.Context) *connect.BidiStreamForClient[v1.ConverseRequest, v1.ConverseResponse]
	// Introduce is a server streaming RPC. Given the caller's name, Eliza
	// returns a stream of sentences to introduce itself.
	Introduce(context.Context, *connect.Request[v1.IntroduceRequest]) (*connect.ServerStreamForClient[v1.IntroduceResponse], error)
	GoodBye(context.Context, *connect.Request[v1.GoodByeRequest]) (*connect.Response[v1.GoodByeResponse], error)
}

// NewElizaServiceClient constructs a client for the buf.v1.ElizaService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewElizaServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ElizaServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	elizaServiceMethods := v1.File_buf_v1_eliza_proto.Services().ByName("ElizaService").Methods()
	return &elizaServiceClient{
		say: connect.NewClient[v1.SayRequest, v1.SayResponse](
			httpClient,
			baseURL+ElizaServiceSayProcedure,
			connect.WithSchema(elizaServiceMethods.ByName("Say")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		converse: connect.NewClient[v1.ConverseRequest, v1.ConverseResponse](
			httpClient,
			baseURL+ElizaServiceConverseProcedure,
			connect.WithSchema(elizaServiceMethods.ByName("Converse")),
			connect.WithClientOptions(opts...),
		),
		introduce: connect.NewClient[v1.IntroduceRequest, v1.IntroduceResponse](
			httpClient,
			baseURL+ElizaServiceIntroduceProcedure,
			connect.WithSchema(elizaServiceMethods.ByName("Introduce")),
			connect.WithClientOptions(opts...),
		),
		goodBye: connect.NewClient[v1.GoodByeRequest, v1.GoodByeResponse](
			httpClient,
			baseURL+ElizaServiceGoodByeProcedure,
			connect.WithSchema(elizaServiceMethods.ByName("GoodBye")),
			connect.WithClientOptions(opts...),
		),
	}
}

// elizaServiceClient implements ElizaServiceClient.
type elizaServiceClient struct {
	say       *connect.Client[v1.SayRequest, v1.SayResponse]
	converse  *connect.Client[v1.ConverseRequest, v1.ConverseResponse]
	introduce *connect.Client[v1.IntroduceRequest, v1.IntroduceResponse]
	goodBye   *connect.Client[v1.GoodByeRequest, v1.GoodByeResponse]
}

// Say calls buf.v1.ElizaService.Say.
func (c *elizaServiceClient) Say(ctx context.Context, req *connect.Request[v1.SayRequest]) (*connect.Response[v1.SayResponse], error) {
	return c.say.CallUnary(ctx, req)
}

// Converse calls buf.v1.ElizaService.Converse.
func (c *elizaServiceClient) Converse(ctx context.Context) *connect.BidiStreamForClient[v1.ConverseRequest, v1.ConverseResponse] {
	return c.converse.CallBidiStream(ctx)
}

// Introduce calls buf.v1.ElizaService.Introduce.
func (c *elizaServiceClient) Introduce(ctx context.Context, req *connect.Request[v1.IntroduceRequest]) (*connect.ServerStreamForClient[v1.IntroduceResponse], error) {
	return c.introduce.CallServerStream(ctx, req)
}

// GoodBye calls buf.v1.ElizaService.GoodBye.
func (c *elizaServiceClient) GoodBye(ctx context.Context, req *connect.Request[v1.GoodByeRequest]) (*connect.Response[v1.GoodByeResponse], error) {
	return c.goodBye.CallUnary(ctx, req)
}

// ElizaServiceHandler is an implementation of the buf.v1.ElizaService service.
type ElizaServiceHandler interface {
	// Say is a unary RPC. Eliza responds to the prompt with a single sentence.
	Say(context.Context, *connect.Request[v1.SayRequest]) (*connect.Response[v1.SayResponse], error)
	// Converse is a bidirectional RPC. The caller may exchange multiple
	// back-and-forth messages with Eliza over a long-lived connection. Eliza
	// responds to each ConverseRequest with a ConverseResponse.
	Converse(context.Context, *connect.BidiStream[v1.ConverseRequest, v1.ConverseResponse]) error
	// Introduce is a server streaming RPC. Given the caller's name, Eliza
	// returns a stream of sentences to introduce itself.
	Introduce(context.Context, *connect.Request[v1.IntroduceRequest], *connect.ServerStream[v1.IntroduceResponse]) error
	GoodBye(context.Context, *connect.Request[v1.GoodByeRequest]) (*connect.Response[v1.GoodByeResponse], error)
}

// NewElizaServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewElizaServiceHandler(svc ElizaServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	elizaServiceMethods := v1.File_buf_v1_eliza_proto.Services().ByName("ElizaService").Methods()
	elizaServiceSayHandler := connect.NewUnaryHandler(
		ElizaServiceSayProcedure,
		svc.Say,
		connect.WithSchema(elizaServiceMethods.ByName("Say")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	elizaServiceConverseHandler := connect.NewBidiStreamHandler(
		ElizaServiceConverseProcedure,
		svc.Converse,
		connect.WithSchema(elizaServiceMethods.ByName("Converse")),
		connect.WithHandlerOptions(opts...),
	)
	elizaServiceIntroduceHandler := connect.NewServerStreamHandler(
		ElizaServiceIntroduceProcedure,
		svc.Introduce,
		connect.WithSchema(elizaServiceMethods.ByName("Introduce")),
		connect.WithHandlerOptions(opts...),
	)
	elizaServiceGoodByeHandler := connect.NewUnaryHandler(
		ElizaServiceGoodByeProcedure,
		svc.GoodBye,
		connect.WithSchema(elizaServiceMethods.ByName("GoodBye")),
		connect.WithHandlerOptions(opts...),
	)
	return "/buf.v1.ElizaService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ElizaServiceSayProcedure:
			elizaServiceSayHandler.ServeHTTP(w, r)
		case ElizaServiceConverseProcedure:
			elizaServiceConverseHandler.ServeHTTP(w, r)
		case ElizaServiceIntroduceProcedure:
			elizaServiceIntroduceHandler.ServeHTTP(w, r)
		case ElizaServiceGoodByeProcedure:
			elizaServiceGoodByeHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedElizaServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedElizaServiceHandler struct{}

func (UnimplementedElizaServiceHandler) Say(context.Context, *connect.Request[v1.SayRequest]) (*connect.Response[v1.SayResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("buf.v1.ElizaService.Say is not implemented"))
}

func (UnimplementedElizaServiceHandler) Converse(context.Context, *connect.BidiStream[v1.ConverseRequest, v1.ConverseResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("buf.v1.ElizaService.Converse is not implemented"))
}

func (UnimplementedElizaServiceHandler) Introduce(context.Context, *connect.Request[v1.IntroduceRequest], *connect.ServerStream[v1.IntroduceResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("buf.v1.ElizaService.Introduce is not implemented"))
}

func (UnimplementedElizaServiceHandler) GoodBye(context.Context, *connect.Request[v1.GoodByeRequest]) (*connect.Response[v1.GoodByeResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("buf.v1.ElizaService.GoodBye is not implemented"))
}
