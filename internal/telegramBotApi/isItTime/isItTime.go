package isItTime

import "time"

func Time(maghribHourTimeParsed time.Time) bool { //Функция проверки времени
	if time.Now().Sub(maghribHourTimeParsed) <= time.Hour && time.Now().Sub(maghribHourTimeParsed) > 0 {
		return true
	}
	return false
}
