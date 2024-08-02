package services

import (
	"fmt"
	"telegram-bot/dto"
	"telegram-bot/servicecalls"
	"telegram-bot/utils"

	"github.com/gin-gonic/gin"
)

type TelegramService struct {
	tsc servicecalls.TelegramServiceCall
}

func NewTelegramService(ctx *gin.Context) TelegramService {
	return TelegramService{
		tsc: servicecalls.NewTelegramServiceCall(ctx),
	}
}

func (t TelegramService) RegisterWebhook(ctx *gin.Context) (string, error) {
	data, err := t.tsc.SetWebhook(ctx)

	if err != nil {
		return "", fmt.Errorf("error setting webhook: %s", err)
	}

	return data.Description, nil
}

func (t TelegramService) Commands(ctx *gin.Context, data dto.CommandsRequest) (string, error) {

	chatID := data.Message.Chat.ID
	fmt.Println(chatID)

	// Greet new chat members
	for _, newMember := range data.Message.NewChatMembers {

		fullName := t.constructFullName(newMember.FirstName, newMember.LastName, newMember.Username)
		greetingMsg := fmt.Sprintf("Welcome, %s !", fullName)

		_, err := t.tsc.SendMessage(ctx, chatID, greetingMsg)
		if err != nil {
			return greetingMsg, fmt.Errorf("error sending greeting message: %s", err)
		}
		_, err = t.tsc.SendSticker(ctx, chatID, "https://media.giphy.com/media/3o7TKz9bX9v9Kz2wJi/giphy.gif")
		if err != nil {
			return greetingMsg, fmt.Errorf("error sending greeting message: %s", err)
		}
	}

	// Handle members leaving
	if data.Message.LeftChatMember != nil {

		fullName := t.constructFullName(data.Message.LeftChatMember.FirstName,
			data.Message.LeftChatMember.LastName, data.Message.LeftChatMember.Username)

		farewellMsg := fmt.Sprintf("Goodbye, %s !", fullName)

		_, err := t.tsc.SendMessage(ctx, chatID, farewellMsg)
		if err != nil {
			return farewellMsg, fmt.Errorf("error sending farewell message: %s", err)
		}
	}
	return "", nil
}

// func (t TelegramService) SendPhoto(ctx *gin.Context, photo string) (string, error) {
// 	_, err := bot.TgBot.SendMessage(req.ChatID, req.Message, nil)
// }

// func (t TelegramService) SendUpdates(ctx *gin.Context, chatID int64, message string) (string, error) {
// _, err := bot.TgBot.SendMessage(chatID, message, nil)
// if err != nil {
// 	return "", fmt.Errorf("error sending message: %s", err)
// }
// return "Message sent", nil
// }

func (t TelegramService) constructFullName(firstName, lastName, userName string) string {

	name := ""
	if !utils.IsEmpty(firstName) {
		name = firstName
	}
	if !utils.IsEmpty(lastName) {
		name += " " + lastName
	}
	if !utils.IsEmpty(userName) {
		name += " (@" + userName + ")"
	}
	return name
}
