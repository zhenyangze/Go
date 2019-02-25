# Go
Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。

#### ① [搭建开发环境](https://juejin.im/book/5b0778756fb9a07aa632301e/section/5b0d466bf265da08ee7edd20)

> Windows - src: %GOPATH%\src - 配置 set:

    GO_INSTALL_DIR=D:\Program
    GOROOT=D:\Program\Go\
    GOPATH=C:\Users\Administrator\go
    PATH=C:\Users\Administrator\go\bin

> Linux - src: $GOPATH/src - 配置 export:

    GO_INSTALL_DIR=/usr/local # 安装位置: 可更改为$HOME或其他
    GOROOT=$GO_INSTALL_DIR/go
    GOPATH=$HOME/gopath
    PATH=/usr/local/go/bin:$GOPATH/bin:$PATH

> 安装依赖包
~~~bash
# 先安装镜像包 (由于网络问题)
git clone https://github.com/golang/tools.git %GOPATH%/src/golang.org/x/tools   # Go Tools
git clone https://github.com/golang/mobile.git %GOPATH%/src/golang.org/x/mobile # Go on Mobile
git clone https://github.com/golang/build.git %GOPATH%/src/golang.org/x/build   # build and release
git clone https://github.com/golang/crypto.git %GOPATH%/src/golang.org/x/crypto # cryptography libraries
git clone https://github.com/golang/sys.git %GOPATH%/src/golang.org/x/sys       # low-level interaction
git clone https://github.com/golang/image.git %GOPATH%/src/golang.org/x/image   # image libraries
git clone https://github.com/golang/oauth2.git %GOPATH%/src/golang.org/x/oauth2 # OAuth 2.0
git clone https://github.com/golang/tour.git %GOPATH%/src/golang.org/x/tour
git clone https://github.com/golang/text.git %GOPATH%/src/golang.org/x/text
git clone https://github.com/golang/net.git %GOPATH%/src/golang.org/x/net
git clone https://github.com/golang/arch.git %GOPATH%/src/golang.org/x/arch
git clone https://github.com/golang/perf.git %GOPATH%/src/golang.org/x/perf
git clone https://github.com/golang/sync.git %GOPATH%/src/golang.org/x/sync
git clone https://github.com/golang/time.git %GOPATH%/src/golang.org/x/time
git clone https://github.com/golang/debug.git %GOPATH%/src/golang.org/x/debug
git clone https://github.com/golang/lint.git %GOPATH%/src/golang.org/x/lint/golint

# 开始实例
# go get github.com/golang/leveldb
go get github.com/golang/playground
go get github.com/golang/example/hello
~~~

#### ② [功能、框架、基础库、应用、工具等](https://github.com/avelino/awesome-go)

 * [QT跨平台应用框架](https://github.com/therecipe/qt)

#### ③ [构建企业级的 RESTful API 服务](https://juejin.im/book/5b0778756fb9a07aa632301e)

#### ④ [标准库文档](https://studygolang.com/pkgdoc)

#### ⑤ 阅读相关文章

 * 高性能
    * [高并发架构解决方案](https://studygolang.com/articles/15479)


----

