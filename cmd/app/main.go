package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"lastHourFriday/internal/restApi"
	"lastHourFriday/internal/telegramBotApi"
	"lastHourFriday/internal/telegramBotApi/SendToTelegramFunc"
	"lastHourFriday/internal/telegramBotApi/isItTime"
	tk "lastHourFriday/internal/token"
	"log"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		telegramBotApi.GetUpdateFromTelegram()
		wg.Done()
	}()

	go func() {
		for {
			// смортрим если уже пятница, то запускаем
			if time.Now().Weekday().String() == "Wednesday" {
				isTime, err := restApi.Parse()

				if err != nil {
					log.Println(err)
				}

				if isItTime.Time(isTime) {

					Bot, err := tgbotapi.NewBotAPI(tk.ParseTokenTg())
					if err != nil {
						log.Println(err)
					}
					SendToTelegramFunc.SendTg(Bot)
				}
				time.Sleep(time.Hour * 24 * 5)
			}
			time.Sleep(time.Hour * 12)
		}
	}()

	wg.Wait()
}
