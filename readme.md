# go make 版本自动++，xx version显示版本等信息

```
-MB1:versionpp yezi$ make
./versionpp: ./versionpp: cannot execute binary file
go build -ldflags "-X 'main.GoVersion=go version go1.14.2 darwin/amd64' -X main.BuildTime=2020-12-14-CST/23:50:48 -X main.ProgramCommitID= -X main.Version=" -v
_/Users/yezi/workspace/versionpp
-MB1:versionpp yezi$ make
go build -ldflags "-X 'main.GoVersion=go version go1.14.2 darwin/amd64' -X main.BuildTime=2020-12-14-CST/23:50:53 -X main.ProgramCommitID= -X main.Version=0.0.1" -v
-MB1:versionpp yezi$ make
go build -ldflags "-X 'main.GoVersion=go version go1.14.2 darwin/amd64' -X main.BuildTime=2020-12-14-CST/23:50:55 -X main.ProgramCommitID= -X main.Version=0.0.2" -v
-MB1:versionpp yezi$ make
go build -ldflags "-X 'main.GoVersion=go version go1.14.2 darwin/amd64' -X main.BuildTime=2020-12-14-CST/23:50:56 -X main.ProgramCommitID= -X main.Version=0.0.3" -v
-MB1:versionpp yezi$ make
go build -ldflags "-X 'main.GoVersion=go version go1.14.2 darwin/amd64' -X main.BuildTime=2020-12-14-CST/23:50:57 -X main.ProgramCommitID= -X main.Version=0.0.4" -v
-MB1:versionpp yezi$ make
go build -ldflags "-X 'main.GoVersion=go version go1.14.2 darwin/amd64' -X main.BuildTime=2020-12-14-CST/23:50:57 -X main.ProgramCommitID= -X main.Version=0.0.5" -v
-MB1:versionpp yezi$ make
go build -ldflags "-X 'main.GoVersion=go version go1.14.2 darwin/amd64' -X main.BuildTime=2020-12-14-CST/23:50:58 -X main.ProgramCommitID= -X main.Version=0.0.6" -v
-MB1:versionpp yezi$ make
go build -ldflags "-X 'main.GoVersion=go version go1.14.2 darwin/amd64' -X main.BuildTime=2020-12-14-CST/23:50:59 -X main.ProgramCommitID= -X main.Version=0.0.7" -v
-MB1:versionpp yezi$ 
```

# 显示版本信息
```
-MB1:versionpp yezi$ ./versionpp version
Go version: 	go version go1.14.2 darwin/amd64
Build Time: 	2020-12-14-CST/23:50:59
Commit ID : 	

Version : 	./versionpp_0.0.7
```
