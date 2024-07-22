package cloopen

import (
	"context"
	"os"
	"testing"

	"github.com/cloopen/go-sms-sdk/cloopen"
)

func TestSender(t *testing.T) {
	accountSId, ok := os.LookupEnv("SMS_ACCOUNT_SID")
	if !ok {
		t.Fatal()
	}
	authToken, ok := os.LookupEnv("SMS_AUTH_TOKEN")
	appId, ok := os.LookupEnv("APP_ID")
	number, ok := os.LookupEnv("NUMBER")

	cfg := cloopen.DefaultConfig().
		WithAPIAccount(accountSId).
		WithAPIToken(authToken)
	c := cloopen.NewJsonClient(cfg).SMS()

	s := NewService(c, appId)

	tests := []struct {
		name    string
		tplId   string
		data    []string
		numbers []string
		wantErr error
	}{
		{
			name:  "发送验证码",
			tplId: "1",
			data:  []string{"1234", "5"},
			// 改成你的手机号码
			numbers: []string{number},
		},
	}
	for _, tt := range tests {
		err := s.Send(context.Background(), tt.tplId, tt.data, tt.numbers...)
		if err != nil {
			t.Fatal(err)
		}
	}
}
