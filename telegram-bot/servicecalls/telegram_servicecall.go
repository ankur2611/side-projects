package servicecalls

import (
	"encoding/json"
	"fmt"
	"telegram-bot/dto"
	"telegram-bot/servicecalls/interfaces"
	"telegram-bot/utils"
	"telegram-bot/utils/endpoints"
	"time"

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

// https://core.telegram.org/bots/api#setwebhook
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

// https://core.telegram.org/bots/api#sendmessage
func (g TelegramServiceCall) SendMessage(ctx *gin.Context, chatID int64, text string) (dto.TgResponse, error) {

	var response dto.TgResponse
	body := dto.SendMessageRequest{
		ChatID: chatID,
		Text:   text,
	}

	endpoint := fmt.Sprintf("%s%s", g.url, fmt.Sprintf(endpoints.TG_SEND_MESSAGE_URL, g.apiKey))

	byteResponse, err := g.sc.Post(ctx, endpoint, body, nil)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(byteResponse, &response); err != nil {
		return response, err
	}
	return response, nil
}

// https://core.telegram.org/bots/api#sendsticker
func (g TelegramServiceCall) SendSticker(ctx *gin.Context, chatID int64, mediaURL string) (dto.TgResponse, error) {

	var response dto.TgResponse
	body := dto.SendStickerRequest{
		ChatID:  chatID,
		Sticker: mediaURL,
	}

	endpoint := fmt.Sprintf("%s%s", g.url, fmt.Sprintf(endpoints.TG_SEND_STICKER_URL, g.apiKey))

	byteResponse, err := g.sc.Post(ctx, endpoint, body, nil)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(byteResponse, &response); err != nil {
		return response, err
	}
	return response, nil
}

// https://core.telegram.org/bots/api#banchatmember
func (g TelegramServiceCall) BanUser(ctx *gin.Context, chatID, userID int64) (dto.TgResponse, error) {

	var response dto.TgResponse

	body := dto.BanUserRequest{
		ChatID:         chatID,
		UserID:         userID,
		UntilDate:      utils.GetUnixTime(time.Now().AddDate(0, 0, 30)),
		RevokeMessages: true,
	}

	endpoint := fmt.Sprintf("%s%s", g.url, fmt.Sprintf(endpoints.TG_BAN_USER, g.apiKey))

	byteResponse, err := g.sc.Post(ctx, endpoint, body, nil)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(byteResponse, &response); err != nil {
		return response, err
	}
	return response, nil
}

// https://core.telegram.org/bots/api#unbanchatmember
func (g TelegramServiceCall) UnBanUser(ctx *gin.Context, chatID, userID int64) (dto.TgResponse, error) {

	var response dto.TgResponse
	body := dto.UnbanUserRequest{
		ChatID:       chatID,
		UserID:       userID,
		OnlyIfBanned: true,
	}

	endpoint := fmt.Sprintf("%s%s", g.url, fmt.Sprintf(endpoints.TG_BAN_USER, g.apiKey))

	byteResponse, err := g.sc.Post(ctx, endpoint, body, nil)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(byteResponse, &response); err != nil {
		return response, err
	}
	return response, nil
}

// https://core.telegram.org/bots/api#restrictchatmember
func (g TelegramServiceCall) RestrictUser(ctx *gin.Context, chatID, userID int64) (dto.TgResponse, error) {

	var response dto.TgResponse
	body := dto.RestrictUser{
		ChatID: chatID,
		UserID: userID,
		Permissions: dto.Permissions{
			CanSendMessages:       false,
			CanSendPolls:          false,
			CanSendOtherMessages:  false,
			CanAddWebPagePreviews: false,
			CanChangeInfo:         false,
			CanInviteUsers:        false,
			CanPinMessages:        false,
			CanSendAudios:         false,
			CanSendDocuments:      false,
			CanSendPhotos:         false,
			CanSendVideos:         false,
		},
		UseIndependentChatPermissions: true,
		UntilDate:                     utils.GetUnixTime(time.Now().AddDate(0, 0, 30)),
	}

	endpoint := fmt.Sprintf("%s%s", g.url, fmt.Sprintf(endpoints.TG_RESTRICT_USER, g.apiKey))

	byteResponse, err := g.sc.Post(ctx, endpoint, body, nil)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(byteResponse, &response); err != nil {
		return response, err
	}
	return response, nil
}
