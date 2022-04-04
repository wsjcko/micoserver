package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/wsjcko/micoserver/domain/model"
)

type IUserRepository interface {
	//初始化数据表
	InitTable() error
	//创建用户
	CreateUser(*model.User) (int64, error)
	//根据用户Id查找用户信息
	FindUserById(int64) (*model.User, error)
	//根据用户名称查找用户信息
	FindUserByName(string) (*model.User, error)
	//更新用户信息
	UpdateUser(*model.User) error
	//根据用户id删除用户
	DeleteUserById(int64) error
	//查找所有用户
	FindAll() ([]model.User, error)
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{sqlDb: db}
}

type UserRepository struct {
	sqlDb *gorm.DB
}

func (u *UserRepository) InitTable() error {
	return u.sqlDb.CreateTable(&model.User{}).Error
}

func (u *UserRepository) CreateUser(user *model.User) (int64, error) {
	err := u.sqlDb.Create(user).Error
	return user.Id, err
}

func (u *UserRepository) FindUserById(id int64) (user *model.User, err error) {
	err = u.sqlDb.First(user, id).Error
	return user, err
}

func (u *UserRepository) FindUserByName(name string) (user *model.User, err error) {
	err = u.sqlDb.Where("name = ?", name).Find(user).Error
	return user, err
}

func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.sqlDb.Model(user).Update(&user).Error
}
func (u *UserRepository) DeleteUserById(id int64) error {
	return u.sqlDb.Where("id = ?", id).Delete(&model.User{}).Error
}
func (u *UserRepository) FindAll() (userAll []model.User, err error) {
	return userAll, u.sqlDb.Find(&userAll).Error
}
