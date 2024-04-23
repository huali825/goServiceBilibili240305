package sms

import "context"

//	type Service interface {
//		Send(ctx context.Context,
//			numbers []string,
//			appId string,
//			signature string,
//			tpl string,
//			args []string) error
//	}
type Service interface {
	Send(ctx context.Context, tplId string,
		args []string, numbers ...string) error
}
