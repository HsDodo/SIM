### 部署项目

#### 1. 构建基础镜像

> 首先将各微服务 build 出来，也就是要将编译和运行分离开
>
> 根据下面构建出的镜像 ，sim_server 中有编译好的各服务的二进制运行文件，
>
> 再将这些运行文件分到其他镜像中，这就能使各镜像尽可能的小

```dockerfile
FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
COPY . .

RUN go mod tidy

RUN go build -o auth/api/auth     auth/api/auth.go

RUN go build -o chat/api/chat     chat/api/chat.go
RUN go build -o chat/rpc/chatrpc  chat/rpc/chatrpc.go

RUN go build -o file/api/file     file/api/file.go
RUN go build -o file/rpc/filerpc  file/rpc/filerpc.go

RUN go build -o gateway/gateway    gateway/gateway.go

RUN go build -o group/api/group    group/api/group.go

RUN go build -o logs/api/logs     logs/api/logs.go

RUN go build -o user/api/users     user/api/user.go
RUN go build -o user/rpc/userrpc   user/rpc/userrpc.go


```

#### 2. 构建服务的具体镜像

如 Auth 服务

```dockerfile
FROM sim_server AS builder

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /build/auth/api/auth .

CMD ["./auth"]

```

从基础镜像中拿编译好的服务的二进制运行文件，其他服务也一样的写法

#### 3. 用 docker-compose 来编排各服务容器