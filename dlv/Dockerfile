# Compile stage
FROM golang:1.22.3-alpine3.20 AS build-env

# Build Delve
RUN go env -w  GOPROXY=https://goproxy.cn,direct \
    && go get github.com/go-delve/delve/cmd/dlv

ADD ../ /dockerdev
WORKDIR /dockerdev

# 编译需要 debug 的程序
RUN go build -gcflags="all=-N -l" -o /server

# Final stage
FROM debian:buster

# 分别暴露 server 和 dlv 端口
EXPOSE 8000 40000

WORKDIR /
COPY --from=build-env /go/bin/dlv /
COPY --from=build-env /server /

CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/server"]
