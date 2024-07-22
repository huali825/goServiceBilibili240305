package tencent

import (
	"context"
	"fmt"
	"github.com/ecodeclub/ekit"
	"github.com/ecodeclub/ekit/slice"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"go.uber.org/zap"
	mysms "go20240218/02webook/internal/service/sms"
	"go20240218/02webook/pkg/ratelimit"
)

type Service struct {
	appId    *string
	signName *string
	client   *sms.Client
	limiter  ratelimit.Limiter
}

func NewService(client *sms.Client, appId string,
	signName string, limiter ratelimit.Limiter) *Service {
	return &Service{
		client:   client,
		appId:    ekit.ToPtr[string](appId),
		signName: ekit.ToPtr[string](signName),
		limiter:  limiter,
	}
}

// Send 一个是 []*string
// 一个是 string，json 串
// biz 直接代表的就是 tplId
func (s *Service) Send(ctx context.Context,
	biz string, args []string, numbers ...string) error {
	req := sms.NewSendSmsRequest()
	req.SmsSdkAppId = s.appId
	req.SignName = s.signName
	req.TemplateId = ekit.ToPtr[string](biz)
	req.PhoneNumberSet = s.toStringPtrSlice(numbers)
	req.TemplateParamSet = s.toStringPtrSlice(args)
	resp, err := s.client.SendSms(req)
	zap.L().Debug("发送短信", zap.Any("req", req),
		zap.Any("resp", resp), zap.Error(err))
	if err != nil {
		return fmt.Errorf("腾讯短信服务发送失败 %w", err)
	}
	for _, status := range resp.Response.SendStatusSet {
		if status.Code == nil || *(status.Code) != "Ok" {
			return fmt.Errorf("发送短信失败 %s, %s ", *status.Code, *status.Message)
		}
	}
	return nil
}

func (s *Service) SendV1(ctx context.Context, tplId string, args []mysms.NamedArg, numbers ...string) error {
	req := sms.NewSendSmsRequest()
	req.SmsSdkAppId = s.appId
	req.SignName = s.signName
	req.TemplateId = ekit.ToPtr[string](tplId)
	req.PhoneNumberSet = s.toStringPtrSlice(numbers)
	req.TemplateParamSet = slice.Map[mysms.NamedArg, *string](args, func(idx int, src mysms.NamedArg) *string {
		return &src.Val
	})
	resp, err := s.client.SendSms(req)
	if err != nil {
		return err
	}
	for _, status := range resp.Response.SendStatusSet {
		if status.Code == nil || *(status.Code) != "Ok" {
			return fmt.Errorf("发送短信失败 %s, %s ", *status.Code, *status.Message)
		}
	}
	return nil
}

func (s *Service) toStringPtrSlice(src []string) []*string {
	return slice.Map[string, *string](src, func(idx int, src string) *string {
		return &src
	})
}
