package dao

import (
	"context"
	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
}

func (d UserDAO) Insert(ctx context.Context, u User) error {
	panic("zuo ye wu ")
}

// User 直接对应 数据库表结构
type User struct {
	Id       int64
	Email    string
	Password string

	CreatTime  int64
	UpdateTime int64
}
