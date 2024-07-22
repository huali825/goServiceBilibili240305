package domain

import (
	"time"
)

// User 领域对象，是 DDD 中的 entity
// BO(business object)
type User struct {
	Id       int64
	Email    string
	Password string
	Phone    string

	// 不要组合，万一你将来可能还有 DingDingInfo，里面有同名字段 UnionID
	WechatInfo WechatInfo
	Ctime      time.Time
}

//type Address struct {
//}
