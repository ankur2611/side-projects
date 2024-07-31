package main

import (
	"fmt"
	"net/http"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Text           string `json:"text"`
	Chat           Chat   `json:"chat"`
	NewChatMembers []User `json:"new_chat_members"`
	MessageID      int    `json:"message_id"`
	From           User   `json:"from"`
	Date           int    `json:"date"`
	LeftChatMember *User  `json:"left_chat_member,omitempty"`
}

type Chat struct {
	ID int64 `json:"id"`
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

var bot *gotgbot.Bot

func main() {

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	app := gin.Default()

	opts := gotgbot.BotOpts{}
	bot, err = gotgbot.NewBot(viper.GetString("BOT_TOKEN"), &opts)

	if err != nil {
		panic(err)
	}

	// Set up webhook with all types of updates
	webhookURL := viper.GetString("WEBHOOK_URL")
	_, err = bot.SetWebhook(webhookURL, &gotgbot.SetWebhookOpts{
		AllowedUpdates: []string{"message", "my_chat_member", "chat_member"},
	})

	if err != nil {
		panic(err)
	}

	app.POST("/webhook", UpdatesHandler)
	app.POST("/send", SendUpdatesHandler)

	app.Run(":9090")
}

func UpdatesHandler(c *gin.Context) {

	var update Update
	if err := c.BindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid update format"})
		return
	}

	chatID := update.Message.Chat.ID
	messageText := update.Message.Text

	fmt.Printf("Received update: %+v\n", update)

	// Check if there are new chat members
	if len(update.Message.NewChatMembers) > 0 {
		for _, newMember := range update.Message.NewChatMembers {
			greeting := fmt.Sprintf("Welcome, %s!", newMember.FirstName)
			if newMember.LastName != "" {
				greeting += " " + newMember.LastName
			}
			if newMember.Username != "" {
				greeting += " (@" + newMember.Username + ")"
			}
			_, err := bot.SendMessage(update.Message.Chat.ID, greeting, nil)
			if err != nil {
				fmt.Printf("Error sending greeting message: %s\n", err)
			}
		}
	}

	// Handle members leaving
	if update.Message.LeftChatMember != nil {
		farewell := fmt.Sprintf("Goodbye, %s!", update.Message.LeftChatMember.FirstName)
		if update.Message.LeftChatMember.LastName != "" {
			farewell += " " + update.Message.LeftChatMember.LastName
		}
		if update.Message.LeftChatMember.Username != "" {
			farewell += " (@" + update.Message.LeftChatMember.Username + ")"
		}
		_, err := bot.SendMessage(update.Message.Chat.ID, farewell, nil)
		if err != nil {
			fmt.Printf("Error sending farewell message: %s\n", err)
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})

	fmt.Println(chatID, messageText)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func SendUpdatesHandler(c *gin.Context) {
	var req struct {
		ChatID  int64  `json:"chat_id"`
		Message string `json:"message"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	_, err := bot.SendMessage(req.ChatID, req.Message, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "message sent"})
}

func handleTextMessage(message Message) {
	fmt.Printf("Text message from %s: %s\n", message.From.Username, message.Text)
	// Process the message as needed
}
