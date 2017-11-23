# hello-go

## For Ubuntu16.04 server

1.install by apt-get
```
sudo apt-get install goloang-go
```
2.config go env
```
sudo vim vim /etc/profile
export GOROOT=/usr/lib/go-1.6
export GOPATH=$HOME/go-path
```
3.go cmd
```
$ go run xxx.go
$ go test
$ go test -v
$ go test -v -test.run 

$ sudo apt-get install golang-golang-x-tools
$ godoc -http=:8083
$ nohup godoc -http=:8083 &
```
```
$ go version
go version go1.6.2 linux/amd64
```
```
$ go env
GOARCH="amd64"
GOBIN=""
GOEXE=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOOS="linux"
GOPATH="/home/jihao/go-path"
GORACE=""
GOROOT="/usr/lib/go-1.6"
GOTOOLDIR="/usr/lib/go-1.6/pkg/tool/linux_amd64"
GO15VENDOREXPERIMENT="1"
CC="gcc"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0"
CXX="g++"
CGO_ENABLED="1
```

## Golang Format
```
import (
    "encoding/json"
    "strings"

    "myproject/models"
    "myproject/controller"
    "myproject/utils"

    "github.com/astaxie/beego"
    "github.com/go-sql-driver/mysql"
)   
```

## References
- http://www.yiibai.com/go/go_start.html
- http://www.runoob.com/go/go-tutorial.html
- http://studygolang.com/articles/2059
- https://godoc.org/
- https://play.golang.org/
