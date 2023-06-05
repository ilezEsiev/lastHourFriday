package token

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Token string `json:"token"`
}

func TokenTg() (Token string) {
	// Чтение файла конфигурации
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal("Ошибка чтения файла конфигурации:", err)
	}

	// Распарсивание файла конфигурации
	var Configjs Config
	err = json.Unmarshal(file, &Configjs)
	if err != nil {
		log.Fatal("Ошибка парсинга файла конфигурации:", err)
	}

	return Configjs.Token
}
