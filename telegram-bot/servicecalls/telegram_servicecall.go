package servicecalls

import (
	"encoding/json"
	"fmt"
	"telegram-bot/dto"
	"telegram-bot/servicecalls/interfaces"
	"telegram-bot/utils/endpoints"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type TelegramServiceCall struct {
	sc     interfaces.IServiceCalls
	url    string
	apiKey string
}

func NewTelegramServiceCall(ctx *gin.Context) TelegramServiceCall {
	return TelegramServiceCall{
		sc:     NewServiceCall(ctx),
		url:    viper.GetString("TELEGRAM_URL"),
		apiKey: viper.GetString("TELEGRAM_BOT_TOKEN"),
	}
}

func (g TelegramServiceCall) SetWebhook(ctx *gin.Context) (dto.WebHookResponse, error) {

	var response dto.WebHookResponse
	endpoint := fmt.Sprintf("%s%s", g.url, fmt.Sprintf(endpoints.TG_WEBHOOK_URL, g.apiKey,
		viper.GetString("MY_WEBHOOK_URL")))

	byteResponse, err := g.sc.Post(ctx, endpoint, nil, nil)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(byteResponse, &response); err != nil {
		return response, err
	}
	return response, nil
}

func (g TelegramServiceCall) SendMessage(ctx *gin.Context, chatID int64, text string) (dto.TgResponse, error) {

	var response dto.TgResponse
	endpoint := fmt.Sprintf("%s%s", g.url, fmt.Sprintf(endpoints.TG_SEND_MESSAGE_URL, g.apiKey))

	byteResponse, err := g.sc.Post(ctx, endpoint, nil, nil)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(byteResponse, &response); err != nil {
		return response, err
	}
	return response, nil
}

func (g TelegramServiceCall) SendSticker(ctx *gin.Context, chatID int64, mediaURL string) (dto.TgResponse, error) {

	var response dto.TgResponse
	endpoint := fmt.Sprintf("%s%s", g.url, fmt.Sprintf(endpoints.TG_SEND_STICKER_URL, g.apiKey))

	byteResponse, err := g.sc.Post(ctx, endpoint, nil, nil)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(byteResponse, &response); err != nil {
		return response, err
	}
	return response, nil
}
