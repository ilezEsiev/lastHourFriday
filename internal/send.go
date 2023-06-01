package internal

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
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

func send() {
	/*bot, err := tgbotapi.NewBotAPI("5440075577:AAFS5UBeVWrOKOp6J0odq0NyhQJFFCuqcfg")
	if err != nil {
		log.Panic(err)
	}*/
	var timeMaghrib TimeSalat
	weekday := time.Now().Weekday()
	//Проверям сегодня пятница
	if weekday.String() == "Wednesday" {
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
			if isItTime(maghribHourTimeParsed) {
				SendToTelegram(bot)
				return
			}
		}

	}

}
func isItTime(maghribHourTimeParsed time.Time) bool {
	if time.Now().Sub(maghribHourTimeParsed) <= time.Hour && time.Now().Sub(maghribHourTimeParsed) > 0 {
		return true
	}
	return false
}

func SendToTelegram(bot *tgbotapi.BotAPI) {

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Последний час пятницы, не забуьте сделать дуа!")
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)

		}
	}

}
