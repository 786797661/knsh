package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "knsh/api/realworld/v1"
)

type RealworldRepo interface {
	Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error)
}

type GreeterRealworld struct {
	repo RealworldRepo
	log  *log.Helper
}

func NewGreeterRealworld(repo RealworldRepo, logger log.Logger) *GreeterRealworld {
	return &GreeterRealworld{repo: repo, log: log.NewHelper(logger)}
}
