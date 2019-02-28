# Go
Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。

#### ① [搭建开发环境](https://juejin.im/book/5b0778756fb9a07aa632301e/section/5b0d466bf265da08ee7edd20)
    安装版本> go version
    环境配置> go env

> Windows - src: %GOPATH%\src - 配置 set: cd %USERPROFILE% (C:\Users\Administrator)

    https://studygolang.com/dl/golang/go1.12.windows-amd64.msi
    GOROOT=D:\Program\Go\
    GOPATH=C:\Users\Administrator\go
    PATH=D:\Program\Go\bin;%GOPATH%\bin;%PATH%

> Linux - src: $GOPATH/src - 配置 export: cd $HOME (/root 或 /home)
    
    wget https://studygolang.com/dl/golang/go1.12.linux-amd64.tar.gz
    GO_INSTALL_DIR=/usr/local # 默认安装目录: 可更改 (选项 tar -C)
    tar -xvzf go1.12.linux-amd64.tar.gz -C $GO_INSTALL_DIR
    GOROOT=/usr/local/go
    GOPATH=/home/go
    PATH=/usr/local/go/bin:$GOPATH/bin:$PATH

> 安装依赖包 [集成工具](https://godoc.org/golang.org/x/tools)、 [项目依赖-包管理工具govendor](https://github.com/kardianos/govendor)
~~~bash
# 安装全局依赖-镜像包 (解决网络问题)
git clone https://github.com/golang/arch.git %GOPATH%/src/golang.org/x/arch     # 数据结构
git clone https://github.com/golang/build.git %GOPATH%/src/golang.org/x/build   # 构建、发布
git clone https://github.com/golang/crypto.git %GOPATH%/src/golang.org/x/crypto # 加密、安全
git clone https://github.com/golang/debug.git %GOPATH%/src/golang.org/x/debug   # 调试、跟踪
git clone https://github.com/golang/image.git %GOPATH%/src/golang.org/x/image   # 图片库
git clone https://github.com/golang/lint.git %GOPATH%/src/golang.org/x/lint/golint # 语法检查
git clone https://github.com/golang/mobile.git %GOPATH%/src/golang.org/x/mobile # 移动端
git clone https://github.com/golang/net.git %GOPATH%/src/golang.org/x/net       # 网络库
git clone https://github.com/golang/oauth2.git %GOPATH%/src/golang.org/x/oauth2 # OAuth 2.0 认证授权
git clone https://github.com/golang/perf.git %GOPATH%/src/golang.org/x/perf     # 性能测量、存储和分析
git clone https://github.com/golang/sync.git %GOPATH%/src/golang.org/x/sync     # 并发访问-同步锁
git clone https://github.com/golang/sys.git %GOPATH%/src/golang.org/x/sys       # 系统底层
git clone https://github.com/golang/text.git %GOPATH%/src/golang.org/x/text     # 文本处理
git clone https://github.com/golang/time.git %GOPATH%/src/golang.org/x/time     # 时间处理
git clone https://github.com/golang/tools.git %GOPATH%/src/golang.org/x/tools   # 工具包
git clone https://github.com/golang/tour.git %GOPATH%/src/golang.org/x/tour     # 其他

# 安装api* Gin依赖包 > cd %GOPATH%/src
git clone https://github.com/lexkong/vendor                                     # 项目依赖 govendor
git clone https://github.com/lexkong/apiserver_demos apiserver                  # 项目源码 复制^demo至 工作目录
go get github.com/StackExchange/wmi                                             # 项目依赖-缺失的包
# 编译api* apiserver例子> cd apiserver ; gofmt -w . ; go tool vet . ; go build -v .
# 运行 > apiserver.exe

# 开始playground*
go get github.com/golang/playground
go get github.com/golang/example/hello
go get github.com/golang/leveldb            # 内存数据库
~~~

#### ② [功能、框架、基础库、应用、工具等](https://github.com/avelino/awesome-go)

 * [QT跨平台应用框架](https://github.com/therecipe/qt)

#### ③ [构建企业级的 RESTful API 服务](https://juejin.im/book/5b0778756fb9a07aa632301e)

#### ④ [标准库文档](https://studygolang.com/pkgdoc)

#### ⑤ 阅读相关文章

 * 高性能
    * [高并发架构解决方案](https://studygolang.com/articles/15479)


----

