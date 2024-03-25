package repository

import (
	"context"
	"github.com/gin-gonic/gin"
	"go20240218/01webook/internal/domain"
	"go20240218/01webook/internal/repository/dao"
)

var (
	ErrUserNotFound       = dao.ErrUserNotFound
	ErrUserDuplicateEmail = dao.ErrUserDuplicateEmail
)

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{dao: dao}
}

func (ur UserRepository) FindByEmail(
	ctx *gin.Context, email string) (domain.User, error) {
	u, err := ur.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
	}, nil
}

func (ur UserRepository) Create(ctx context.Context, u domain.User) error {
	return ur.dao.Insert(ctx, dao.DaoisUser{
		//Id:         0,
		Email:    u.Email,
		Password: u.Password,
		//CreatTime:  0,
		//UpdateTime: 0,
	})
}
