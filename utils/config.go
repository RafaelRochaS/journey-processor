package utils

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var KAFKA_HOSTS string
var KAFKA_GROUP_ID string
var KAFKA_TOPIC string
var MIN_COMMIT_COUNT int

func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	KAFKA_HOSTS = os.Getenv("KAFKA_HOSTS")
	KAFKA_TOPIC = os.Getenv("KAFKA_TOPIC")
	KAFKA_GROUP_ID = os.Getenv("KAFKA_GROUP_ID")
	MIN_COMMIT_COUNT, err = strconv.Atoi(os.Getenv("MIN_COMMIT_COUNT"))

	if err != nil {
		MIN_COMMIT_COUNT = 5
	}
}
