package service

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go20240218/02webook/internal/domain"
	"go20240218/02webook/internal/repository"
	repomocks "go20240218/02webook/internal/repository/mocks"
	"go20240218/02webook/pkg/logger"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"
)

func Test_userService_Login(t *testing.T) {
	// 做成一个测试用例都用到的时间
	now := time.Now()

	testCases := []struct {
		name string
		mock func(ctrl *gomock.Controller) repository.UserRepository

		// 输入
		//ctx      context.Context
		email    string
		password string

		// 输出
		wantUser domain.User
		wantErr  error
	}{
		{
			name: "登录成功", // 用户名和密码是对的
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().FindByEmail(gomock.Any(), "123@qq.com").
					Return(domain.User{
						Email:    "123@qq.com",
						Password: "$2a$10$MN9ZKKIbjLZDyEpCYW19auY7mvOG9pcpiIcUUoZZI6pA6OmKZKOVi",
						Phone:    "15212345678",
						Ctime:    now,
					}, nil)
				return repo
			},
			email:    "123@qq.com",
			password: "hello#world123",

			wantUser: domain.User{
				Email:    "123@qq.com",
				Password: "$2a$10$MN9ZKKIbjLZDyEpCYW19auY7mvOG9pcpiIcUUoZZI6pA6OmKZKOVi",
				Phone:    "15212345678",
				Ctime:    now,
			},
			wantErr: nil,
		},
		{
			name: "用户不存在",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().FindByEmail(gomock.Any(), "123@qq.com").
					Return(domain.User{}, repository.ErrUserNotFound)
				return repo
			},
			email:    "123@qq.com",
			password: "hello#world123",

			wantUser: domain.User{},
			wantErr:  ErrInvalidUserOrPassword,
		},
		{
			name: "DB错误",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().FindByEmail(gomock.Any(), "123@qq.com").
					Return(domain.User{}, errors.New("mock db 错误"))
				return repo
			},
			email:    "123@qq.com",
			password: "hello#world123",

			wantUser: domain.User{},
			wantErr:  errors.New("mock db 错误"),
		},
		{
			name: "密码不对",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().FindByEmail(gomock.Any(), "123@qq.com").
					Return(domain.User{
						Email:    "123@qq.com",
						Password: "$2a$10$MN9ZKKIbjLZDyEpCYW19auY7mvOG9pcpiIcUUoZZI6pA6OmKZKOVi",
						Phone:    "15212345678",
						Ctime:    now,
					}, nil)
				return repo
			},
			email:    "123@qq.com",
			password: "112443rsdffhello#world123",

			wantUser: domain.User{},
			wantErr:  ErrInvalidUserOrPassword,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			// 具体的测试代码
			svc := NewUserService(tc.mock(ctrl), &logger.NopLogger{})
			u, err := svc.Login(context.Background(), tc.email, tc.password)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantUser, u)
		})
	}
}

func TestEncrypted(t *testing.T) {
	res, err := bcrypt.GenerateFromPassword([]byte("hello#world123"), bcrypt.DefaultCost)
	if err == nil {
		t.Log(string(res))
	}
}
