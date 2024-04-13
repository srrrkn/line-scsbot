FROM golang:1.21.6-alpine

# update apk repo
RUN echo "http://dl-4.alpinelinux.org/alpine/v3.7/main" >> /etc/apk/repositories && \
    echo "http://dl-4.alpinelinux.org/alpine/v3.7/community" >> /etc/apk/repositories

# copy src
COPY cmd/ /cmd

# builds
RUN cd /cmd/notif-scs/ && go build main.go
RUN cd /cmd/snooze-scs/ && go build /cmd/snooze-scs/main.go
RUN cd /cmd/webhook-scs/ && go build /cmd/webhook-scs/main.go
