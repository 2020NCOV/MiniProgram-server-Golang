FROM golang as build

ADD . /usr/local/go/src/MiniProgram-server-Golang

WORKDIR /usr/local/go/src/MiniProgram-server-Golang

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOPROXY="https://goproxy.io" go build -o ncov_golang

FROM alpine:3.7

ENV MYSQL_DSN=""
ENV GIN_MODE="release"
ENV PORT=8080

RUN echo "http://mirrors.aliyun.com/alpine/v3.7/main/" > /etc/apk/repositories && \
    apk update && \
    apk add ca-certificates && \
    echo "hosts:files dns" > /etc/nsswitch.conf && \
    mkdir -p /www/conf

WORKDIR /www

COPY --from=build /usr/local/go/src/MiniProgram-server-Golang/ncov_golang /usr/bin/ncov_golang
ADD ./conf /www/conf

RUN chmod +x /usr/bin/ncov_golang

ENTRYPOINT ["ncov_golang"]