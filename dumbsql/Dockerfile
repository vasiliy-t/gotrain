FROM golang:1.10-alpine as builder
WORKDIR /go/src/github.com/vasiliy-t/gotrain/dumbsql
COPY . .
RUN go build -v -o app

FROM alpine:3.6
WORKDIR /
COPY --from=builder /go/src/github.com/vasiliy-t/gotrain/dumbsql/app ./
CMD ["./app"]
