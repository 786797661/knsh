package service

import (
	"github.com/go-kratos/kratos/v2/log"
	v1 "knsh/api/realworld/v1"
	"knsh/internal/biz"
)

// GreeterService is a greeter service.
type RealworldService struct {
	v1.UnimplementedRealworldServer
	uc  *biz.GreeterRealworld
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewrealworldService(uc *biz.GreeterRealworld, logger log.Logger) *RealworldService {
	return &RealworldService{uc: uc, log: log.NewHelper(logger)}
}
