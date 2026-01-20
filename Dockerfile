FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -ldflags "-X github.com/i-stanko/infra-bot/internal/version.Version=0.1.0" -o infra-bot ./main.go

FROM gcr.io/distroless/base-debian11

USER nonroot:nonroot
WORKDIR /app
COPY --from=builder /app/infra-bot .

ENTRYPOINT ["/app/infra-bot"]
CMD ["--help"]
