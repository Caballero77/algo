FROM golang:1.13.5-alpine3.10

COPY . .
RUN go env -w GO111MODULE=on
RUN go get github.com/kataras/iris/v12@latest

EXPOSE 80
ENTRYPOINT go run api.go