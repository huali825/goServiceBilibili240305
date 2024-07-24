package integration

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go20240218/02webook/internal/integration/startup"
	"go20240218/02webook/internal/repository/dao"
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
				var art dao.Article
				err := s.db.Where("id=?", 1).First(&art).Error
				assert.NoError(t, err)
				assert.True(t, art.Ctime > 0)
				assert.True(t, art.Utime > 0)
				art.Ctime = 0
				art.Utime = 0
				assert.Equal(t, dao.Article{
					Id:       1,
					Title:    "标题001",
					Content:  "内容001",
					AuthorId: 123,
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
				err := s.db.Create(dao.Article{
					Id:       2,
					Title:    "标题002",
					Content:  "内容002",
					AuthorId: 123,
					//跟时间有关的测试不是逼不得已不用time.now()
					Ctime: 123,
					Utime: 234,
				}).Error
				assert.NoError(t, err)
			},
			after: func(t *testing.T) {
				//验证数据库
				var art dao.Article
				err := s.db.Where("id=?", 2).First(&art).Error
				assert.NoError(t, err)
				//assert.True(t, art.Ctime > 0)
				assert.True(t, art.Utime > 234)
				art.Utime = 0
				assert.Equal(t, dao.Article{
					Id:       2,
					Title:    "标题003",
					Content:  "内容003",
					Ctime:    123,
					AuthorId: 123,
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
				err := s.db.Create(dao.Article{
					Id:      3,
					Title:   "标题004",
					Content: "内容004",
					// 123在修改124的数据
					AuthorId: 124,
					//跟时间有关的测试不是逼不得已不用time.now()
					Ctime: 123,
					Utime: 234,
				}).Error
				assert.NoError(t, err)
			},
			after: func(t *testing.T) {
				//验证数据库
				var art dao.Article
				err := s.db.Where("id=?", 3).First(&art).Error
				assert.NoError(t, err)
				//assert.True(t, art.Ctime > 0)
				//assert.True(t, art.Utime > 234)
				assert.Equal(t, dao.Article{
					Id:       3,
					Title:    "标题004",
					Content:  "内容004",
					AuthorId: 124,

					Ctime: 123,
					Utime: 234,
				}, art)
			},
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

func TestArticle(t *testing.T) {
	suite.Run(t, new(ArticleTestSuite))
}

type Article struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
type Result[T any] struct {
	// 这个叫做业务错误码
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}
