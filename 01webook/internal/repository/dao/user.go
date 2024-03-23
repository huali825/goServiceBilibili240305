package dao

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
}

func (d UserDAO) Insert(ctx context.Context, u DaoUser) error {
	now := time.Now().UnixMilli() //毫秒数
	u.CreatTime = now
	u.UpdateTime = now

	return d.db.WithContext(ctx).Create(&u).Error //gorm
	//panic("zuo ye wu ")
}

// User 直接对应 数据库表结构
type DaoUser struct {
	Id       int64  `gorm:"primaryKey,autoIncrement"`
	Email    string `gorm:"unique"`
	Password string //`gorm:"password"`

	CreatTime  int64 //`gorm:"creat_time"`
	UpdateTime int64 //`gorm:"update_time"`
}
