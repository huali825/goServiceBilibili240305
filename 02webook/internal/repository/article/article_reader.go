package article

import (
	"context"
	"go20240218/02webook/internal/domain"
)

type ArticleReaderRepository interface {
	// Save 有就更新，没有就新建，即 upsert 的语义
	Save(ctx context.Context, art domain.Article) (int64, error)
}
