package biz

import (
	"context"
	"encoding/json"
	"errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"uptrace-example/store"
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
	if exist, err := u.store.Exist(ctx, map[string]string{"name": user.Name}); err != nil {
		return err
	} else if exist {
		return errors.New("user already exists")
	}

	return u.store.Create(ctx, user)
}

// FindByID 通过 id 查找用户
// 手动创建 span，来追踪自身的业务逻辑
func (u *UserServiceImpl) FindByID(ctx context.Context, id int64) (*store.User, error) {
	c, span := u.tracer.Start(ctx, "UserService#FindByID")
	defer span.End()

	span.SetAttributes(attribute.Int64("method.args.id", id))

	user, err := u.store.FindById(c, id)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.SetAttributes(attribute.String("method.resp.error", err.Error()))
		return nil, err
	} else {
		span.SetStatus(codes.Ok, "ok")
		data, _ := json.Marshal(user)
		span.SetAttributes(attribute.String("method.resp.user", string(data)))
		return user, err
	}
}
