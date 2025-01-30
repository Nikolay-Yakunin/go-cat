FROM golang:1.22.6-alpine as builder

RUN apk add --no-cache git make

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN make build-linux

FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /root/

COPY --from=builder /app/build/linux/go-cat .

ENTRYPOINT ["./go-cat"]