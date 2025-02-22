package service

import (
	"context"
	"testing"

	"github.com/PiaoAdmin/gomall/app/user/biz/dal"
	user "github.com/PiaoAdmin/gomall/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
)

func TestDeleteUser_Run(t *testing.T) {
	_ = godotenv.Load("../../.env")
	dal.Init()
	ctx := context.Background()
	s := NewDeleteUserService(ctx)
	// init req and assert value

	req := &user.DeleteUserRequest{
		UserId: 1891718274337476608,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
