### micro
```
### 设置代理
go env -w GO111MODULE=on

go env -w GOPROXY=https://goproxy.cn,https://proxy.golang.org,direct


### 安装micro
go install go-micro.dev/v4/cmd/micro@master
micro -v
### 安装protoc-gen-micro工具, protoc --micro_out
go install go-micro.dev/v4/cmd/protoc-gen-micro@master

go install github.com/golang/protobuf/protoc-gen-go@master
go install github.com/gogo/protobuf/protoc-gen-gofast
go install github.com/golang/protobuf/proto@master 
go get -u -v google.golang.org/grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@master

### 指定相对于$GOPATH的目录路径，快速创建一个新服务。
micro new service $GOPATH/michserver
code michserver
go mod init
go mod tidy


### 根据自己的需求去修改proto文件
touch cmicro.proto
没有指定 option go_package = "./proto;michserver";
protoc -I=./proto --micro_out=./proto --gofast_out=./proto  cmicro.proto

有go_package 用如下
protoc -I=./proto --go_out=plugins=grpc:. mich.proto  //生成grpc mich.pb.go

protoc -I=./proto --micro_out=. --go_out=.  mich.proto //生成go mich.pb.go、mich.pb.micro.go

protoc -I=./proto --micro_out=. --gofast_out=. mich.proto  //生成gofast mich.pb.go、mich.pb.micro.go

protoc -I=./proto --micro_out=. --go_out=.  --go-grpc_out=. mich.proto 



### 客户端，服务端
搜索Server API for  |  Client API for 分别实现对应的方法
```





### docker 
```docker


创建或修改 /etc/docker/daemon.json 文件,加载配置重启docker
{
    "registry-mirrors": [
    "https://reg-mirror.qiniu.com",
    "https://registry.docker-cn.com",
    "https://v0z45uix.mirror.aliyuncs.com",
    "http://hub-mirror.c.163.com",
    "https://docker.mirrors.ustc.edu.cn"
  ]
}
systemctl daemon-reload
systemctl restart docker
docker info

[docker文档](https://docs.docker.com/engine/reference/commandline/search)

docker search --help
docker search --filter is-official=true --filter stars=3 nginx

docker ps -a 查看所有的容器
docker ps 查看所有运行的容器
CONTAINER ID   IMAGE
a28ea788bfc8   mysql:5.7 

docker images 查看本地images
docker images nginx 查看本地images nginx

docker rmi 9ad401599ff2 删除本地images

docker pull nginx:last
docker stop b8697e34ff96 停止运行容器
docker rm b8697e34ff96 #删除不运行的容器，正在运行，需要停止


docker tag mysql:5.7 wsjcko/mysql57
docker push wsjcko/mysql57:game
docker commit a28ea788bfc8 wsjcko/mysql57:game


修改名称mysql57 => mysql5.7
docker tag wsjcko/mysql57:game wsjcko/mysql5.7:game
docker login
docker push wsjcko/mysql5.7:game


Dcokerfile 基本结构


Dockerfile的四部分
基础镜像信息
维护者信息
镜像操作指令
容器启动指令


基础镜像信息
FROM ubuntu
FROM alpine
FROM nginx

维护者信息
MAINTAINER wsjcko wsjcko@163.com

docker run -d -p 8443:443 -p 8480:80 -p 8422:22 --name gitlab --restart always -v /mnt/f/Docker/images/gitlab/config:/etc/gitlab -v /mnt/f/Docker/images/gitlab/logs:/var/log/gitlab -v /mnt/f/Docker/images/gitlab/data:/data/gitlab gitlab/gitlab-ce:latest

docker run -p 3306:3306 --name mysql57   -e MYSQL_ROOT_PASSWORD=123456 -d mysql:5.7 --character-set-server=utf8mb4 --collation-server=utf8mb4_general_ci

docker cp ./hpsm_maindb.sql mysql57:/home

docker run -p 6379:6379 --name gRedis  -v /mnt/f/Docker/images/redis/redis.conf:/etc/redis/redis.conf --privileged=true -d redis redis-server /etc/redis/redis.conf

docker logs -f  -t --tail 100 gRedis
docker exec -it  a28ea788bfc8 /bin/bash


RUN 镜像操作指令, 指定镜像被构建时要运行的命令
RUN echo "deb http://archive.ubuntu.com/ubuntu/ raring main universe" >> /et
RUN apt-get update && apt-get install -y nginx
RUN echo "\ndaemon off;" >> /etc/nginx/nginx.conf

CMD 指定容器被启动时要运行的命令,CMD 指令可以被 docker run 命令覆盖
CMD  Dockerfile中只能指定一条CMD命令。 如果有多个CMD，只有最后一个生效，在这一点上，CMD和ENTRYPOINT一样
CMD /usr/sbin/nginx
CMD ["nginx", "-g", "daemon off;"]
CMD ["/bin/bash"] 
docker run -it nginx:latest /bin/uname -a


ENTRYPOINT ["/usr/sbin/nginx"]
CMD ["-h"]
如果在启动容器时没有指定任何参数
则CMD 指令-h 参数会传递给nginx 守护进程，即执行/usr/sbin/nginx -h


ENTRYPOINT 指令可以被 docker run 命令传参
两种格式:

ENTRYPOINT ["executable", "param1", "param2"]
ENTRYPOINT command param1 param2
配置容器启动后执行的命令, 并且不会被docker run提供的参数覆盖. 如果有多个ENTRPOINT, 只有最后一个生效.

ENTRYPOINT ["/usr/sbin/nginx"]
docker run -it nginx:latest -g "daemon off;"


其他指令
EXPOSE
EXPOSE <port> [<port>...]

容器暴露的端口, 启动时要指定 -P 自动给容器分配端口, 或者 -p aaa:bbb 手动分配端口

ENV
ENV PG_MAJOR 9.3
ENV PG_VERSION 9.3.4
ENV PATH /usr/local/postgres-$PG_MAJOR/bin:$PATH
指定一个环境变量,会被后续RUN指令使用,并在容器运行时保持。

ADD
ADD <src> <dest>
复制源目录到容器中的目录, 源目录是dockerfile的相对路径, 也可以是URL或者一个tar文件.

COPY
COPY <src> <dest>
跟ADD差不多, 不过只能复制相对路径, 推荐使用COPY.



VOLUME
VOLUME ["/date"]
挂载数据卷.

USER
USER root
指定容器启动时默认的用户名.

WORKDIR
WORKDIR /path/to/workdir
为后续指令提供工作目录(默认目录).
可以理解为cd到这个目录中, 然后后面的指令都是基于这个目录的.

可以使用多个WORKDIR, 比如:
WORKDIR /a
WORKDIR b
WORKDIR c
那么当前的目录则为/a/b/c

ONBUILD
ONBUILD [INSTRUCTION]
如果一个镜像的dockerfile中含有这个指令, 则基于这个镜像创建新的镜像的时候,都会执行指令后的内容.
```

### docker example micoserver
cd microserver
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 

docker build -t micoserver:latest .
或
docker build -t micoserver:latest -f Dockerfile 指定路径

docker run -t micoserver:latest 直接可以看log
或
docker run -d micoserver:latest
docker ps -a|grep micoserver
docker logs 0aeb7916e0d3





### go module
```
#####  设置代理
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,https://proxy.golang.org,direct

#####  go 命令会从公共镜像 http://goproxy.cn 上下载依赖包
#####  私有的代码仓库,跳过 proxy server 和校验检查
go env -w GOPRIVATE=gitcode.wonderful.com,gitlab.sw.com

#####  module校验数据库
go env -w GOSUMDB=sum.golang.org 默认
go env -w GOSUMDB=*.corp.com,rsc.io/private
#####  私有仓库和模块，不用校验
go env -w GONOSUMDB=*.corp.com,rsc.io/private 

go mod init 初始化当前文件夹, 创建go.mod文件
go mod edit 编辑go.mod文件
go mod graph 打印模块依赖图
go mod tidy 增加缺少的module，删除无用的module
go mod vendor 将依赖复制到vendor下
go mod verify 校验依赖
go mod why 解释为什么需要依赖
go mod download 下载依赖的module到本地cache（默认为$GOPATH/pkg/mod目录）

手动修改go.mod
go mod edit -fmt 格式化
或

go mod 命令修改 go.mod

#####  go version
go mod edit -go=1.17

#####  modulename
go mod edit --module= gameserver

#####  添加或删除 依赖
go mod edit -require=golang.org/x/text
go mod edit --droprequire=golang.org/x/text

##### 指定依赖版本，但仅限小版本号。
go mod edit -require=github.com/dgraph-io/badger/v3@v3.2011.1

# gojenkins二次开发后，替换github 或者gitee路径
go mod edit -replace="github.com/bndr/gojenkins => github.com/wsjcko/gojenkins@v1.10"

# 特殊版本
go get github.com/osrg/gobgp@915bfc2 //版本号的hash获取该版本
go mod edit -replace="github.com/tealeg/xlsx => github.com/tealeg/xlsx v1.0.4-0.20190807182118-a6243d92b369"
go mod edit -replace="google.golang.org/grpc => google.golang.org/grpc v1.26.0"

# 本地替代
go mod edit -replace="gitlab.hd.com/prometheus/sdk_go => /home/sun/gopath/src/gitlab.hd.com/prometheus/sdk_go"

# 修改版本号
# panic: proto: file "common.proto" is already registered
go mod edit -replace=github.com/gogo/protobuf=github.com/gogo/protobuf@v1.3.1
go mod edit -replace=github.com/golang/protobuf=github.com/golang/protobuf@v1.3.5

# 编译
go build -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn"

go mod tidy
特殊需要：  go mod tidy -go=1.16 && go mod tidy -go=1.17

# 私有库
go env -w GOPRIVATE=github.com/wsjcko
git config --global url."https://wsjcko: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@github.com".insteadOf "https://github.com"

git config --global http.extraheader "PRIVATE-TOKEN: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"
git config --global url."git@github.com:wsjcko".insteadOf "https://github.com/wsjcko"

git config --global url."git@github.com:wsjcko/user.git".insteadOf "https://github.com/wsjcko/user.git"
go mod edit -replace="github.com/wsjcko/user =github.com/wsjcko/user@v1.0.0"

git config --global url."git@github.com:wsjcko/micoserver.git".insteadOf "https://github.com/wsjcko/micoserver.git"
go mod edit -replace="github.com/wsjcko/micoserver =github.com/wsjcko/micoserver@v1.0.0"
go env -w GOPRIVATE=github.com

 cat ~/.gitconfig
 go env
```