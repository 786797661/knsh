package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Greeter struct {
	Hello string
}

type GreeterRepo interface {
	CreateGreeter(context.Context, *Greeter) error
	UpdateGreeter(context.Context, *Greeter) error
}

type GreeterRealworld struct {
	repo GreeterRepo
	log  *log.Helper
}

func NewGreeterRealworld(repo GreeterRepo, logger log.Logger) *GreeterRealworld {
	return &GreeterRealworld{repo: repo, log: log.NewHelper(logger)}
}
