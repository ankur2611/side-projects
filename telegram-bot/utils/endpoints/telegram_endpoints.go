package endpoints

// Telegram APIs
const (
	TG_WEBHOOK_URL      = "/bot%s/setWebhook?url=%s"
	TG_SEND_MESSAGE_URL = "/bot%s/sendMessage"
	TG_SEND_STICKER_URL = "/bot%s/sendSticker"
	TG_BAN_USER         = "/bot%s/banChatMember"
	TG_UNBAN_USER       = "/bot%s/unbanChatMember"
	TG_RESTRICT_USER    = "/bot%s/restrictChatMember"
)
