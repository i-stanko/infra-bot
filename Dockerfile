FROM quay.io/projectquay/golang:1.22 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go build -o infra-bot \
	-ldflags "-X github.com/i-stanko/infra-bot/cmd.appVersion=container" .

FROM scratch
WORKDIR /
COPY --from=builder /app/infra-bot /infra-bot
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/infra-bot"]
CMD ["start"]
