package tencent

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"testing"
)

// 这个需要手动跑，也就是你需要在本地搞好这些环境变量
func TestSender(t *testing.T) {
	//secretId, ok := os.LookupEnv("SMS_SECRET_ID")
	//if !ok {
	//	t.Fatal()
	//}
	//secretKey, ok := os.LookupEnv("SMS_SECRET_KEY")
	secretId := "1400904371"
	secretKey := "76610c7c8349e3e68653dc768faff5db"

	c, err := sms.NewClient(common.NewCredential(secretId, secretKey),
		"ap-nanjing",
		profile.NewClientProfile())
	if err != nil {
		t.Fatal(err)
	}

	s := NewService(c, "1400842696", "妙影科技")

	testCases := []struct {
		name    string
		tplId   string
		params  []string
		numbers []string
		wantErr error
	}{
		{
			name:   "发送验证码",
			tplId:  "1877556",
			params: []string{"123456"},
			// 改成你的手机号码
			numbers: []string{"13343464683"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			er := s.Send(context.Background(), tc.tplId, tc.params, tc.numbers...)
			assert.Equal(t, tc.wantErr, er)
		})
	}
}
