# ローカル開発環境
FROM golang:1.19.1-bullseye as debugger
RUN go install github.com/rakyll/gotest@latest \
  && go install github.com/jwilder/dockerize@latest
WORKDIR /go/src/github.com/gold-kou/cfn-practice
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 80
CMD ["dockerize", "-wait", "tcp://cfn-practice-db:3306", "-timeout", "60s", "go", "run", "main.go"]

# 本番で使うバイナリのビルド用環境
FROM golang:1.19.1-bullseye as builder
WORKDIR /go/src/github.com/gold-kou/cfn-practice/backend
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build

# 本番の実行環境
FROM debian:11.5-slim as executor
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /
COPY --from=builder /go/src/github.com/gold-kou/cfn-practice/backend/backend /backend
EXPOSE 80
CMD ["/backend"]