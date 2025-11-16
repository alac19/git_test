# 构建阶段
FROM golang:1.21-alpine AS builder
ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /app
COPY hello-http.go .
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o main hello-http.go

# 运行阶段
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]