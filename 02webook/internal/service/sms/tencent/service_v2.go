package tencent

import (
	"context"
	"fmt"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

type ServiceV1 struct {
	client   *sms.Client
	appId    *string
	signName *string
}

func NewServiceV1(c *sms.Client, appId string, signName string) *Service {
	return &Service{
		client:   c,
		appId:    toPtr[string](appId),
		signName: toPtr[string](signName),
	}
}

func (s *ServiceV1) Send(ctx context.Context, tplId string, args map[string]string, numbers ...string) error {
	req := sms.NewSendSmsRequest()
	req.PhoneNumberSet = toStringPtrSlice(numbers)
	req.SmsSdkAppId = s.appId
	// ctx 继续往下传
	req.SetContext(ctx)
	req.TemplateParamSet = mapToStringPtrSlice(args)
	req.TemplateId = toPtr[string](tplId)
	req.SignName = s.signName
	resp, err := s.client.SendSms(req)
	if err != nil {
		return err
	}
	for _, status := range resp.Response.SendStatusSet {
		if status.Code == nil || *(status.Code) != "Ok" {
			return fmt.Errorf("发送失败，code: %s, 原因：%s",
				*status.Code, *status.Message)
		}
	}
	return nil
}

// string切片转换string 指针切片
func toStringPtrSlice(src []string) []*string {
	dst := make([]*string, len(src))
	for i, s := range src {
		dst[i] = &s
	}
	return dst
}

// map转换string 指针切片（当key没用的时候才用此方法）
func mapToStringPtrSlice(src map[string]string) []*string {
	dst := make([]*string, len(src))
	var i int
	for _, v := range src {
		dst[i] = &v
		i++
	}
	return dst
}

func toPtr[T any](t T) *T {
	return &t
}
