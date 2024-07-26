package service

import (
	"context"
	"go20240218/02webook/internal/repository/article"
	"go20240218/02webook/pkg/logger"
	//"errors"
	"go20240218/02webook/internal/domain"
	"go20240218/02webook/internal/repository"
)

type ArticleService interface {
	Save(ctx context.Context, art domain.Article) (int64, error)
	Publish(ctx context.Context, art domain.Article) (int64, error)
	PublishV1(ctx context.Context, art domain.Article) (int64, error)
	//Update(ctx context.Context, art domain.Article) error
}

type articleService struct {
	repo repository.ArticleRepository

	// V1 用下面的就不能用上面的
	author article.ArticleAuthorRepository
	reader article.ArticleReaderRepository
	l      logger.LoggerV1
}

func NewArticleService(repo repository.ArticleRepository) ArticleService {
	return &articleService{
		repo: repo,
	}
}
func NewArticleServiceV1(author article.ArticleAuthorRepository, reader article.ArticleReaderRepository) ArticleService {
	return &articleService{
		author: author,
		reader: reader,
	}
}

func (a *articleService) Publish(ctx context.Context, art domain.Article) (int64, error) {
	panic("implement me")
}

func (a *articleService) PublishV1(ctx context.Context, art domain.Article) (int64, error) {
	id, err := a.author.Create(ctx, art)
	if err != nil {
		return 0, err
	}
	art.Id = id
	return a.reader.Save(ctx, art)

	//var (
	//	id  = art.Id
	//	err error
	//)
	//if art.Id > 0 {
	//	err = a.author.Update(ctx, art)
	//} else {
	//	id, err = a.author.Create(ctx, art)
	//}
	//if err != nil {
	//	return 0, err
	//}
	//art.Id = id
	//for i := 0; i < 3; i++ {
	//	time.Sleep(time.Second * time.Duration(i))
	//	id, err = a.reader.Save(ctx, art)
	//	if err == nil {
	//		break
	//	}
	//	a.l.Error("部分失败，保存到线上库失败",
	//		logger.Int64("art_id", art.Id),
	//		logger.Error(err))
	//}
	//if err != nil {
	//	a.l.Error("部分失败，重试彻底失败",
	//		logger.Int64("art_id", art.Id),
	//		logger.Error(err))
	//	// 接入你的告警系统，手工处理一下
	//	// 走异步，我直接保存到本地文件
	//	// 走 Canal
	//	// 打 MQ
	//}
	//return id, err
}

func (a *articleService) Save(ctx context.Context, art domain.Article) (int64, error) {
	if art.Id > 0 {
		err := a.repo.Update(ctx, art)
		return art.Id, err
	}
	return a.repo.Create(ctx, art)
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
