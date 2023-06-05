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

type Item struct {
	MagrhibTime string `json:"maghrib"`
}
type TimeSalat struct {
	Items []Item `json:"items"`
}

func Send() {

	var timeMaghrib TimeSalat
	weekday := time.Now().Weekday()
	//Проверям сегодня пятница
	if weekday.String() == "Monday" {
		// Получаем json с Rest API
		data, err := http.Get("https://muslimsalat.com/nazran/daily.json?key=906a413e13c24f0c43459ed9f04cb0e2")
		if err != nil {
			log.Fatal(err)
		}

		dataBodyToByte, err := io.ReadAll(data.Body)

		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(dataBodyToByte, &timeMaghrib)
		if err != nil {
			log.Fatal(err)
		}
		// Из всех данных достаем время вечернего намаза
		maghribHourTime := strings.ToUpper(timeMaghrib.Items[0].MagrhibTime)
		maghribHourTime = strings.ReplaceAll(maghribHourTime, " ", "")
		maghribHourTimeParsed, err := time.Parse(time.Kitchen, maghribHourTime) //Время маг1риб намаза
		if err != nil {
			log.Fatal(err)
		}
		//
		houreBefore := maghribHourTimeParsed.Add(-time.Hour)
		hoursLeft := time.Now().Sub(houreBefore)
		time.Sleep(hoursLeft)
		for {
			if isItTime.Time(maghribHourTimeParsed) {
				SendToTelegramFunc.Send(telegramBotApi.Bot)
				return
			}
		}

	}

}
