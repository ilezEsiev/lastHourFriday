package DB

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

// Redis database configuration
var Rdb *redis.Client

func InitRedis() *redis.Client {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println("Подключение  Redis выполнено")
	return Rdb
}
