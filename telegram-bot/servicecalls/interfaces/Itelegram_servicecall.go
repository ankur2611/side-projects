package interfaces

import (
	"telegram-bot/dto"

	"github.com/gin-gonic/gin"
)

type ITelegramServiceCall interface {
	SetWebhook(ctx *gin.Context) (dto.WebHookResponse, error)
	SendMessage(ctx *gin.Context, chatID int64, text string) (dto.TgResponse, error)
	SendSticker(ctx *gin.Context, chatID int64, mediaURL string) (dto.TgResponse, error)
	BanUser(ctx *gin.Context, chatID, userID int64) (dto.TgResponse, error)
	UnBanUser(ctx *gin.Context, chatID, userID int64) (dto.TgResponse, error)
	RestrictUser(ctx *gin.Context, chatID, userID int64) (dto.TgResponse, error)
	SendInvoice(ctx *gin.Context, chatID int64, invoice dto.Invoice) (dto.TgResponse, error)
	SendPreCheckoutQuery(ctx *gin.Context, data dto.SendPreCheckoutQueryRequest) (dto.TgResponse, error)
}
