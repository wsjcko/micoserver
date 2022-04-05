package service

import (
	"errors"
	"github.com/wsjcko/micoserver/domain/model"
	"github.com/wsjcko/micoserver/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	AddUser(user *model.User) (int64, error)
	DeleteUser(int64) error
	UpdateUser(user *model.User, isChangePwd bool) error
	FindUserByName(string) (*model.User, error)
	CheckPwd(userName string, pwd string) (isOk bool, err error)
}

type UserService struct {
	UserRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

//加密用户密码 bcrypt算法对于同一个密码，每次生成的hash不一样
func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

//验证用户密码
func ValidatePassword(userPassword string, hashed string) (isOk bool, err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword))
	if err != nil {
		return false, errors.New("密码错误")
	}
	return true, nil
}

func (u *UserService) AddUser(user *model.User) (int64, error) {
	pwdByte, err := GeneratePassword(user.HashPassword)
	if err != nil {
		return user.Id, err
	}
	user.HashPassword = string(pwdByte)
	return u.UserRepository.CreateUser(user)
}

func (u *UserService) DeleteUser(id int64) error {
	return u.UserRepository.DeleteUserById(id)
}

func (u *UserService) UpdateUser(user *model.User, isChangePwd bool) error {
	if isChangePwd {
		pwdByte, err := GeneratePassword(user.HashPassword)
		if err != nil {
			return err
		}
		user.HashPassword = string(pwdByte)
	}
	return u.UserRepository.UpdateUser(user)
}

func (u *UserService) FindUserByName(userName string) (*model.User, error) {
	return u.UserRepository.FindUserByName(userName)
}

func (u *UserService) CheckPwd(userName string, pwd string) (isOk bool, err error) {
	user, err := u.UserRepository.FindUserByName(userName)
	if err != nil {
		return false, err
	}
	return ValidatePassword(pwd, user.HashPassword)
}
