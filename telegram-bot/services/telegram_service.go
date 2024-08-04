package services

import (
	"fmt"
	"math/rand"
	"strings"
	"telegram-bot/dto"
	"telegram-bot/servicecalls"
	"telegram-bot/servicecalls/interfaces"
	"telegram-bot/utils"
	"telegram-bot/utils/constants"
	"telegram-bot/utils/logger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// TelegramService struct defines the TelegramService struct type with tsc and logger fields
type TelegramService struct {
	tsc    interfaces.ITelegramServiceCall
	gs     GiphyService
	logger *logrus.Logger
}

// NewTelegramService creates a new instance of TelegramService
func NewTelegramService(ctx *gin.Context) TelegramService {
	return TelegramService{
		tsc:    servicecalls.NewTelegramServiceCall(ctx),
		gs:     NewGiphyService(ctx),
		logger: logger.GetLogger(),
	}
}

// RegisterWebhook registers the webhook for the bot to receive updates from Telegram
func (t TelegramService) RegisterWebhook(ctx *gin.Context) (string, error) {
	data, err := t.tsc.SetWebhook(ctx)

	if err != nil {
		t.logger.Error("error setting webhook: ", err)
		return "", fmt.Errorf("error setting webhook: %s", err)
	}

	return data.Description, nil
}

// Commands handles the commands received from the bot
func (t TelegramService) Commands(ctx *gin.Context, data dto.CommandsRequest) (string, error) {

	chatID := data.Message.Chat.ID
	userID := data.Message.From.ID
	message := data.Message.Text
	preCheckoutQuery := data.PreCheckoutQuery
	t.logger.Debug("Data Received from TG Server, chatID: ", data.Message.Chat.ID, " message: ", data.Message.Text)

	switch message {
	case "/desc":
		msg := groupDescription()
		if _, err := t.tsc.SendMessage(ctx, chatID, msg); err != nil {
			t.logger.Error("error sending group description: ", err)
		}
	case "/restrict":
		if _, err := t.tsc.RestrictUser(ctx, chatID, userID); err != nil {
			t.logger.Error("error restricting user: ", err)
		}
	case "/invoice":
		invoice := dto.Invoice{
			ChatID:        chatID,
			Title:         "Sample Invoice",
			Description:   "This is a sample invoice",
			Payload:       "sample_invoice",
			ProviderToken: viper.GetString("STRIPE_TOKEN"),
			Currency:      "JPY",
			Prices: []struct {
				Label  string `json:"label"`
				Amount int    `json:"amount"`
			}{
				{
					Label:  "Sample Item",
					Amount: 150,
				},
			},
		}
		if _, err := t.tsc.SendInvoice(ctx, chatID, invoice); err != nil {
			t.logger.Error("error sending invoice: ", err)
		}
	}

	// Greet new chat members with text msg & sticker
	for _, newMember := range data.Message.NewChatMembers {

		// Construct fullname & username of the new member
		fullName := utils.ConstructFullName(newMember.FirstName, newMember.LastName, newMember.Username)
		greetingMsg := fmt.Sprintf("Welcome, %s !", fullName)

		// Send a greeting message
		if _, err := t.tsc.SendMessage(ctx, chatID, greetingMsg); err != nil {
			t.logger.Error("error sending greeting message: ", err)
		}

		// Fetch a random sticker from Giphy.com
		stickerUrls, err := t.gs.SearchStickers(ctx, "welcome", "en")
		randomStickerUrl := stickerUrls[rand.Intn(len(stickerUrls))]

		if err != nil {
			t.logger.Error("error fetching random gif: ", err)
		}

		// Send a sticker
		if _, err := t.tsc.SendSticker(ctx, chatID, randomStickerUrl); err != nil {
			t.logger.Error("error sending sticker: ", err)
		}
	}

	// Handle members leaving
	if data.Message.LeftChatMember != (dto.User{}) {

		// Construct fullname & username of the new member
		fullName := utils.ConstructFullName(data.Message.LeftChatMember.FirstName,
			data.Message.LeftChatMember.LastName, data.Message.LeftChatMember.Username)

		farewellMsg := fmt.Sprintf("Goodbye, %s !", fullName)

		// Send a farewell message
		if _, err := t.tsc.SendMessage(ctx, chatID, farewellMsg); err != nil {
			t.logger.Error("error sending farewell message: ", err)
			return farewellMsg, fmt.Errorf("error sending farewell message: %s", err)
		}
	}

	// Ban user for offensive language
	for _, word := range constants.OffensiveWords {
		if strings.Contains(strings.ToLower(message), word) {
			if _, err := t.tsc.BanUser(ctx, chatID, userID); err != nil {
				t.logger.Error("error banning user: ", err)
			}
		}
	}

	if preCheckoutQuery != (dto.PreCheckoutQuery{}) {

		checkoutData := dto.SendPreCheckoutQueryRequest{
			PreCheckoutQueryID: preCheckoutQuery.ID,
			Ok:                 true,
		}
		if _, err := t.tsc.SendPreCheckoutQuery(ctx, checkoutData); err != nil {
			t.logger.Error("error sending precheckout query: ", err)
		}
	}
	return "", nil
}

// groupDescription returns the description of the group
func groupDescription() string {
	return "This is a sample telegram bot"
}
