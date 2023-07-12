package telegramBotApi

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"lastHourFriday/internal/DB"
	tk "lastHourFriday/internal/token"
	"log"
	"strconv"
)

var Bot *tgbotapi.BotAPI

func GetUpdateFromTelegram() {
	var err error
	Bot, err = tgbotapi.NewBotAPI(tk.ParseTokenTg())
	if err != nil {
		log.Panic(err)
	}
	Bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := Bot.GetUpdatesChan(u)
	ctx := context.Background()
	for update := range updates {
		if update.Message != nil && update.Message.Text == "/startFD" {
			// Получаем чат ID and ChatName
			chatID := update.Message.Chat.ID
			//Подключение к базе данных Redis
			DB.InitRedis()
			// Записываем chatID в базу данных Redis
			err := DB.Rdb.Set(ctx, strconv.Itoa(int(chatID)), strconv.Itoa(int(chatID)), 0).Err()
			if err != nil {
				log.Println(err)
			}
			// Отправляем сообщение об успешном сохранении ChatId в базу данных и запуске бота.
			msg := tgbotapi.NewMessage(chatID, "Вы запустили бота, мы напомним вам о наступлении последнего часа пятницы!")
			Bot.Send(msg)
		}
	}
}
