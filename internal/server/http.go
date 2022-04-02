package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	v1 "knsh/api/realworld/v1"
	"knsh/internal/conf"
	"knsh/internal/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, cau *conf.Auth, realworld *service.RealworldService, logger log.Logger) *http.Server {

	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			tracing.Server(), // 新增 tracing
			selector.Server(jwt.Server(func(token *jwtv4.Token) (interface{}, error) {
				return []byte(cau.JwtKey), nil
			})).Match(NewWhiteListMatcher()).Build(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterRealworldHTTPServer(srv, realworld)
	return srv
}

// NewWhiteListMatcher 白名单不需要token验证的接口
func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/shop.shop.v1.Shop/Captcha"] = struct{}{}
	whiteList["/shop.shop.v1.Shop/Login"] = struct{}{}
	whiteList["/shop.shop.v1.Shop/Register"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
