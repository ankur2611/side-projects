package interfaces

import (
	"telegram-bot/dto"

	"github.com/gin-gonic/gin"
)

type ITelegramServiceCall interface {
	SetWebhook(ctx *gin.Context) (dto.WebHookResponse, error)
	SendMessage(ctx *gin.Context, chatID, text string) (dto.TgResponse, error)
	SendSticker(ctx *gin.Context, chatID, mediaURL string) (dto.TgResponse, error)
}
