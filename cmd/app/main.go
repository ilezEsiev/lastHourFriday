package main

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {
	initRedis()
	initBot()

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil && update.Message.Text == "/startFriday" {
			// Get chat ID from the incoming message
			chatID := update.Message.Chat.ID
			// Add chat ID to the Redis database
			err := rdb.SAdd(context.Background(), "channels", chatID).Err()
			if err != nil {
				log.Println(err)
			}

			// Send confirmation message to the user
			msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Ваш чат ид %d дабавлен в базу данных!", chatID))
			bot.Send(msg)
		}
	}
}
