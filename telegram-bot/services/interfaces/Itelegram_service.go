package interfaces

import (
	"telegram-bot/dto"

	"github.com/gin-gonic/gin"
)

type ITelegramService interface {
	RegisterWebhook(ctx *gin.Context) (string, error)
	Commands(ctx *gin.Context, data dto.CommandsRequest) (string, error)
}
