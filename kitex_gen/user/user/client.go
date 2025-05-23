// Code generated by Kitex v0.13.1. DO NOT EDIT.

package user

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	basic "github.com/xh-polaris/psych-idl-gen/kitex_gen/basic"
	user "github.com/xh-polaris/psych-idl-gen/kitex_gen/user"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	UserSignUp(ctx context.Context, Req *user.UserSignUpReq, callOptions ...callopt.Option) (r *basic.Response, err error)
	UserGetInfo(ctx context.Context, Req *user.UserGetInfoReq, callOptions ...callopt.Option) (r *user.UserGetInfoResp, err error)
	UserUpdateInfo(ctx context.Context, Req *user.UserUpdateInfoReq, callOptions ...callopt.Option) (r *basic.Response, err error)
	UserUpdatePassword(ctx context.Context, Req *user.UserUpdatePasswordReq, callOptions ...callopt.Option) (r *basic.Response, err error)
	UserBelongUnit(ctx context.Context, Req *user.UserBelongUnitReq, callOptions ...callopt.Option) (r *user.UserBelongUnitResp, err error)
	UserSignIn(ctx context.Context, Req *user.UserSignInReq, callOptions ...callopt.Option) (r *basic.Response, err error)
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
	return &kUserClient{
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

type kUserClient struct {
	*kClient
}

func (p *kUserClient) UserSignUp(ctx context.Context, Req *user.UserSignUpReq, callOptions ...callopt.Option) (r *basic.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserSignUp(ctx, Req)
}

func (p *kUserClient) UserGetInfo(ctx context.Context, Req *user.UserGetInfoReq, callOptions ...callopt.Option) (r *user.UserGetInfoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetInfo(ctx, Req)
}

func (p *kUserClient) UserUpdateInfo(ctx context.Context, Req *user.UserUpdateInfoReq, callOptions ...callopt.Option) (r *basic.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserUpdateInfo(ctx, Req)
}

func (p *kUserClient) UserUpdatePassword(ctx context.Context, Req *user.UserUpdatePasswordReq, callOptions ...callopt.Option) (r *basic.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserUpdatePassword(ctx, Req)
}

func (p *kUserClient) UserBelongUnit(ctx context.Context, Req *user.UserBelongUnitReq, callOptions ...callopt.Option) (r *user.UserBelongUnitResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserBelongUnit(ctx, Req)
}

func (p *kUserClient) UserSignIn(ctx context.Context, Req *user.UserSignInReq, callOptions ...callopt.Option) (r *basic.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserSignIn(ctx, Req)
}
