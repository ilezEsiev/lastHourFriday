package internal

import (
	"encoding/json"
	"io"
	"lastHourFriday/internal/telegramBotApi"
	"lastHourFriday/internal/telegramBotApi/SendToTelegramFunc"
	"lastHourFriday/internal/telegramBotApi/isItTime"
	"log"
	"net/http"
	"strings"
	"time"
)

type Response struct {
	Data struct {
		Timings struct {
			Maghrib string `json:"Maghrib"`
		} `json:"timings"`
	} `json:"data"`
}

func Send() {
	// Создаем переменную для хранения JSON-ответа
	var timeMaghrib Response

	// Получение текущего дня недели
	weekday := time.Now().Weekday()

	// Проверяем, является ли сегодня воскресеньем
	if weekday.String() == "Friday" {
		// Получаем JSON-данные из REST API
		data, err := http.Get("https://api.aladhan.com/v1/timingsByCity/18-06-2023?city=Nazran&country=Ru&method=2&school=0")
		if err != nil {
			log.Fatal(err)
		}

		dataBodyToByte, err := io.ReadAll(data.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Разбираем JSON-данные и сохраняем в переменную timeMaghrib
		err = json.Unmarshal(dataBodyToByte, &timeMaghrib)
		if err != nil {
			log.Fatal(err)
		}

		// Извлекаем время вечернего намаза из данных
		maghribHourTime := strings.ToUpper(timeMaghrib.Data.Timings.Maghrib)
		//maghribHourTime = strings.ReplaceAll(maghribHourTime, " ", "")

		// Парсим время вечернего намаза
		maghribHourTimeParsed, err := time.Parse("15:04", maghribHourTime)
		if err != nil {
			log.Fatal(err)
		}

		// Рассчитываем разницу между текущим временем и временем намаза
		houreBefore := maghribHourTimeParsed.Add(-time.Hour)
		hoursLeft := time.Now().Sub(houreBefore)
		time.Sleep(hoursLeft)

		// Проверяем, наступило ли время намаза
		for {
			if isItTime.Time(maghribHourTimeParsed) {
				// Отправляем сообщение с напоминанием в Telegram каналы
				SendToTelegramFunc.SendTg(telegramBotApi.Bot)
				time.Sleep((time.Hour * 24) * 7)
				break //как тут быть
			} else {
				time.Sleep(time.Hour)
			}
		}

	} else {
		// Если сегодня не воскресенье, ждем 24 часа
		time.Sleep(24 * time.Hour)
	}
}
