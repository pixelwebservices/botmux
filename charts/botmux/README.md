# BotMux Helm Chart

A simple Helm chart to deploy BotMux with persistent storage on Kubernetes.

## Installation

```bash
helm install botmux ./charts/botmux \
  --set telegramBotToken="YOUR_TELEGRAM_BOT_TOKEN"
```

## Values Configuration

You can customize the deployment in `values.yaml`:

- `image`: The docker image to run.
- `telegramBotToken`: Your Telegram bot API token.
- `cookieSecure`: Set to `"false"` to allow session cookies over plain HTTP (essential for local dev / port forwarding). Leave blank for auto-detection based on request headers (`TLS` / `X-Forwarded-Proto`).
- `persistence.size`: The storage size for the SQLite database (`1Gi` by default).
- `persistence.storageClass`: The Kubernetes storage class to use (uses cluster default if empty).
