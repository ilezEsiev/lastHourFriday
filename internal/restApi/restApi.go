package restApi

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
		Date struct {
			Readable string `json:"readable"`
		}
	} `json:"data"`
}

func Parse() (time.Time, error) {
	// Создаем переменную для хранения JSON-ответа
	var timeMaghrib Response

	// Получаем JSON-данные из REST API
	data, err := http.Get("https://api.aladhan.com/v1/timingsByCity?city=Nazran&country=Ru&method=2&school=0")
	if err != nil {
		return time.Time{}, err
	}

	dataBodyToByte, err := io.ReadAll(data.Body)
	if err != nil {
		return time.Time{}, err
	}

	// Разбираем JSON-данные и сохраняем в переменную timeMaghrib
	err = json.Unmarshal(dataBodyToByte, &timeMaghrib)
	if err != nil {
		return time.Time{}, err
	}

	// Извлекаем время вечернего намаза из данных
	maghribHourAndDateTime := strings.ToUpper(timeMaghrib.Data.Date.Readable + " " + timeMaghrib.Data.Timings.Maghrib + " +0300")
	//maghribHourTime = strings.ReplaceAll(maghribHour Time, " ", "")

	// Парсим время вечернего намаза
	MaghribHourTimeParsed, err := time.Parse("02 Jan 2006 15:04 -0700", maghribHourAndDateTime)
	if err != nil {
		return time.Time{}, err
	}

	return MaghribHourTimeParsed, nil
}
