package telegramBotApi

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tk "lastHourFriday/internal/token"
	"log"
)

// Telegram bot configuration
var Bot *tgbotapi.BotAPI

func InitBot() {
	var err error
	Bot, err = tgbotapi.NewBotAPI(tk.TokenTg().)
	if err != nil {
		log.Panic(err)
	}
	Bot.Debug = true
}
