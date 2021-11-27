FROM golang:1.17.3-alpine3.14 as builder
COPY go.mod go.sum /go/src/todo/
WORKDIR /go/src/todo
RUN go mod tidy
COPY . /go/src/todo
RUN go build -o todoServer ./server

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/todo/todoServer /usr/bin/todoServer
RUN chmod a+x /usr/bin/todoServer
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/todoServer"]