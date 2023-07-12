package telegramBotApi

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"lastHourFriday/internal/DB"
	tk "lastHourFriday/internal/token"
	"log"
)

func GetUpdateFromTelegram() {
	Bot, err := tgbotapi.NewBotAPI(tk.ParseTokenTg())
	if err != nil {
		log.Panic(err)
	}

	Bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil && update.Message.Text == "/startFD" {
			// Получаем chat ID и ChatName
			chatID := update.Message.Chat.ID

			// Записываем chat ID в текстовый файл
			err := DB.SaveChatIDToFile(chatID)
			if err != nil {
				log.Println(err)
				return
			}

			// Отправляем сообщение об успешном сохранении ChatID в файл и запуске бота
			msg := tgbotapi.NewMessage(chatID, "Вы запустили бота, мы напомним вам о наступлении последнего часа пятницы!")
			Bot.Send(msg)
		}
	}
}
