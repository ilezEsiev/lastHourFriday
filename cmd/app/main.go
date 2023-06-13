package main

import (
	"lastHourFriday/internal"
	"lastHourFriday/internal/DB"
	tg "lastHourFriday/internal/telegramBotApi"
)

func main() {
	DB.GetAllKeysFromRedis()
	DB.InitRedis()
	tg.GetUpdateFromTelegram()
	internal.Send()

}
