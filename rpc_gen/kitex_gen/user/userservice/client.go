// Code generated by Kitex v0.9.1. DO NOT EDIT.

package userservice

import (
	"context"
	user "github.com/PiaoAdmin/gomall/rpc_gen/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateUser(ctx context.Context, Req *user.CreateUserRequest, callOptions ...callopt.Option) (r *user.CreateUserResponse, err error)
	DeleteUser(ctx context.Context, Req *user.DeleteUserRequest, callOptions ...callopt.Option) (r *user.DeleteUserResponse, err error)
	UpdateBaseUser(ctx context.Context, Req *user.UpdateBaseUserRequest, callOptions ...callopt.Option) (r *user.UpdateBaseUserResponse, err error)
	UpdateUserPassword(ctx context.Context, Req *user.UpdateUserPasswordRequest, callOptions ...callopt.Option) (r *user.UpdateUserPasswordResponse, err error)
	UpdateUserBalance(ctx context.Context, Req *user.UpdateUserBalanceRequest, callOptions ...callopt.Option) (r *user.UpdateUserBalanceResponse, err error)
	UpdateUserStatus(ctx context.Context, Req *user.UpdateUserStatusRequest, callOptions ...callopt.Option) (r *user.UpdateUserStatusResponse, err error)
	GetUserInfo(ctx context.Context, Req *user.GetUserInfoRequest, callOptions ...callopt.Option) (r *user.GetUserInfoResponse, err error)
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
	return &kUserServiceClient{
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

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) CreateUser(ctx context.Context, Req *user.CreateUserRequest, callOptions ...callopt.Option) (r *user.CreateUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateUser(ctx, Req)
}

func (p *kUserServiceClient) DeleteUser(ctx context.Context, Req *user.DeleteUserRequest, callOptions ...callopt.Option) (r *user.DeleteUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteUser(ctx, Req)
}

func (p *kUserServiceClient) UpdateBaseUser(ctx context.Context, Req *user.UpdateBaseUserRequest, callOptions ...callopt.Option) (r *user.UpdateBaseUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateBaseUser(ctx, Req)
}

func (p *kUserServiceClient) UpdateUserPassword(ctx context.Context, Req *user.UpdateUserPasswordRequest, callOptions ...callopt.Option) (r *user.UpdateUserPasswordResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateUserPassword(ctx, Req)
}

func (p *kUserServiceClient) UpdateUserBalance(ctx context.Context, Req *user.UpdateUserBalanceRequest, callOptions ...callopt.Option) (r *user.UpdateUserBalanceResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateUserBalance(ctx, Req)
}

func (p *kUserServiceClient) UpdateUserStatus(ctx context.Context, Req *user.UpdateUserStatusRequest, callOptions ...callopt.Option) (r *user.UpdateUserStatusResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateUserStatus(ctx, Req)
}

func (p *kUserServiceClient) GetUserInfo(ctx context.Context, Req *user.GetUserInfoRequest, callOptions ...callopt.Option) (r *user.GetUserInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserInfo(ctx, Req)
}
