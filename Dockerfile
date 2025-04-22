FROM golang:1.24.0-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

RUN go build -o server ./cmd/server

FROM scratch

COPY ./configs /configs

COPY --from=builder /app/server /

ENTRYPOINT ["/server", "configs/dev.yaml"]