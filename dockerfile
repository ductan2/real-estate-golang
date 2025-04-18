FROM golang:1.24.0-alpine as builder

WORKDIR /app

RUN apk add --no-cache --virtual .build-deps git
RUN go install github.com/cosmtrek/air@latest

COPY . . 

RUN go mod download
RUN go build -o server ./cmd/server

FROM scratch

COPY ./configs /configs
COPY ./.env /

COPY --from=builder /app/server /
COPY --from=builder /go/bin/air /usr/local/bin/air

ENTRYPOINT ["air", "-c", "/app/.air.toml"]
