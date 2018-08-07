FROM golang

WORKDIR /go/src/app

ADD hibtc_websocket.go hibtc_websocket.go
ADD redis_init.go redis_init.go

RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT ["go", "run", "hibtc_websocket.go", "redis_init.go"]