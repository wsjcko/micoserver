package main

import (
	"context"
	"strconv"
	"time"

	pb "github.com/wsjcko/micoserver/protobuf/pb"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

func main() {
	//实例化
	clientService := micro.NewService(
		micro.Name("micoserver.client"),
	)
	//初始化
	clientService.Init()
	ctx := context.TODO()

	if err := clientService.Run(); err != nil {
		log.Fatal("Run service ", err)
	}
	log.Info("Run service")

	//测试通信
	micoService := pb.NewMicoserverService("micoserver", clientService.Client())
	res, err := micoService.Call(ctx, &pb.CallRequest{Name: "wuxue"})
	if err != nil {
		log.Fatal("micoserver call ", err)
		return
	}
	log.Info(res.GetMsg())

	//注册
	userService := pb.NewUserService("micoserver", clientService.Client())
	userName := "jingcheng" + strconv.Itoa(int(time.Now().Unix()))
	registerReq := &pb.UserRegisterReq{
		UserName:  userName,
		FirstName: "sun",
		Pwd:       "sjc123456",
	}
	registerRes, err := userService.Register(ctx, registerReq)
	if err != nil {
		log.Fatal("micoserver Register ", err)
		return
	}
	log.Info(registerRes.GetMessage())

	//登录
	loginReq := &pb.UserLoginReq{
		UserName: userName,
		Pwd:      "sjc123456",
	}
	loginRes, err := userService.Login(ctx, loginReq)
	if err != nil {
		log.Fatal("micoserver Login ", err)
		return
	}
	log.Info(loginRes.GetIsSuccess())

	//获取信息
	userInfoReq := &pb.UserInfoReq{
		UserName: userName,
	}
	userInfoRes, err := userService.GetUserInfo(ctx, userInfoReq)
	if err != nil {
		log.Fatal("micoserver GetUserInfo ", err)
		return
	}
	log.Info(userInfoRes.GetFirstName(), userInfoRes.GetUserName(), userInfoRes.GetId())

}
