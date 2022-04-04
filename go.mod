module github.com/wsjcko/micoserver

go 1.17

require (
	github.com/golang/protobuf v1.5.2
	github.com/jinzhu/gorm v1.9.16
	go-micro.dev/v4 v4.6.0
	golang.org/x/crypto v0.0.0-20220331220935-ae2d96664a29
	google.golang.org/protobuf v1.28.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/wsjcko/micoserver => github.com/wsjcko/micoserver v1.0.0
