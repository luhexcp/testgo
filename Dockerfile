FROM golang:1.10-alpine

LABEL                   \
   version="1.10"       \
   maintainer="xuchengpeng"

RUN                     \
    apk add --update    \
        curl            \
        git             \
        make            \
        &&              \
        mkdir /build

ADD Makefile /build

RUN make -f /build/Makefile all
