FROM golang:1.11beta2-alpine3.8 as builder

WORKDIR /go/src/github.com/vasiliy-t/gotrain/redis/webhook

COPY . .

RUN go build -v -o app


FROM golang:1.11beta2-alpine3.8

WORKDIR /app

COPY --from=builder /go/src/github.com/vasiliy-t/gotrain/redis/webhook .

CMD ["./app"]
