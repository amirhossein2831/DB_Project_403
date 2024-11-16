package bootstrap

import (
	"github.com/joho/godotenv"
	"log"
	"time"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("ENV Service: Failed to  loading .env file. timestamp: %s", time.Now().String())
	}
	log.Printf("ENV Service: Env variable initial successfully. timestamp: %s \n", time.Now().String())

}
