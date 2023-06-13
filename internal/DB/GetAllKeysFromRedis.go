package DB

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func GetAllKeysFromRedis() []string {
	// Подключение к Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Адрес Redis сервера
		Password: "",               // Пароль, если установленzzz
		DB:       0,                // Выбор базы данных
	})
	defer client.Close()
	// Создание контекста
	// Создание контекста
	ctx := context.Background()

	// Итерация через все значения в Redis
	var cursor uint64
	var values []string

	for {
		var keys []string
		var err error

		// Использование команды "SCAN" для получения пакета значений
		keys, cursor, err = client.Scan(ctx, cursor, "*", 10).Result()
		if err != nil {
			panic(err)
		}

		// Использование команды "MGET" для получения значений по ключам
		results, err := client.MGet(ctx, keys...).Result()
		if err != nil {
			panic(err)
		}

		// Добавление значений в общий срез
		for _, value := range results {
			if value != nil {
				values = append(values, value.(string))
			}
		}

		// Прекращение итерации, если достигнут конец
		if cursor == 0 {
			break
		}
	}

	// Вывод всех значений
	for _, value := range values {
		fmt.Println(value)
	}
	return values
}
