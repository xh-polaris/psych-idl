// Code generated by Kitex v0.13.1. DO NOT EDIT.

package view

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	basic "github.com/xh-polaris/psych-idl-gen/kitex_gen/basic"
	user "github.com/xh-polaris/psych-idl-gen/kitex_gen/user"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	ViewSignUp(ctx context.Context, Req *user.ViewSignUpReq, callOptions ...callopt.Option) (r *basic.Response, err error)
	ViewGetInfo(ctx context.Context, Req *user.ViewGetInfoReq, callOptions ...callopt.Option) (r *user.ViewGetInfoResp, err error)
	ViewUpdateInfo(ctx context.Context, Req *user.ViewUpdateInfoReq, callOptions ...callopt.Option) (r *basic.Response, err error)
	ViewUpdatePassword(ctx context.Context, Req *user.ViewUpdatePasswordReq, callOptions ...callopt.Option) (r *basic.Response, err error)
	ViewBelongUnit(ctx context.Context, Req *user.ViewBelongUnitReq, callOptions ...callopt.Option) (r *user.ViewBelongUnitResp, err error)
	ViewSignIn(ctx context.Context, Req *user.ViewSignInReq, callOptions ...callopt.Option) (r *basic.Response, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kViewClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kViewClient struct {
	*kClient
}

func (p *kViewClient) ViewSignUp(ctx context.Context, Req *user.ViewSignUpReq, callOptions ...callopt.Option) (r *basic.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ViewSignUp(ctx, Req)
}

func (p *kViewClient) ViewGetInfo(ctx context.Context, Req *user.ViewGetInfoReq, callOptions ...callopt.Option) (r *user.ViewGetInfoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ViewGetInfo(ctx, Req)
}

func (p *kViewClient) ViewUpdateInfo(ctx context.Context, Req *user.ViewUpdateInfoReq, callOptions ...callopt.Option) (r *basic.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ViewUpdateInfo(ctx, Req)
}

func (p *kViewClient) ViewUpdatePassword(ctx context.Context, Req *user.ViewUpdatePasswordReq, callOptions ...callopt.Option) (r *basic.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ViewUpdatePassword(ctx, Req)
}

func (p *kViewClient) ViewBelongUnit(ctx context.Context, Req *user.ViewBelongUnitReq, callOptions ...callopt.Option) (r *user.ViewBelongUnitResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ViewBelongUnit(ctx, Req)
}

func (p *kViewClient) ViewSignIn(ctx context.Context, Req *user.ViewSignInReq, callOptions ...callopt.Option) (r *basic.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ViewSignIn(ctx, Req)
}
