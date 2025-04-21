FROM golang:1.24.0-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o server ./cmd/server

FROM scratch

COPY ./configs /configs

COPY --from=builder /app/server /

ENTRYPOINT ["/server", "configs/dev.yaml"]