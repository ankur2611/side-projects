package services

import (
	"fmt"
	"math/rand"
	"telegram-bot/dto"
	"telegram-bot/servicecalls"
	"telegram-bot/utils"
	"telegram-bot/utils/logger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// TelegramService struct defines the TelegramService struct type with tsc and logger fields
type TelegramService struct {
	tsc    servicecalls.TelegramServiceCall
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
	message := data.Message.Text
	t.logger.Debug("Data Received from TG Server, chatID: ", data.Message.Chat.ID, "message: ", data.Message.Text)

	switch message {
	case "/desc":
		msg := groupDescription()
		if _, err := t.tsc.SendMessage(ctx, chatID, msg); err != nil {
			t.logger.Error("error sending group description: ", err)
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
	if data.Message.LeftChatMember != nil {

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

	return "", nil
}

// groupDescription returns the description of the group
func groupDescription() string {
	return "This is a simple telegram bot"
}
