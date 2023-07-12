package DB

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func GetAllKeysFromRedis() ([]string, error) {
	// Подключение к Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer client.Close()
	ctx := context.Background()

	var cursor uint64
	var values []string

	for {
		var keys []string
		var err error

		// Использование команды "SCAN" для получения пакета значений
		keys, cursor, err = client.Scan(ctx, cursor, "*", 10).Result()
		if err != nil {
			return nil, err
		}

		// Использование команды "MGET" для получения значений по ключам
		results, err := client.MGet(ctx, keys...).Result()
		if err != nil {
			return nil, err
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
	return values, nil
}
