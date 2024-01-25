package biz

import (
	"context"
	"errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"uptrace-example/store"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)

type UserService interface {
	Create(ctx context.Context, user *store.User) error
	FindByID(ctx context.Context, id int64) (*store.User, error)
}

type UserServiceImpl struct {
	store  store.UserStore
	tracer trace.Tracer
}

func NewUserService(store store.UserStore) UserService {
	return &UserServiceImpl{
		store:  store,
		tracer: otel.GetTracerProvider().Tracer("UserService"),
	}
}

func (u *UserServiceImpl) Create(ctx context.Context, user *store.User) error {
	c, span := u.tracer.Start(ctx, "Create")
	defer span.End()

	if exist, err := u.store.Exist(c, map[string]string{"name": user.Name}); err != nil {
		span.SetAttributes(attribute.String("method.args.user", user.ToJson()))
		span.RecordError(err, trace.WithStackTrace(true))
		return err
	} else if exist {
		span.SetAttributes(attribute.String("method.args.user", user.ToJson()))
		span.RecordError(ErrUserAlreadyExists, trace.WithStackTrace(true))
		return ErrUserAlreadyExists
	}

	return u.store.Create(c, user)
}

// FindByID 通过 id 查找用户
// 手动创建 span，来追踪自身的业务逻辑
// 成功的处理，不需要记录方法参数和方法返回值
func (u *UserServiceImpl) FindByID(ctx context.Context, id int64) (*store.User, error) {
	c, span := u.tracer.Start(ctx, "FindByID")
	defer span.End()

	user, err := u.store.FindById(c, id)
	if err != nil {
		// https://opentelemetry.io/docs/specs/semconv/exceptions/exceptions-spans/
		// 记录错误信息、堆栈信息
		span.SetAttributes(attribute.Int64("method.args.id", id))
		span.RecordError(err, trace.WithStackTrace(true))
		return nil, err
	} else {
		span.SetStatus(codes.Ok, "success")
		//span.SetAttributes(attribute.String("method.response", user.ToJson()))
		return user, err
	}
}
