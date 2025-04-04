# 使用Go官方镜像作为构建环境
FROM golang:1.20-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制go.mod和go.sum文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server

# 使用轻量级alpine镜像作为运行环境
FROM alpine:latest

# 安装必要的CA证书
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 从构建阶段复制编译好的应用
COPY --from=builder /app/main .

# 暴露应用端口
EXPOSE 8080

# 运行应用
CMD ["./main"]