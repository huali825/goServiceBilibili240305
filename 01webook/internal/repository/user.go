package repository

import (
	"context"
	"go20240218/01webook/internal/domain"
	"go20240218/01webook/internal/repository/dao"
)

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{dao: dao}
}

func (ur UserRepository) Create(ctx context.Context, u domain.User) error {
	return ur.dao.Insert(ctx, dao.DaoUser{
		//Id:         0,
		Email:    u.Email,
		Password: u.Password,
		//CreatTime:  0,
		//UpdateTime: 0,
	})
}
