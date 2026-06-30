# BotMux Helm Chart

A simple Helm chart to deploy BotMux with persistent storage on Kubernetes.

## Installation

```bash
helm install botmux ./charts/botmux \
  --set telegramBotToken="YOUR_TELEGRAM_BOT_TOKEN"
```

## Values Configuration

You can customize the deployment in `values.yaml`:

- `image`: The docker image to run (`ghcr.io/skrashevich/botmux:main` by default).
- `telegramBotToken`: Your Telegram bot API token.
- `persistence.size`: The storage size for the SQLite database (`1Gi` by default).
- `persistence.storageClass`: The Kubernetes storage class to use (uses cluster default if empty).
