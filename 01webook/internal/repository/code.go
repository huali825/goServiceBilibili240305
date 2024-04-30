package repository

import (
	"context"
	"fmt"
	"go20240218/01webook/internal/repository/cache"
)

var ErrCodeSendTooMany = cache.ErrCodeSendTooMany
var ErrCodeVerifyTooMany = cache.ErrCodeVerifyTooMany

type CodeRepository interface {
	Set(ctx context.Context, biz, phone, code string) error
	Verify(ctx context.Context, biz, phone, code string) (bool, error)
	isRepoCode(s string) error
}

type CachedCodeRepository struct {
	cache cache.CodeCache
}

func NewCodeRepository(c cache.CodeCache) CodeRepository {
	return &CachedCodeRepository{
		cache: c,
	}
}

func (c *CachedCodeRepository) isRepoCode(s string) error {
	fmt.Println(s)
	return nil
}

func (c *CachedCodeRepository) Set(ctx context.Context, biz, phone, code string) error {
	return c.cache.Set(ctx, biz, phone, code)
}

func (c *CachedCodeRepository) Verify(ctx context.Context, biz, phone, code string) (bool, error) {
	return c.cache.Verify(ctx, biz, phone, code)
}
