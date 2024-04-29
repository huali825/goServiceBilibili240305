package intergration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go20240218/01webook/internal/web"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandler_e2e_SendSMSLoginCode(t *testing.T) {
	server := InitWebServer()
	//rdb := ioc.InitRedis()
	testCases := []struct {
		name string

		//准备数据
		before func(t *testing.T)

		//验证数据 redis的数据对不对
		after func(t *testing.T)

		phone string

		reqBody string

		// 构造请求，预期中输入
		//reqBuilder func(t *testing.T) *http.Request

		// 预期中的输出
		wantCode int
		wantBody web.Result
	}{
		{
			name: "发送成功",
			before: func(t *testing.T) {
				//nothing
			},
			after: func(t *testing.T) {

			},
			phone: "133434346464",
			//			reqBody: `
			//{
			//	"phone": "13343434646"
			//}
			//`,
			wantCode: 200,
			wantBody: web.Result{
				Msg: "发送成功",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// 构造 handler
			//userSvc, codeSvc := tc.mock(ctrl)
			//hdl := NewUserHandler(userSvc, codeSvc)

			// 准备服务器，注册路由
			//server := gin.Default()
			//hdl.RegisterRoutes(server)

			// 准备Req和记录的 recorder
			//req := tc.reqBuilder(t)
			recorder := httptest.NewRecorder()

			// 准备Req和记录的 recorder
			req, _ := http.NewRequest(http.MethodPost,
				"/users/login_sms/code/send",
				bytes.NewReader([]byte(fmt.Sprintf(`{"phone": "%s"}`, tc.phone))))
			req.Header.Set("Content-Type", "application/json")

			// 执行
			server.ServeHTTP(recorder, req)

			// 断言结果
			assert.Equal(t, tc.wantCode, recorder.Code)

			var res web.Result
			err := json.NewDecoder(recorder.Body).Decode(&res)
			assert.NoError(t, err)
			assert.Equal(t, tc.wantBody, res)
			//assert.Equal(t, tc.wantBody, recorder.Body.String())
		})
	}
}