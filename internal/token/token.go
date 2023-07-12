package token

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Token string `json:"token"`
}

func ParseTokenTg() (Token string) {
	// Чтение файла конфигурации
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal("Ошибка чтения файла конфигурации:", err)
	}

	// Распарсивание файла конфигурации
	var configjs Config
	err = json.Unmarshal(file, &configjs)
	if err != nil {
		log.Fatal("Ошибка парсинга файла конфигурации:", err)
	}
	//fmt.Print(configjs.Token)
	return configjs.Token
}
