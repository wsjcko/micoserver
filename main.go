package main

import (
	"micoserver/handler"
	pb "micoserver/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "micoserver"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()
	logger.Info("Create service")

	//数据库初始化
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/microUser?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		logger.Fatal(err)
		return
	}
	logger.Info("Connect Mysql")

	defer db.Close()               //释放数据库连接资源
	db.SingularTable(true)         //gorm创建表时，false:表添加s后缀
	db.LogMode(true)               //开启sql log
	db.DB().SetMaxIdleConns(10)    //最大空闲连接
	db.DB().SetMaxOpenConns(25)    //最大连接数
	db.DB().SetConnMaxLifetime(30) //最大生存时间(s)

	//创建表
	rp := repository.NewUserRepository(db)
	rp.InitTable() //只执行一次
	logger.Info("InitTable")

	// Register handler
	userService := mservice.NewUserService(rp)
	pb.RegisterUserHandler(srv.Server(), new(handler.UserServer{UserSevice: userService}))
	if err != nil {
		logger.Fatal(err)
		return
	}
	logger.Info("Register handler")

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
	logger.Info("Run service")
}
