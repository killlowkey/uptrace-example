package store

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"uptrace-example/store/model"
)

type UserStore interface {
	Create(ctx context.Context, user *model.User) error
	FindById(ctx context.Context, id int64) (*model.User, error)
	List(ctx context.Context) ([]*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id int64) error
	BatchDelete(ctx context.Context, ids []int64) error
	BatchCreate(ctx context.Context, users []*model.User) error
	Exist(ctx context.Context, query interface{}, args ...interface{}) (bool, error)
}

type UserStoreImpl struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) UserStore {
	return &UserStoreImpl{db: db}
}

func (u *UserStoreImpl) Create(ctx context.Context, user *model.User) error {
	return u.db.WithContext(ctx).Create(user).Error
}

func (u *UserStoreImpl) FindById(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	err := u.db.WithContext(ctx).First(&user, id).Error
	return &user, err
}

func (u *UserStoreImpl) List(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	err := u.db.WithContext(ctx).Find(&users).Error
	return users, err
}

func (u *UserStoreImpl) Update(ctx context.Context, user *model.User) error {
	return u.db.WithContext(ctx).Save(user).Error
}

func (u *UserStoreImpl) Delete(ctx context.Context, id int64) error {
	return u.db.WithContext(ctx).Delete(&model.User{}, id).Error
}

func (u *UserStoreImpl) BatchDelete(ctx context.Context, ids []int64) error {
	return u.db.WithContext(ctx).Delete(&model.User{}, ids).Error
}

func (u *UserStoreImpl) BatchCreate(ctx context.Context, users []*model.User) error {
	return u.db.WithContext(ctx).Create(users).Error
}

func (u *UserStoreImpl) Exist(ctx context.Context, query interface{}, args ...interface{}) (bool, error) {
	var user model.User
	err := u.db.WithContext(ctx).Where(query, args...).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
