# discord-webhook

<!-- TOC -->

- [discord-webhook](#discord-webhook)
  - [Usage](#usage)
    - [via command](#via-command)
    - [as a step in cloud build](#as-a-step-in-cloud-build)

<!-- /TOC -->
## Usage

### via command

```shell
DISCORD_WEBHOOK_URL=https://...... CONTENT=...... discord-webhook send
```

You can give a Discord Webhook URL via environment variable named DISCORD_WEBHOOK_URL.
The variable can be derived from We Config file at the path WE_CONFIG.

You can specify the content of the message in the following ways

- via CONTENT environment variable
- via file whose name is set CONTENT_FILE environment variable
- or standard input of this command.

### as a step in cloud build

Create a container in your GCP project

```shell
gcloud --project=YOUR-PROJECT builds submit --config cloudbuild.yaml .
```

Then you can use it in your cloudbuild.yaml like this

```yaml
steps:
  - name: gcr.io/$PROJECT_ID/discord-webhook
    env:
      - "DISCORD_WEBHOOK_URL=http://....."
      - "CONTENT=...."
```
