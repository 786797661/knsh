package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	v1 "knsh/api/realworld/v1"
	"knsh/internal/biz"
)

type realworldRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewRealworldRepo(data *Data, logger log.Logger) biz.RealworldRepo {
	return &realworldRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type User struct {
	Id       int64  `gorm:"primarykey"`
	Username string `gorm:"column:username"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

func (User) TableName() string {
	return "scholar_userinfo"
}

func (dac *realworldRepo) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	fmt.Println(req)
	var user User
	fmt.Println("进入查询")
	dac.data.db.Where("email = ? and password =?", req.UserInfo.Email, req.UserInfo.Password).First(&user)

	fmt.Println("执行搜索")
	userInfo := v1.UserInfo{
		Email:    user.Email,
		Username: user.Username,
	}
	res := v1.LoginReply{
		User: &userInfo,
	}
	return &res, nil
}
