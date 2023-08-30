FROM golang:latest AS build

COPY common /go/src/common/
COPY config /go/src/config/
COPY controller /go/src/controller/
COPY log /go/src/log/
COPY middleware /go/src/middleware/
COPY model /go/src/model/
COPY repo /go/src/repo/
COPY service /go/src/service/
COPY test /go/src/test/
COPY utils /go/src/utils/
COPY minioHandler /go/src/minioHandler/

COPY go.mod go.sum *.go /go/src/

WORKDIR "/go/src/"
RUN go env -w GO111MODULE=on \
  && go env -w GOPROXY=https://goproxy.cn,direct \
  && go env -w GOOS=linux \
  && go env -w GOARCH=amd64
RUN go mod tidy
RUN go build -o douyin_project

FROM jrottenberg/ffmpeg

RUN mkdir "/app"
COPY --from=build /go/src/douyin_project /app/douyin_project

RUN chmod +x /app/douyin_project

EXPOSE 18005
ENTRYPOINT ["/app/douyin_project"]
