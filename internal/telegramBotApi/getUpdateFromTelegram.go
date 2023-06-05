package telegramBotApi

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"lastHourFriday/internal/DB"
	tk "lastHourFriday/internal/token"
	"log"
	"strconv"
)

var Bot *tgbotapi.BotAPI

func GetUpdateFromTelegram() {
	var err error
	Bot, err = tgbotapi.NewBotAPI(tk.TokenTg())
	if err != nil {
		log.Panic(err)
	}
	Bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := Bot.GetUpdatesChan(u)
	ctx := context.Background()
	for update := range updates {
		if update.Message != nil && update.Message.Text == "/startFridayBot" {
			// Получаем чат ID and ChatName
			chatID := update.Message.Chat.ID
			chatName := update.Message.Chat.FirstName

			// Записываем chatID в базу данных Redis
			err := DB.Rdb.Set(ctx, chatName, strconv.Itoa(int(chatID)), 0).Err()
			if err != nil {
				log.Println(err)
			}

			// Отправляем сообщение об успешном сохранении ChatId в базу данных и запуске бота.
			msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("%s, бот запушен! ", chatName))
			Bot.Send(msg)
		}
	}
}
