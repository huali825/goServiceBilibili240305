package service

import (
	"context"
	//"errors"
	"go20240218/02webook/internal/domain"
	"go20240218/02webook/internal/repository"
)

type ArticleService interface {
	Save(ctx context.Context, art domain.Article) (int64, error)
	Publish(ctx context.Context, art domain.Article) (int64, error)
	//Update(ctx context.Context, art domain.Article) error
}

type articleService struct {
	repo repository.ArticleRepository
}

func NewArticleService(repo repository.ArticleRepository) ArticleService {
	return &articleService{
		repo: repo,
	}
}

func (as *articleService) Publish(ctx context.Context, art domain.Article) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (as *articleService) Save(ctx context.Context, art domain.Article) (int64, error) {
	if art.Id > 0 {
		err := as.repo.Update(ctx, art)
		return art.Id, err
	}
	return as.repo.Create(ctx, art)
}

//func (as *articleService) Update(ctx context.Context, art domain.Article) error {
//	// 只要你不更新 author_id
//	// 但是性能比较差
//	artInDB, err := as.repo.FindByID(ctx, art.Id)
//	if err != nil {
//		return err
//	}
//	if art.Author.Id != artInDB.Author.Id {
//		return errors.New("更新别人的数据")
//	}
//	return as.repo.Update(ctx, art)
//}
