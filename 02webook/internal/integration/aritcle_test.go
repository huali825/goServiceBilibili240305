package integration

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go20240218/02webook/internal/domain"
	"go20240218/02webook/internal/integration/startup"
	daoArticle "go20240218/02webook/internal/repository/dao/article"
	ijwt "go20240218/02webook/internal/web/jwt"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ArticleTestSuite struct {
	//单元测试套件的写法
	suite.Suite
	server *gin.Engine
	db     *gorm.DB
}

// TearDownTest 每一个都会执行
func (s *ArticleTestSuite) TearDownTest() {
	// 清空所有数据，并且自增主键恢复到 1
	s.db.Exec("TRUNCATE TABLE articles")
	s.db.Exec("TRUNCATE TABLE publish_articles")
}

// 在所有测试开始之前 初始化一些内容
func (s *ArticleTestSuite) SetupSuite() {
	//直接使用
	//s.server = startup.InitWebServer()

	//定制化
	s.server = gin.Default()
	s.server.Use(func(ctx *gin.Context) {
		ctx.Set("claims", &ijwt.UserClaims{
			Uid: 123,
		})
	})
	s.db = startup.InitTestDB()
	artHdl := startup.InitArticleHandler()
	// 注册好了路由
	artHdl.RegisterRoutes(s.server)
}

func (s *ArticleTestSuite) TestAbc() {
	s.T().Log("hello ,这是测试套件")
}

func (s *ArticleTestSuite) TestEdit() {
	t := s.T()
	testCases := []struct {
		name string

		//集成测试准备数据
		before func(t *testing.T)
		after  func(t *testing.T)

		req Article

		wantCode int
		wantRes  Result[int64]
	}{
		{
			name: "新建帖子-保存成功",
			before: func(t *testing.T) {
			},
			after: func(t *testing.T) {
				//验证数据库
				var art daoArticle.Article
				err := s.db.Where("id=?", 1).First(&art).Error
				assert.NoError(t, err)
				assert.True(t, art.Ctime > 0)
				assert.True(t, art.Utime > 0)
				art.Ctime = 0
				art.Utime = 0
				assert.Equal(t, daoArticle.Article{
					Id:       1,
					Title:    "标题001",
					Content:  "内容001",
					AuthorId: 123,
					Status:   domain.ArticleStatusUnpublished.ToUint8(),
				}, art)
			},
			req: Article{
				Title:   "标题001",
				Content: "内容001",
			},
			wantCode: http.StatusOK,
			wantRes: Result[int64]{
				Data: 1,
				Msg:  "OK",
			},
		},
		{
			name: "修改帖子-保存成功",
			before: func(t *testing.T) {
				//提前准备数据
				err := s.db.Create(daoArticle.Article{
					Id:       2,
					Title:    "标题002",
					Content:  "内容002",
					AuthorId: 123,
					//跟时间有关的测试不是逼不得已不用time.now()
					Ctime:  123,
					Utime:  234,
					Status: domain.ArticleStatusUnpublished.ToUint8(),
				}).Error
				assert.NoError(t, err)
			},
			after: func(t *testing.T) {
				//验证数据库
				var art daoArticle.Article
				err := s.db.Where("id=?", 2).First(&art).Error
				assert.NoError(t, err)
				//assert.True(t, art.Ctime > 0)
				assert.True(t, art.Utime > 234)
				art.Utime = 0
				assert.Equal(t, daoArticle.Article{
					Id:       2,
					Title:    "标题003",
					Content:  "内容003",
					Ctime:    123,
					AuthorId: 123,
					Status:   domain.ArticleStatusUnpublished.ToUint8(),
				}, art)
			},
			req: Article{
				Id:      2,
				Title:   "标题003",
				Content: "内容003",
			},
			wantCode: http.StatusOK,
			wantRes: Result[int64]{
				Data: 2,
				Msg:  "OK",
			},
		},
		{
			name: "修改别人的帖子",
			before: func(t *testing.T) {
				//提前准备数据
				err := s.db.Create(daoArticle.Article{
					Id:      3,
					Title:   "标题004",
					Content: "内容004",
					// 123在修改124的数据
					AuthorId: 124,
					//跟时间有关的测试不是逼不得已不用time.now()
					Ctime:  123,
					Utime:  234,
					Status: domain.ArticleStatusPublished.ToUint8(),
				}).Error
				assert.NoError(t, err)
			},
			after: func(t *testing.T) {
				//验证数据库
				var art daoArticle.Article
				err := s.db.Where("id=?", 3).First(&art).Error
				assert.NoError(t, err)
				//assert.True(t, art.Ctime > 0)
				//assert.True(t, art.Utime > 234)
				assert.Equal(t, daoArticle.Article{
					Id:       3,
					Title:    "标题004",
					Content:  "内容004",
					AuthorId: 124,

					Ctime:  123,
					Utime:  234,
					Status: domain.ArticleStatusPublished.ToUint8(),
				}, art)
			},

			//没有办法改的
			req: Article{
				Id:      3,
				Title:   "标题003222",
				Content: "内容003222",
			},
			wantCode: http.StatusOK,
			wantRes: Result[int64]{
				Code: 5,
				Msg:  "系统错误",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//构造请求
			//执行
			//验证结果
			tc.before(t)
			reqbody, err := json.Marshal(tc.req)
			assert.NoError(t, err)
			req, err := http.NewRequest(http.MethodPost,
				"/articles/edit", bytes.NewBuffer(reqbody))
			require.NoError(t, err)
			// 数据是 JSON 格式
			req.Header.Set("Content-Type", "application/json")
			// 这里你就可以继续使用 req

			resp := httptest.NewRecorder()
			// 这就是 HTTP 请求进去 GIN 框架的入口。
			// 当你这样调用的时候，GIN 就会处理这个请求
			// 响应写回到 resp 里
			s.server.ServeHTTP(resp, req)

			assert.Equal(t, tc.wantCode, resp.Code)
			if resp.Code != 200 {
				return
			}
			var webRes Result[int64]
			err = json.NewDecoder(resp.Body).Decode(&webRes)
			require.NoError(t, err)
			assert.Equal(t, tc.wantRes, webRes)
			tc.after(t)
		})
	}

}

func (s *ArticleTestSuite) TestPublish() {
	t := s.T()

	testCases := []struct {
		name string
		// 要提前准备数据
		before func(t *testing.T)
		// 验证并且删除数据
		after func(t *testing.T)
		req   Article

		// 预期响应
		wantCode   int
		wantResult Result[int64]
	}{
		{
			name: "新建帖子并发表",
			before: func(t *testing.T) {
				// 什么也不需要做
			},
			after: func(t *testing.T) {
				// 验证一下数据
				var art daoArticle.Article
				err := s.db.Where("author_id = ?", 123).First(&art).Error
				assert.NoError(t, err)
				// 确保已经生成了主键
				assert.True(t, art.Id > 0)
				assert.True(t, art.Ctime > 0)
				assert.True(t, art.Utime > 0)
				art.Ctime = 0
				art.Utime = 0
				art.Id = 0
				assert.Equal(t, daoArticle.Article{
					Title:    "hello，你好",
					Content:  "随便试试",
					AuthorId: 123,
					Status:   uint8(domain.ArticleStatusPublished),
				}, art)
				var publishedArt daoArticle.PublishArticle
				err = s.db.Where("author_id = ?", 123).First(&publishedArt).Error
				assert.NoError(t, err)
				assert.True(t, publishedArt.Id > 0)
				assert.True(t, publishedArt.Ctime > 0)
				assert.True(t, publishedArt.Utime > 0)
				publishedArt.Ctime = 0
				publishedArt.Utime = 0
				publishedArt.Id = 0
				assert.Equal(t, daoArticle.PublishArticle{
					Article: daoArticle.Article{
						Title:    "hello，你好",
						Content:  "随便试试",
						AuthorId: 123,
						Status:   uint8(domain.ArticleStatusPublished),
					},
				}, publishedArt)
			},
			req: Article{
				Title:   "hello，你好",
				Content: "随便试试",
			},
			wantCode: 200,
			wantResult: Result[int64]{
				Msg:  "OK",
				Data: 1,
			},
		},
		{
			// 制作库有，但是线上库没有
			name: "更新帖子并新发表",
			before: func(t *testing.T) {
				// 模拟已经存在的帖子
				err := s.db.Create(&daoArticle.Article{
					Id:       2,
					Title:    "我的标题",
					Content:  "我的内容",
					Ctime:    456,
					Utime:    234,
					AuthorId: 123,
					Status:   uint8(domain.ArticleStatusUnpublished),
				}).Error
				assert.NoError(t, err)
			},
			after: func(t *testing.T) {
				// 验证一下数据
				var art daoArticle.Article
				s.db.Where("id = ?", 2).First(&art)
				// 更新时间变了
				assert.True(t, art.Utime > 234)
				art.Utime = 0
				// 创建时间没变
				assert.Equal(t, daoArticle.Article{
					Id:       2,
					Ctime:    456,
					Status:   uint8(domain.ArticleStatusPublished),
					Content:  "新的内容",
					Title:    "新的标题",
					AuthorId: 123,
				}, art)

				var publishedArt daoArticle.PublishArticle
				s.db.Where("id = ?", 2).First(&publishedArt)
				assert.True(t, publishedArt.Ctime > 0)
				assert.True(t, publishedArt.Utime > 0)
				publishedArt.Ctime = 0
				publishedArt.Utime = 0
				assert.Equal(t, daoArticle.PublishArticle{
					Article: daoArticle.Article{
						Id:       2,
						Status:   uint8(domain.ArticleStatusPublished),
						Content:  "新的内容",
						Title:    "新的标题",
						AuthorId: 123,
					},
				}, publishedArt)
			},
			req: Article{
				Id:      2,
				Title:   "新的标题",
				Content: "新的内容",
			},
			wantCode: 200,
			wantResult: Result[int64]{
				Msg:  "OK",
				Data: 2,
			},
		},
		{
			name: "更新帖子，并且重新发表",
			before: func(t *testing.T) {
				art := daoArticle.Article{
					Id:       3,
					Title:    "我的标题",
					Content:  "我的内容",
					Ctime:    456,
					Utime:    234,
					AuthorId: 123,
					Status:   uint8(domain.ArticleStatusPublished),
				}
				err := s.db.Create(&art).Error
				assert.NoError(t, err)
				part := daoArticle.PublishArticle{
					Article: art,
				}
				err = s.db.Create(&part).Error
				assert.NoError(t, err)
			},
			after: func(t *testing.T) {
				var art daoArticle.Article
				err := s.db.Where("id = ?", 3).First(&art).Error
				assert.NoError(t, err)
				// 更新时间变了
				assert.True(t, art.Utime > 234)
				art.Utime = 0
				// 创建时间没变
				assert.Equal(t, daoArticle.Article{
					Id:       3,
					Ctime:    456,
					Status:   uint8(domain.ArticleStatusPublished),
					Content:  "新的内容",
					Title:    "新的标题",
					AuthorId: 123,
				}, art)

				var publishedArt daoArticle.PublishArticle
				err = s.db.Where("id = ?", 3).First(&publishedArt).Error
				assert.NoError(t, err)
				assert.True(t, publishedArt.Ctime > 0)
				assert.True(t, publishedArt.Utime > 0)
				publishedArt.Ctime = 0
				publishedArt.Utime = 0
				assert.Equal(t, daoArticle.PublishArticle{
					Article: daoArticle.Article{
						Id:       3,
						Status:   uint8(domain.ArticleStatusPublished),
						Content:  "新的内容",
						Title:    "新的标题",
						AuthorId: 123,
					},
				}, publishedArt)
			},
			req: Article{
				Id:      3,
				Title:   "新的标题",
				Content: "新的内容",
			},
			wantCode: 200,
			wantResult: Result[int64]{
				Msg:  "OK",
				Data: 3,
			},
		},
		{
			name: "更新别人的帖子，并且发表失败",
			before: func(t *testing.T) {
				art := daoArticle.Article{
					Id:      4,
					Title:   "我的标题",
					Content: "我的内容",
					Ctime:   456,
					Utime:   234,
					// 注意。这个 AuthorID 我们设置为另外一个人的ID
					AuthorId: 789,
				}
				s.db.Create(&art)
				part := daoArticle.PublishArticle{
					Article: daoArticle.Article{
						Id:       4,
						Title:    "我的标题",
						Content:  "我的内容",
						Ctime:    456,
						Utime:    234,
						AuthorId: 789,
					},
				}
				s.db.Create(&part)
			},
			after: func(t *testing.T) {
				// 更新应该是失败了，数据没有发生变化
				var art daoArticle.Article
				s.db.Where("id = ?", 4).First(&art)
				assert.Equal(t, "我的标题", art.Title)
				assert.Equal(t, "我的内容", art.Content)
				assert.Equal(t, int64(456), art.Ctime)
				assert.Equal(t, int64(234), art.Utime)
				assert.Equal(t, int64(789), art.AuthorId)

				var part daoArticle.PublishArticle
				// 数据没有变化
				s.db.Where("id = ?", 4).First(&part)
				assert.Equal(t, "我的标题", part.Title)
				assert.Equal(t, "我的内容", part.Content)
				assert.Equal(t, int64(789), part.AuthorId)
				// 创建时间没变
				assert.Equal(t, int64(456), part.Ctime)
				// 更新时间变了
				assert.Equal(t, int64(234), part.Utime)
			},
			req: Article{
				Id:      4,
				Title:   "新的标题",
				Content: "新的内容",
			},
			wantCode: 200,
			wantResult: Result[int64]{
				Code: 5,
				Msg:  "系统错误",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(t)
			data, err := json.Marshal(tc.req)
			// 不能有 error
			assert.NoError(t, err)
			req, err := http.NewRequest(http.MethodPost,
				"/articles/publish", bytes.NewReader(data))
			assert.NoError(t, err)
			req.Header.Set("Content-Type",
				"application/json")
			recorder := httptest.NewRecorder()

			s.server.ServeHTTP(recorder, req)
			code := recorder.Code
			assert.Equal(t, tc.wantCode, code)
			if code != http.StatusOK {
				return
			}
			// 反序列化为结果
			// 利用泛型来限定结果必须是 int64
			var result Result[int64]
			err = json.Unmarshal(recorder.Body.Bytes(), &result)
			assert.NoError(t, err)
			assert.Equal(t, tc.wantResult, result)
			tc.after(t)
		})
	}
}

func TestArticle(t *testing.T) {
	suite.Run(t, new(ArticleTestSuite))
}

type Article struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  uint8  `json:"status"`
}
type Result[T any] struct {
	// 这个叫做业务错误码
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}
