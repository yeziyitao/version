NAME=versionpp
VERSION=$(shell ./${NAME} newversion)
GO_VERSION=$(shell go version)
BUILD_TIME=$(shell date +%F-%Z/%T)
#ProgramCommitID=$(shell git rev-parse HEAD) # 项目是 Git 版本控制，则获取当前 commit id

LDFLAGS=-ldflags "-X 'main.GoVersion=${GO_VERSION}' -X main.BuildTime=${BUILD_TIME} -X main.ProgramCommitID=${ProgramCommitID} -X main.Version=${VERSION}"

all:
	go build ${LDFLAGS} -v
