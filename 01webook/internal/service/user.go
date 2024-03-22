package service

import (
	"context"
	"go20240218/01webook/internal/domain"
	"go20240218/01webook/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) Signup(ctx context.Context, u domain.User) error {
	//加密
	return svc.repo.Create(ctx, u)
}
