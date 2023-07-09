package internal

import (
	"encoding/json"
	"io"
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

func Send() error {
	// Создаем переменную для хранения JSON-ответа
	var timeMaghrib Response

	// Получаем JSON-данные из REST API
	data, err := http.Get("https://api.aladhan.com/v1/timingsByCity/18-06-2023?city=Nazran&country=Ru&method=2&school=0")
	if err != nil {
		return err
	}

	dataBodyToByte, err := io.ReadAll(data.Body)
	if err != nil {
		return err
	}

	// Разбираем JSON-данные и сохраняем в переменную timeMaghrib
	err = json.Unmarshal(dataBodyToByte, &timeMaghrib)
	if err != nil {
		return err
	}

	// Извлекаем время вечернего намаза из данных
	maghribHourTime := strings.ToUpper(timeMaghrib.Data.Timings.Maghrib)
	//maghribHourTime = strings.ReplaceAll(maghribHourTime, " ", "")

	// Парсим время вечернего намаза
	MaghribHourTimeParsed, err := time.Parse("15:04", maghribHourTime)
	if err != nil {
		return err
	}

	// Рассчитываем разницу между текущим временем и временем намаза
	hourBefore := MaghribHourTimeParsed.Add(-time.Hour)
	hoursLeft := time.Now().Sub(hourBefore)
	time.Sleep(hoursLeft)
	return nil
}
