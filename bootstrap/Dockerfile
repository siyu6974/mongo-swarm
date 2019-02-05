FROM golang:1.8.3-alpine as builder

ARG APP_VERSION=unkown

# pull deps
RUN apk add --no-cache --virtual git
RUN go get -u github.com/golang/dep/cmd/dep

# copy code
ADD . /go/src/github.com/stefanprodan/mongo-swarm/bootstrap

# solution root
WORKDIR /go/src/github.com/stefanprodan/mongo-swarm/bootstrap

RUN dep ensure

# output
RUN mkdir /go/dist
RUN go build -ldflags "-X main.version=$APP_VERSION" \
    -o /go/dist/bootstrap github.com/stefanprodan/mongo-swarm/bootstrap


FROM alpine:3.9

COPY --from=builder /go/dist/bootstrap /mongo-swarm/bootstrap

RUN chmod 777 /mongo-swarm/bootstrap

EXPOSE 9090
WORKDIR /mongo-swarm
ENTRYPOINT ["/mongo-swarm/bootstrap"]
