# 使用Go的官方映像作为基础映像。这个映像是Go 1.9预安装的。该映像的$GOPATH值已被设置为/go。所有安装在/go/src的程序包都能通过go命令访问。
FROM golang:1.9

# 安装beego程序包和bee工具。beego程序包将在应用程序内部使用，bee工具将用于在开发过程中实时重载代码。
RUN go get github.com/golang/dep/cmd/dep

RUN dep ensure

# 通过开发计算机上容器的8080端口暴露该应用程序。最后一行，
EXPOSE 8080

# 使用bee命令开始对我们的应用程序进行实时重载。
CMD ["go", "run", 'main.go']