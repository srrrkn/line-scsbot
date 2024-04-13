FROM golang:1.21.6-alpine

# update apk repo
RUN echo "http://dl-4.alpinelinux.org/alpine/v3.7/main" >> /etc/apk/repositories && \
    echo "http://dl-4.alpinelinux.org/alpine/v3.7/community" >> /etc/apk/repositories

# copy src
COPY cmd/ /go/cmd

# builds
RUN cd /go/cmd/notif-scs/ && go build main.go
RUN cd /go/cmd/snooze-scs/ && go build main.go
RUN cd /go/cmd/webhook-scs/ && go build main.go
