// Code generated by Kitex v0.9.1. DO NOT EDIT.

package userservice

import (
	"context"
	"errors"
	user "github.com/PiaoAdmin/gomall/rpc_gen/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"CreateUser": kitex.NewMethodInfo(
		createUserHandler,
		newCreateUserArgs,
		newCreateUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"DeleteUser": kitex.NewMethodInfo(
		deleteUserHandler,
		newDeleteUserArgs,
		newDeleteUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"UpdateBaseUser": kitex.NewMethodInfo(
		updateBaseUserHandler,
		newUpdateBaseUserArgs,
		newUpdateBaseUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"UpdateUserPassword": kitex.NewMethodInfo(
		updateUserPasswordHandler,
		newUpdateUserPasswordArgs,
		newUpdateUserPasswordResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"UpdateUserBalance": kitex.NewMethodInfo(
		updateUserBalanceHandler,
		newUpdateUserBalanceArgs,
		newUpdateUserBalanceResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"UpdateUserStatus": kitex.NewMethodInfo(
		updateUserStatusHandler,
		newUpdateUserStatusArgs,
		newUpdateUserStatusResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetUserInfo": kitex.NewMethodInfo(
		getUserInfoHandler,
		newGetUserInfoArgs,
		newGetUserInfoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	userServiceServiceInfo                = NewServiceInfo()
	userServiceServiceInfoForClient       = NewServiceInfoForClient()
	userServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForClient
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
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
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
		"PackageName": "user",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func createUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.CreateUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).CreateUser(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *CreateUserArgs:
		success, err := handler.(user.UserService).CreateUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateUserResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newCreateUserArgs() interface{} {
	return &CreateUserArgs{}
}

func newCreateUserResult() interface{} {
	return &CreateUserResult{}
}

type CreateUserArgs struct {
	Req *user.CreateUserRequest
}

func (p *CreateUserArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.CreateUserRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CreateUserArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CreateUserArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CreateUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CreateUserArgs) Unmarshal(in []byte) error {
	msg := new(user.CreateUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateUserArgs_Req_DEFAULT *user.CreateUserRequest

func (p *CreateUserArgs) GetReq() *user.CreateUserRequest {
	if !p.IsSetReq() {
		return CreateUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateUserArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CreateUserArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CreateUserResult struct {
	Success *user.CreateUserResponse
}

var CreateUserResult_Success_DEFAULT *user.CreateUserResponse

func (p *CreateUserResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.CreateUserResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CreateUserResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CreateUserResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CreateUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CreateUserResult) Unmarshal(in []byte) error {
	msg := new(user.CreateUserResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateUserResult) GetSuccess() *user.CreateUserResponse {
	if !p.IsSetSuccess() {
		return CreateUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.CreateUserResponse)
}

func (p *CreateUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CreateUserResult) GetResult() interface{} {
	return p.Success
}

func deleteUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.DeleteUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).DeleteUser(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *DeleteUserArgs:
		success, err := handler.(user.UserService).DeleteUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteUserResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newDeleteUserArgs() interface{} {
	return &DeleteUserArgs{}
}

func newDeleteUserResult() interface{} {
	return &DeleteUserResult{}
}

type DeleteUserArgs struct {
	Req *user.DeleteUserRequest
}

func (p *DeleteUserArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.DeleteUserRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeleteUserArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeleteUserArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeleteUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteUserArgs) Unmarshal(in []byte) error {
	msg := new(user.DeleteUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteUserArgs_Req_DEFAULT *user.DeleteUserRequest

func (p *DeleteUserArgs) GetReq() *user.DeleteUserRequest {
	if !p.IsSetReq() {
		return DeleteUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteUserArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DeleteUserArgs) GetFirstArgument() interface{} {
	return p.Req
}

type DeleteUserResult struct {
	Success *user.DeleteUserResponse
}

var DeleteUserResult_Success_DEFAULT *user.DeleteUserResponse

func (p *DeleteUserResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.DeleteUserResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeleteUserResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeleteUserResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeleteUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteUserResult) Unmarshal(in []byte) error {
	msg := new(user.DeleteUserResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteUserResult) GetSuccess() *user.DeleteUserResponse {
	if !p.IsSetSuccess() {
		return DeleteUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.DeleteUserResponse)
}

func (p *DeleteUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteUserResult) GetResult() interface{} {
	return p.Success
}

func updateBaseUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.UpdateBaseUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).UpdateBaseUser(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *UpdateBaseUserArgs:
		success, err := handler.(user.UserService).UpdateBaseUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateBaseUserResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newUpdateBaseUserArgs() interface{} {
	return &UpdateBaseUserArgs{}
}

func newUpdateBaseUserResult() interface{} {
	return &UpdateBaseUserResult{}
}

type UpdateBaseUserArgs struct {
	Req *user.UpdateBaseUserRequest
}

func (p *UpdateBaseUserArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.UpdateBaseUserRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateBaseUserArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateBaseUserArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateBaseUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateBaseUserArgs) Unmarshal(in []byte) error {
	msg := new(user.UpdateBaseUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateBaseUserArgs_Req_DEFAULT *user.UpdateBaseUserRequest

func (p *UpdateBaseUserArgs) GetReq() *user.UpdateBaseUserRequest {
	if !p.IsSetReq() {
		return UpdateBaseUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateBaseUserArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdateBaseUserArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdateBaseUserResult struct {
	Success *user.UpdateBaseUserResponse
}

var UpdateBaseUserResult_Success_DEFAULT *user.UpdateBaseUserResponse

func (p *UpdateBaseUserResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.UpdateBaseUserResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateBaseUserResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateBaseUserResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateBaseUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateBaseUserResult) Unmarshal(in []byte) error {
	msg := new(user.UpdateBaseUserResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateBaseUserResult) GetSuccess() *user.UpdateBaseUserResponse {
	if !p.IsSetSuccess() {
		return UpdateBaseUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateBaseUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.UpdateBaseUserResponse)
}

func (p *UpdateBaseUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateBaseUserResult) GetResult() interface{} {
	return p.Success
}

func updateUserPasswordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.UpdateUserPasswordRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).UpdateUserPassword(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *UpdateUserPasswordArgs:
		success, err := handler.(user.UserService).UpdateUserPassword(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateUserPasswordResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newUpdateUserPasswordArgs() interface{} {
	return &UpdateUserPasswordArgs{}
}

func newUpdateUserPasswordResult() interface{} {
	return &UpdateUserPasswordResult{}
}

type UpdateUserPasswordArgs struct {
	Req *user.UpdateUserPasswordRequest
}

func (p *UpdateUserPasswordArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.UpdateUserPasswordRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateUserPasswordArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateUserPasswordArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateUserPasswordArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateUserPasswordArgs) Unmarshal(in []byte) error {
	msg := new(user.UpdateUserPasswordRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateUserPasswordArgs_Req_DEFAULT *user.UpdateUserPasswordRequest

func (p *UpdateUserPasswordArgs) GetReq() *user.UpdateUserPasswordRequest {
	if !p.IsSetReq() {
		return UpdateUserPasswordArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateUserPasswordArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdateUserPasswordArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdateUserPasswordResult struct {
	Success *user.UpdateUserPasswordResponse
}

var UpdateUserPasswordResult_Success_DEFAULT *user.UpdateUserPasswordResponse

func (p *UpdateUserPasswordResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.UpdateUserPasswordResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateUserPasswordResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateUserPasswordResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateUserPasswordResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateUserPasswordResult) Unmarshal(in []byte) error {
	msg := new(user.UpdateUserPasswordResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateUserPasswordResult) GetSuccess() *user.UpdateUserPasswordResponse {
	if !p.IsSetSuccess() {
		return UpdateUserPasswordResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateUserPasswordResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.UpdateUserPasswordResponse)
}

func (p *UpdateUserPasswordResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateUserPasswordResult) GetResult() interface{} {
	return p.Success
}

func updateUserBalanceHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.UpdateUserBalanceRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).UpdateUserBalance(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *UpdateUserBalanceArgs:
		success, err := handler.(user.UserService).UpdateUserBalance(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateUserBalanceResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newUpdateUserBalanceArgs() interface{} {
	return &UpdateUserBalanceArgs{}
}

func newUpdateUserBalanceResult() interface{} {
	return &UpdateUserBalanceResult{}
}

type UpdateUserBalanceArgs struct {
	Req *user.UpdateUserBalanceRequest
}

func (p *UpdateUserBalanceArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.UpdateUserBalanceRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateUserBalanceArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateUserBalanceArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateUserBalanceArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateUserBalanceArgs) Unmarshal(in []byte) error {
	msg := new(user.UpdateUserBalanceRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateUserBalanceArgs_Req_DEFAULT *user.UpdateUserBalanceRequest

func (p *UpdateUserBalanceArgs) GetReq() *user.UpdateUserBalanceRequest {
	if !p.IsSetReq() {
		return UpdateUserBalanceArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateUserBalanceArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdateUserBalanceArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdateUserBalanceResult struct {
	Success *user.UpdateUserBalanceResponse
}

var UpdateUserBalanceResult_Success_DEFAULT *user.UpdateUserBalanceResponse

func (p *UpdateUserBalanceResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.UpdateUserBalanceResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateUserBalanceResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateUserBalanceResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateUserBalanceResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateUserBalanceResult) Unmarshal(in []byte) error {
	msg := new(user.UpdateUserBalanceResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateUserBalanceResult) GetSuccess() *user.UpdateUserBalanceResponse {
	if !p.IsSetSuccess() {
		return UpdateUserBalanceResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateUserBalanceResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.UpdateUserBalanceResponse)
}

func (p *UpdateUserBalanceResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateUserBalanceResult) GetResult() interface{} {
	return p.Success
}

func updateUserStatusHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.UpdateUserStatusRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).UpdateUserStatus(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *UpdateUserStatusArgs:
		success, err := handler.(user.UserService).UpdateUserStatus(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateUserStatusResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newUpdateUserStatusArgs() interface{} {
	return &UpdateUserStatusArgs{}
}

func newUpdateUserStatusResult() interface{} {
	return &UpdateUserStatusResult{}
}

type UpdateUserStatusArgs struct {
	Req *user.UpdateUserStatusRequest
}

func (p *UpdateUserStatusArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.UpdateUserStatusRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateUserStatusArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateUserStatusArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateUserStatusArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateUserStatusArgs) Unmarshal(in []byte) error {
	msg := new(user.UpdateUserStatusRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateUserStatusArgs_Req_DEFAULT *user.UpdateUserStatusRequest

func (p *UpdateUserStatusArgs) GetReq() *user.UpdateUserStatusRequest {
	if !p.IsSetReq() {
		return UpdateUserStatusArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateUserStatusArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdateUserStatusArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdateUserStatusResult struct {
	Success *user.UpdateUserStatusResponse
}

var UpdateUserStatusResult_Success_DEFAULT *user.UpdateUserStatusResponse

func (p *UpdateUserStatusResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.UpdateUserStatusResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateUserStatusResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateUserStatusResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateUserStatusResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateUserStatusResult) Unmarshal(in []byte) error {
	msg := new(user.UpdateUserStatusResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateUserStatusResult) GetSuccess() *user.UpdateUserStatusResponse {
	if !p.IsSetSuccess() {
		return UpdateUserStatusResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateUserStatusResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.UpdateUserStatusResponse)
}

func (p *UpdateUserStatusResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateUserStatusResult) GetResult() interface{} {
	return p.Success
}

func getUserInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.GetUserInfoRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).GetUserInfo(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetUserInfoArgs:
		success, err := handler.(user.UserService).GetUserInfo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetUserInfoResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetUserInfoArgs() interface{} {
	return &GetUserInfoArgs{}
}

func newGetUserInfoResult() interface{} {
	return &GetUserInfoResult{}
}

type GetUserInfoArgs struct {
	Req *user.GetUserInfoRequest
}

func (p *GetUserInfoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.GetUserInfoRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetUserInfoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetUserInfoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetUserInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetUserInfoArgs) Unmarshal(in []byte) error {
	msg := new(user.GetUserInfoRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetUserInfoArgs_Req_DEFAULT *user.GetUserInfoRequest

func (p *GetUserInfoArgs) GetReq() *user.GetUserInfoRequest {
	if !p.IsSetReq() {
		return GetUserInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetUserInfoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetUserInfoResult struct {
	Success *user.GetUserInfoResponse
}

var GetUserInfoResult_Success_DEFAULT *user.GetUserInfoResponse

func (p *GetUserInfoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.GetUserInfoResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetUserInfoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetUserInfoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetUserInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetUserInfoResult) Unmarshal(in []byte) error {
	msg := new(user.GetUserInfoResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserInfoResult) GetSuccess() *user.GetUserInfoResponse {
	if !p.IsSetSuccess() {
		return GetUserInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.GetUserInfoResponse)
}

func (p *GetUserInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserInfoResult) GetResult() interface{} {
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

func (p *kClient) CreateUser(ctx context.Context, Req *user.CreateUserRequest) (r *user.CreateUserResponse, err error) {
	var _args CreateUserArgs
	_args.Req = Req
	var _result CreateUserResult
	if err = p.c.Call(ctx, "CreateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteUser(ctx context.Context, Req *user.DeleteUserRequest) (r *user.DeleteUserResponse, err error) {
	var _args DeleteUserArgs
	_args.Req = Req
	var _result DeleteUserResult
	if err = p.c.Call(ctx, "DeleteUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateBaseUser(ctx context.Context, Req *user.UpdateBaseUserRequest) (r *user.UpdateBaseUserResponse, err error) {
	var _args UpdateBaseUserArgs
	_args.Req = Req
	var _result UpdateBaseUserResult
	if err = p.c.Call(ctx, "UpdateBaseUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateUserPassword(ctx context.Context, Req *user.UpdateUserPasswordRequest) (r *user.UpdateUserPasswordResponse, err error) {
	var _args UpdateUserPasswordArgs
	_args.Req = Req
	var _result UpdateUserPasswordResult
	if err = p.c.Call(ctx, "UpdateUserPassword", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateUserBalance(ctx context.Context, Req *user.UpdateUserBalanceRequest) (r *user.UpdateUserBalanceResponse, err error) {
	var _args UpdateUserBalanceArgs
	_args.Req = Req
	var _result UpdateUserBalanceResult
	if err = p.c.Call(ctx, "UpdateUserBalance", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateUserStatus(ctx context.Context, Req *user.UpdateUserStatusRequest) (r *user.UpdateUserStatusResponse, err error) {
	var _args UpdateUserStatusArgs
	_args.Req = Req
	var _result UpdateUserStatusResult
	if err = p.c.Call(ctx, "UpdateUserStatus", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserInfo(ctx context.Context, Req *user.GetUserInfoRequest) (r *user.GetUserInfoResponse, err error) {
	var _args GetUserInfoArgs
	_args.Req = Req
	var _result GetUserInfoResult
	if err = p.c.Call(ctx, "GetUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
