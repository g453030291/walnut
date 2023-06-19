# walnut

安装后的环境配置：
go env
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GO111MODULE=on

GOMODCACHE 环境变量，自定义本地 module 的缓存路径。

编译执行：
go build .\hello.go
.\hello.exe

直接执行：
go run .\hello.go

go moudle:
go mod init hellomodule
go mod tidy
go build .\hellomodule.go
go clean -i -n 			清除打包的可执行文件

go get github.com/google/uuid	go module添加一个依赖
go mod tidy			go module自动分析源码中的依赖，并下载
go list -m -versions github.com/sirupsen/logrus	查询某个包的版本
go get github.com/sirupsen/logrus@v1.7.0	将当前项目中某个包，更改为手动指定的版本

使用go mod tidy来自动依赖指定的包版本，先edit指定某个包的版本，再执行go mod tidy
go mod edit -require=github.com/sirupsen/logrus@v1.7.0
go mod tidy

go list -m all	列出当前项目所有的依赖

vendor-->go module依赖保留了vendor机制，可以以vendor依赖为准，离线构建项目
go mod vendor	建立vendor目录（go mod vendor 命令在 vendor 目录下，创建了一份这个项目的依赖包的副本，并且通过 vendor/modules.txt 记录了 vendor 下的 module 以及版本）
go build -mod=vendor	不用go module的缓存构建项目，而是基于项目中的vendor构建项目（go1.14以后的版本，go build都会优先根据vendor构建项目，除非传入go build -mod=mod）

walnut记录:

启动项目:

`go run main.go`

服务器启动:

`nohup ./walunt-0.1.sh >walunt.log 2>&1 &`

打包：

`go build -o bin/walunt-0.1.sh main.go`

`go get -u github.com/gin-gonic/gin@v1.9.0`

异构编译:

`GOOS=linux GOARCH=amd64 go build -o bin/walunt-0.1.sh main.go`

nginx加载配置重启:

`sudo ./nginx -s reload`
