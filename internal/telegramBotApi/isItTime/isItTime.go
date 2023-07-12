package isItTime

import (
	"fmt"
	"time"
)

func Time(maghribHourTimeParsed time.Time) bool {
	hourBefore := maghribHourTimeParsed.Add(time.Hour)
	timeLeft := time.Until(hourBefore)

	if timeLeft > 0 && timeLeft < time.Hour {
		fmt.Println("Пробуждение! Дошло до указанного времени, до вечернего намаза осталось:", timeLeft)
		return true
	} else if timeLeft > time.Hour {
		fmt.Println("Нужно подождать :", timeLeft)
		time.Sleep(timeLeft)
		fmt.Println("Послпали до указанного времяни:", timeLeft)
		return true
	} else {
		fmt.Println("Уже прошло указанное время:", timeLeft, "... Ждем следующей пятницы.")
		return false
	}

}
