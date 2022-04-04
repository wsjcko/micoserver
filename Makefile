GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go install google.golang.org/protobuf/proto
	@go install github.com/gogo/protobuf/protoc-gen-gofast@latest
	@go install github.com/asim/go-micro/cmd/protoc-gen-micro/v4@latest

.PHONY: proto
proto:
	@protoc --proto_path=./proto/pb --micro_out=./protobuf/pb --gofast_out=./protobuf/pb proto/pb/*.proto

.PHONY: update
update:
	@go get -u go-micro.dev/v4@latest

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: dockerBuild
docker:
	@docker build -t micoserver:latest .
