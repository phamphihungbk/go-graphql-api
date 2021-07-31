FROM golang:1.15-alpine

RUN mkdir -p /server \
    mkdir -p /var/log/app && \
    chmod 755 -R /var/log && \
    apk update && \
    apk add curl \
            git \
            bash \
            make && \
    rm -rf /var/cache/apk/*
