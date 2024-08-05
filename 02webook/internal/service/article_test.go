package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go20240218/02webook/internal/domain"
	"go20240218/02webook/internal/repository/article"
	artrepomocks "go20240218/02webook/internal/repository/article/mocks"
	"testing"
)

func Test_articleService_Publish(t *testing.T) {

	testCases := []struct {
		name string
		mock func(ctrl *gomock.Controller) (article.ArticleAuthorRepository,
			article.ArticleReaderRepository)

		art domain.Article

		wantErr error
		wantId  int64
	}{
		{
			name: "新建发表成功",
			mock: func(ctrl *gomock.Controller) (article.ArticleAuthorRepository,
				article.ArticleReaderRepository) {
				author := artrepomocks.NewMockArticleAuthorRepository(ctrl)
				author.EXPECT().Create(gomock.Any(), domain.Article{
					Id:      0,
					Title:   "rep标题发表成功",
					Content: "rep内容发表成功",
					Author:  domain.Author{Id: 123},
				}).Return(int64(1), nil)
				reader := artrepomocks.NewMockArticleReaderRepository(ctrl)
				reader.EXPECT().Save(gomock.Any(), domain.Article{
					Id:      1,
					Title:   "rep标题发表成功",
					Content: "rep内容发表成功",
					Author:  domain.Author{Id: 123},
				}).Return(int64(1), nil)
				return author, reader
			},
			art: domain.Article{
				Title:   "rep标题发表成功",
				Content: "rep内容发表成功",
				Author:  domain.Author{Id: 123},
			},
			wantId:  1,
			wantErr: nil,
		},
		{
			name: "修改并发表成功",
			mock: func(ctrl *gomock.Controller) (article.ArticleAuthorRepository,
				article.ArticleReaderRepository) {
				author := artrepomocks.NewMockArticleAuthorRepository(ctrl)
				author.EXPECT().Update(gomock.Any(), domain.Article{
					Id:      2,
					Title:   "rep标题发表成功",
					Content: "rep内容发表成功",
					Author:  domain.Author{Id: 123},
				}).Return(nil)
				reader := artrepomocks.NewMockArticleReaderRepository(ctrl)
				reader.EXPECT().Save(gomock.Any(), domain.Article{
					Id:      2,
					Title:   "rep标题发表成功",
					Content: "rep内容发表成功",
					Author:  domain.Author{Id: 123},
				}).Return(int64(2), nil)
				return author, reader
			},
			art: domain.Article{
				Id:      2,
				Title:   "rep标题发表成功",
				Content: "rep内容发表成功",
				Author:  domain.Author{Id: 123},
			},
			wantId:  2,
			wantErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			author, reader := tc.mock(ctrl)
			svc := NewArticleServiceV1(author, reader)
			id, err := svc.PublishV1(context.Background(), tc.art)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantId, id)
		})
	}
}
