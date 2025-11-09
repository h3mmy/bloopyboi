# BloopyBoi Agent's Guide

This document provides guidance for AI agents working on the BloopyBoi codebase.

## Getting Started

To get started, you'll need to have Go installed on your system. You can download it from the official [Go website](https://golang.org/).

### Running the Bot

To run the bot, you'll need to set up the following environment variables:

* `DISCORD_BOT_TOKEN`: Your Discord bot token.
* `DISCORD_GUILD_ID`: The ID of the Discord guild where you want to run the bot.

Once you've set these environment variables, you can run the bot using the following command:

```bash
go run main.go
```

### Running Tests

To run the tests, use the following command:

```bash
go test ./...
```

## Coding Conventions

This project follows the standard Go coding conventions. Please make sure your code is formatted correctly before submitting a change. You can format your code using the following command:

```bash
go fmt ./...
```

## Linter

The project uses `golangci-lint`. It can be run with `~/go/bin/golangci-lint run --timeout 5m`. If not present, it can be installed with `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`.

## Project Structure

The project is organized into the following directories:

* `bot/`: Contains the bot's core logic.
* `bot/discord/`: Contains the Discord-specific logic.
* `bot/handlers/`: Contains the bot's command handlers.
* `bot/services/`: Contains the bot's services, which interact with external APIs.
* `main.go`: The application's entry point.

## How to Contribute

To contribute to the project, please follow these steps:

1. Fork the repository.
2. Create a new branch for your changes.
3. Make your changes and commit them with a descriptive commit message.
4. Push your changes to your fork.
5. Open a pull request.

### Contribution Guidelines

Please review the [Code of Conduct](CODE_OF_CONDUCT.md) and [Contribution Guidelines](CONTRIBUTING.md) before contributing. All contributions must be signed off on to certify that you have the right to submit them.
