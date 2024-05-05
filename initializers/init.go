package initializers

import (
	"anya/db/Kafka"
	"anya/db/redis"
	"anya/enums"
	"anya/utilities"
	"github.com/joho/godotenv"
)

var err error

func Init(env_file string) {
	load_env(env_file)
	Kafka.Connect()
	redis.ConnectRedis()
}

func load_env(fl string) {
	err = godotenv.Load(fl)
	utilities.LogError(err, enums.EnvLoadFailed)
}
