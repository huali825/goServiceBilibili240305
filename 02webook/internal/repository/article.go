package repository

import (
	"context"
	"go20240218/02webook/internal/domain"
	"go20240218/02webook/internal/repository/dao"
)

type ArticleRepository interface {
	Create(ctx context.Context, art domain.Article) (int64, error)
}

type CachedArticleRepository struct {
	dao dao.ArticleDAO
}

func (c *CachedArticleRepository) Create(ctx context.Context, art domain.Article) (int64, error) {
	return c.dao.Insert(ctx, dao.Article{
		Title:    art.Title,
		Content:  art.Content,
		AuthorId: art.Author.Id,
	})
}

func NewArticleRepository(dao dao.ArticleDAO) ArticleRepository {
	return &CachedArticleRepository{
		dao: dao,
	}
}
