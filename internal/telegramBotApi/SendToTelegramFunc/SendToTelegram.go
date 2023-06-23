package SendToTelegramFunc

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/redis/go-redis/v9"
	"log"
)

func SendTg(bot *tgbotapi.BotAPI) { //Функция отправки уведомления

	// Подключение к базе Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // Если не требуется пароль для доступа
		DB:       0,  // Индекс базы данных
	})

	defer redisClient.Close()

	// Получение всех идентификаторов чатов из Redis
	channelIDs, err := redisClient.LRange(context.Background(), "telegram_channels", 0, -1).Result()
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
