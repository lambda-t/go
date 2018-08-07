FROM golang

WORKDIR /go/src/app

ADD binance_websocket.go binance_websocket.go
ADD redis_init.go redis_init.go

RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT ["go", "run", "binance_websocket.go", "redis_init.go"]

