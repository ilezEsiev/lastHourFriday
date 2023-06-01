package telegramBotApi

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

// Telegram bot configuration
var bot *tgbotapi.BotAPI

func initBot() {
	var err error
	bot, err = tgbotapi.NewBotAPI("5440075577:AAFS5UBeVWrOKOp6J0odq0NyhQJFFCuqcfg")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
}
