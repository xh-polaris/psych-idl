// Code generated by Kitex v0.14.1. DO NOT EDIT.

package coreapi

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "github.com/cloudwego/prutal"
	core_api "github.com/xh-polaris/psych-idl/kitex_gen/core_api"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"SignIn": kitex.NewMethodInfo(
		signInHandler,
		newSignInArgs,
		newSignInResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	coreApiServiceInfo                = NewServiceInfo()
	coreApiServiceInfoForClient       = NewServiceInfoForClient()
	coreApiServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return coreApiServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return coreApiServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return coreApiServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "CoreApi"
	handlerType := (*core_api.CoreApi)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "core_api",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.14.1",
		Extra:           extra,
	}
	return svcInfo
}

func signInHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.UserSignInReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.CoreApi).SignIn(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *SignInArgs:
		success, err := handler.(core_api.CoreApi).SignIn(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SignInResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newSignInArgs() interface{} {
	return &SignInArgs{}
}

func newSignInResult() interface{} {
	return &SignInResult{}
}

type SignInArgs struct {
	Req *core_api.UserSignInReq
}

func (p *SignInArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *SignInArgs) Unmarshal(in []byte) error {
	msg := new(core_api.UserSignInReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SignInArgs_Req_DEFAULT *core_api.UserSignInReq

func (p *SignInArgs) GetReq() *core_api.UserSignInReq {
	if !p.IsSetReq() {
		return SignInArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SignInArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SignInArgs) GetFirstArgument() interface{} {
	return p.Req
}

type SignInResult struct {
	Success *core_api.UserSignInResp
}

var SignInResult_Success_DEFAULT *core_api.UserSignInResp

func (p *SignInResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *SignInResult) Unmarshal(in []byte) error {
	msg := new(core_api.UserSignInResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SignInResult) GetSuccess() *core_api.UserSignInResp {
	if !p.IsSetSuccess() {
		return SignInResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SignInResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.UserSignInResp)
}

func (p *SignInResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SignInResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) SignIn(ctx context.Context, Req *core_api.UserSignInReq) (r *core_api.UserSignInResp, err error) {
	var _args SignInArgs
	_args.Req = Req
	var _result SignInResult
	if err = p.c.Call(ctx, "SignIn", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
