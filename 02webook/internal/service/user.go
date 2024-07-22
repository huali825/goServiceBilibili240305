package service

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"go20240218/02webook/internal/domain"
	"go20240218/02webook/internal/repository"
	"go20240218/02webook/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

var ErrUserDuplicateEmail = repository.ErrUserDuplicate
var ErrInvalidUserOrPassword = errors.New("账号/邮箱或密码不对")

type UserService interface {
	Login(ctx context.Context, email, password string) (domain.User, error)
	SignUp(ctx context.Context, u domain.User) error
	FindOrCreate(ctx context.Context, phone string) (domain.User, error)
	FindOrCreateByWechat(ctx context.Context, wechatInfo domain.WechatInfo) (domain.User, error)
	Profile(ctx context.Context, id int64) (domain.User, error)
}

type userService struct {
	repo repository.UserRepository
	l    logger.LoggerV1
}

// NewUserService 我用的人，只管用，怎么初始化我不管，我一点都不关心如何初始化
func NewUserService(repo repository.UserRepository, l logger.LoggerV1) UserService {
	return &userService{
		repo: repo,
		l:    l,
	}
}

func NewUserServiceV1(repo repository.UserRepository, l *zap.Logger) UserService {
	return &userService{
		repo: repo,
		// 预留了变化空间
		//logger: zap.L(),
	}
}

//func NewUserServiceV1(f repository.UserRepositoryFactory) UserService {
//	return &userService{
//		// 我在这里，不同的 factory，会创建出来不同实现
//		repo: f.NewRepo(),
//	}
//}

func (svc *userService) Login(ctx context.Context, email, password string) (domain.User, error) {
	// 先找用户
	u, err := svc.repo.FindByEmail(ctx, email)
	if err == repository.ErrUserNotFound {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	if err != nil {
		return domain.User{}, err
	}
	// 比较密码了
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		// DEBUG
		return domain.User{}, ErrInvalidUserOrPassword
	}

	return u, nil
}

func (svc *userService) SignUp(ctx context.Context, u domain.User) error {
	// 你要考虑加密放在哪里的问题了
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	// 然后就是，存起来
	return svc.repo.Create(ctx, u)
}

func (svc *userService) FindOrCreate(ctx context.Context,
	phone string) (domain.User, error) {
	// 这时候，这个地方要怎么办？
	// 这个叫做快路径
	u, err := svc.repo.FindByPhone(ctx, phone)
	// 要判断，有咩有这个用户
	if err != repository.ErrUserNotFound {
		// 绝大部分请求进来这里
		// nil 会进来这里
		// 不为 ErrUserNotFound 的也会进来这里
		return u, err
	}
	// 这里，把 phone 脱敏之后打出来
	//zap.L().Info("用户未注册", zap.String("phone", phone))
	//svc.logger.Info("用户未注册", zap.String("phone", phone))
	svc.l.Info("用户未注册", logger.String("phone", phone))
	//loggerxx.Logger.Info("用户未注册", zap.String("phone", phone))
	// 在系统资源不足，触发降级之后，不执行慢路径了
	//if ctx.Value("降级") == "true" {
	//	return domain.User{}, errors.New("系统降级了")
	//}
	// 这个叫做慢路径
	// 你明确知道，没有这个用户
	u = domain.User{
		Phone: phone,
	}
	err = svc.repo.Create(ctx, u)
	if err != nil && err != repository.ErrUserDuplicate {
		return u, err
	}
	// 因为这里会遇到主从延迟的问题
	return svc.repo.FindByPhone(ctx, phone)
}

func (svc *userService) FindOrCreateByWechat(ctx context.Context,
	info domain.WechatInfo) (domain.User, error) {
	u, err := svc.repo.FindByWechat(ctx, info.OpenID)
	if err != repository.ErrUserNotFound {
		return u, err
	}
	u = domain.User{
		WechatInfo: info,
	}
	err = svc.repo.Create(ctx, u)
	if err != nil && err != repository.ErrUserDuplicate {
		return u, err
	}
	// 因为这里会遇到主从延迟的问题
	return svc.repo.FindByWechat(ctx, info.OpenID)
}

func (svc *userService) Profile(ctx context.Context,
	id int64) (domain.User, error) {
	u, err := svc.repo.FindById(ctx, id)
	return u, err
}

func PathsDownGrade(ctx context.Context, quick, slow func()) {
	quick()
	if ctx.Value("降级") == "true" {
		return
	}
	slow()
}
