FROM golang

WORKDIR /go/src/app

ADD bittrex_websocket.go bittrex_websocket.go
ADD redis_init.go redis_init.go

RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT ["go", "run", "bittrex_websocket.go", "redis_init.go"]

