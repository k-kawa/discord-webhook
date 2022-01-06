package discord

type WebhookURL string

// WebhooksPost is https://discord.com/developers/docs/resources/webhook#execute-webhook
type WebhooksPostRequest struct {
	Content   string `json:"content"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
}
