package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/wsjcko/micoserver/domain/repository"
	"github.com/wsjcko/micoserver/domain/service"
	"github.com/wsjcko/micoserver/handler"
	pb "github.com/wsjcko/micoserver/protobuf/pb"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	serviceName = "micoserver"
	version     = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(serviceName),
		micro.Version(version),
		// micro.Address(addr string)
		// micro.Transport(t transport.Transport)
		// micro.Registry(r registry.Registry)
		// micro.RegisterHandler(s server.Server, h interface{}, opts ...server.HandlerOption) //pb.RegisterMicoserverHandler 实现了
	)
	srv.Init()
	log.Info("Create service")

	//数据库初始化
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/microUser?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error("Connect Mysql ", err)
	}
	log.Info("Connect Mysql")

	defer db.Close()               //释放数据库连接资源
	db.SingularTable(true)         //gorm创建表时，false:表添加s后缀
	db.LogMode(true)               //开启sql log
	db.DB().SetMaxIdleConns(10)    //最大空闲连接
	db.DB().SetMaxOpenConns(25)    //最大连接数
	db.DB().SetConnMaxLifetime(30) //最大生存时间(s)

	//创建表
	rp := repository.NewUserRepository(db)
	// rp.InitTable() //gorm 创建表 user 只需执行一次
	log.Info("InitTable")

	// Register handler
	err = pb.RegisterMicoserverHandler(srv.Server(), new(handler.Micoserver))
	if err != nil {
		log.Fatal("Register handler Call ", err)
		return
	}

	userServer := new(handler.UserServer)
	userServer.Init(service.NewUserService(rp))
	err = pb.RegisterUserHandler(srv.Server(), userServer)
	if err != nil {
		log.Fatal("Register handler ", err)
		return
	}
	log.Info("Register handler")

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal("Run service ", err)
	}
	log.Info("Run service")
}
