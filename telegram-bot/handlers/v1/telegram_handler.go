package handlers_v1

import (
	"net/http"
	"telegram-bot/dto"
	"telegram-bot/services"

	"github.com/gin-gonic/gin"
)

type TelegramHandler struct {
	ts services.TelegramService
}

func NewTelegramHandler(ctx *gin.Context) TelegramHandler {
	return TelegramHandler{
		ts: services.NewTelegramService(ctx),
	}
}

func (t TelegramHandler) RegisterWebhook(ctx *gin.Context) {

	response, err := t.ts.RegisterWebhook(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": response})
}

func (t TelegramHandler) CommandsHandler(c *gin.Context) {

	var data dto.CommandsRequest
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid update format"})
		return
	}

	_, err := t.ts.Commands(c, data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (t TelegramHandler) SendUpdatesHandler(c *gin.Context) {
	var req struct {
		ChatID  int64  `json:"chat_id"`
		Message string `json:"message"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "message sent"})
}
