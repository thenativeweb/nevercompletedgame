FROM golang:1.21-alpine AS builder

RUN apk update

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o /go/bin/ncg

# ---

FROM alpine:3.18

COPY --from=builder /go/bin/ncg /usr/local/bin/ncg

ENTRYPOINT [ "ncg" ]
