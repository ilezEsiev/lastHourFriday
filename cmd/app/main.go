package main

import (
	"lastHourFriday/internal"
	"lastHourFriday/internal/DB"
	tg "lastHourFriday/internal/telegramBotApi"
)

func main() {
	DB.InitRedis()
	tg.GetUpdateFromTelegram()
	internal.Send()

}
