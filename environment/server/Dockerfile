FROM golang:1.15-alpine as debug

RUN mkdir -p /server \
    mkdir -p /var/log/app && \
    chmod 755 -R /var/log && \
    apk update && \
    apk add curl \
            git \
            bash \
            make && \
    rm -rf /var/cache/apk/*

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

###########PRODUCTION IMAGE###################

FROM golang:1.15-alpine as prod
