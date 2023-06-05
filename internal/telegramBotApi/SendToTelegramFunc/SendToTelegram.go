package SendToTelegramFunc

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func Send(bot *tgbotapi.BotAPI) { //Функция отправки уведомления

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Последний час пятницы, не забуьте сделать дуа!")
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)

		}
	}

}