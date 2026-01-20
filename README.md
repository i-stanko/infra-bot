# KVN

[Kevin](https://t.me/kevin_companion_bot) is a personal infrastructure automation and monitoring Telegram bot.


## Features

* CLI tool using Cobra
* Telegram bot integration
* Dry-run mode (no Telegram connection) for safe testing
* Version command to show current build
* Docker support


## Requirements

* Go 1.25+ for local builds
* Docker 20+ for containerized run
* Telegram bot token (only needed for actual bot run)


## Project Structure

```
infra-bot/
├── cmd/                 # CLI commands
├── internal/            # Bot, config, version logic
├── main.go
├── go.mod
├── Makefile
├── Dockerfile
└── README.md
```


## Local Usage

### Build

```bash
make build
```

Builds the CLI binary with version info.

### Run bot

```bash
make run
```

Starts the Telegram bot using `TELE_TOKEN` from your environment.

### Dry-run mode

```bash
./infra-bot start --dry-run
```

Simulates bot startup without connecting to Telegram — useful for testing.

### Check version

```bash
./infra-bot version
```

Displays current version, commit hash, and build date.


## Docker Usage

### Build Docker image

```bash
make docker-build
```

### Run Docker container

```bash
make docker-run
```

Requires `TELE_TOKEN` environment variable.

### Dry-run in Docker

```bash
docker run --rm infra-bot:latest start --dry-run
```

Dry-run works without a Telegram token.


## Makefile Commands

| Command             | Description          |
| ------------------- | -------------------- |
| `make build`        | Build the CLI binary |
| `make run`          | Run the bot locally  |
| `make test`         | Run Go tests         |
| `make lint`         | Run Go vet           |
| `make clean`        | Remove binary        |
| `make docker-build` | Build Docker image   |
| `make docker-run`   | Run Docker container |


## Next Steps

* Add unit tests for `internal/bot` and `internal/config`
* Set up GitHub Actions for CI/CD
* Manage secrets (TELE_TOKEN) securely
* Optionally prepare for Kubernetes deployment


## Notes

This is a **learning pet-project**, focused on experimenting with Go, CLI tools, Docker, and DevOps practices. Dry-run mode is useful for safely testing the bot without connecting to Telegram.
