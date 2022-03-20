package service

import (
	"context"
	v1 "knsh/api/realworld/v1"
)

func (realworld *RealworldService) Login(context.Context, *v1.LoginRequest) (*v1.LoginReply, error) {
	userInfo := v1.UserInfo{
		Email:    "786797661@qq.com",
		Username: "yzc",
		Token:    "",
	}

	return &v1.LoginReply{User: &userInfo}, nil
}
