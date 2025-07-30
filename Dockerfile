# 1. 使用官方 Golang 镜像作为构建阶段
FROM golang:1.21 AS builder

WORKDIR /app

# 将 go.mod 和 go.sum 拷贝进容器
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 拷贝剩余的源代码
COPY . .

# 编译为静态文件
RUN CGO_ENABLED=0 GOOS=linux go build -o hello

# 2. 使用一个更小的镜像作为运行阶段
FROM alpine:latest

WORKDIR /root/

# 从构建阶段复制编译后的可执行文件
COPY --from=builder /app/hello .

# 暴露端口
EXPOSE 8080

# 设置启动命令
CMD ["./hello"]