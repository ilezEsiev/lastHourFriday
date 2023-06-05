package DB

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

func GetAllKeysFromRedis() ([]string, error) {
	// Создаем клиента Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Укажите адрес и порт Redis сервера
		Password: "",               // Укажите пароль, если он установлен
		DB:       0,                // Укажите номер базы данных Redis, если требуется
	})

	// Проверяем подключение к Redis
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	// Инициализируем пустой список ключей
	keys := []string{}

	// Итерируем по страницам с помощью команды SCAN
	var cursor uint64
	for {
		// Выполняем команду SCAN для получения следующей страницы ключей
		keysPage, cur, err := client.Scan(cursor, "", 0).Result()
		if err != nil {
			return nil, err
		}

		// Добавляем ключи текущей страницы в общий список ключей
		keys = append(keys, keysPage...)

		// Обновляем текущий указатель курсора
		cursor = cur

		// Если курсор равен 0, значит достигнут конец итерации
		if cursor == 0 {
			break
		}
	}

	// Закрываем соединение с Redis
	err = client.Close()
	if err != nil {
		log.Println("Failed to close Redis connection:", err)
	}

	return keys, nil
}

func main() {
	keys, err := GetAllKeysFromRedis()
	if err != nil {
		log.Println("Failed to get keys from Redis:", err)
		return
	}

	// Выводим все ключи
	for _, key := range keys {
		fmt.Println(key)
	}
}
