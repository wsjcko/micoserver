package handler

import (
	"context"
	"github.com/wsjcko/micoserver/domain/model"
	"github.com/wsjcko/micoserver/domain/service"
	pb "github.com/wsjcko/micoserver/protobuf/pb"
)

type UserServer struct {
	UserService service.IUserService
}

func (us *UserServer) Init(userService service.IUserService) {
	us.UserService = userService
}

//注册
func (us *UserServer) Register(ctx context.Context, req *pb.UserRegisterReq, ack *pb.UserRegisterAck) error {
	user := &model.User{
		UserName:     req.UserName,
		FirstName:    req.FirstName,
		HashPassword: req.Pwd,
	}
	_, err := us.UserService.AddUser(user)
	ack.Message = "注册成功"
	return err
}

//登录
func (us *UserServer) Login(ctx context.Context, req *pb.UserLoginReq, ack *pb.UserLoginAck) error {
	isOk, err := us.UserService.CheckPwd(req.UserName, req.Pwd)
	if err != nil {
		return err
	}
	ack.IsSuccess = isOk
	return nil
}

//获取信息
func (us *UserServer) GetUserInfo(ctx context.Context, req *pb.UserInfoReq, ack *pb.UserInfoAck) error {
	userInfo, err := us.UserService.FindUserByName(req.UserName)
	if err != nil {
		return err
	}

	ack.Id = userInfo.Id
	ack.UserName = userInfo.UserName
	ack.FirstName = userInfo.FirstName
	return nil
}
