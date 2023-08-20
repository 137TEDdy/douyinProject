FROM golang:latest AS build

COPY biz /go/src/biz/
COPY pkg /go/src/pkg/
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
COPY --from=build /go/src/douyinProject /app/douyin_project

RUN chmod +x /app/douyin_project

EXPOSE 18005
ENTRYPOINT ["/app/douyin_project"]
