package token

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Config struct {
	Tokenstr string `json:"tokenstr"`
}

func TokenTg() {
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

	// Использование значения токена
	fmt.Println("Токен:", Configjs.Tokenstr)
}
