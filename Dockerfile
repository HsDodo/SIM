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

