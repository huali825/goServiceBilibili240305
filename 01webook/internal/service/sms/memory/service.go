package memory

import (
	"context"
	"fmt"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s Service) Send(ctx context.Context, tplId string, args []string, numbers ...string) error {
	fmt.Println("本地测试验证码已发送:", args)
	return nil
}
