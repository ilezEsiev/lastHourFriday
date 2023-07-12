package SendToTelegramFunc

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"lastHourFriday/internal/DB"
	"log"
)

func SendTg(bot *tgbotapi.BotAPI) { //Функция отправки уведомления

	// Получение всех идентификаторов чатов из Redis
	channelIDs, err := DB.GetAllKeysFromRedis()
	if err != nil {
		log.Println("Ошибка при получении идентификаторов чатов из Redis:", err)
		return
	}

	// Отправка уведомления в каждый чат
	for _, channelID := range channelIDs {
		// Отправка уведомления в чат
		msg := tgbotapi.NewMessageToChannel(channelID, "Последний час пятницы, не забуьте сделать дуа!")
		_, err := bot.Send(msg)
		if err != nil {
			log.Println("Ошибка при отправке уведомления в чат:", err)
			continue
		}

		fmt.Println("Уведомление отправлено в чат:", channelID)
	}

}
