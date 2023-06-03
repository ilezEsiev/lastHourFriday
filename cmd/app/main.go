package main

import (
	"lastHourFriday/internal"
	"lastHourFriday/internal/DB"
	tg "lastHourFriday/internal/telegramBotApi"
	"lastHourFriday/internal/token"
)

func main() {
	token.TokenTg()
	tg.InitBot()
	DB.InitRedis()
	internal.Send()
}
