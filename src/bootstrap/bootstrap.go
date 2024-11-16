package bootstrap

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Init() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("ENV Service: Failed to  loading .env file. timestamp: %s", time.Now().String())
	}
	log.Printf("ENV Service: Env variable initial successfully. timestamp: %s \n", time.Now().String())

	time.Sleep(50 * time.Millisecond)
	log.Printf("Application is now running.Press CTRL-C to exit. timestamp: %s \n", time.Now().String())
	<-sc

	log.Printf("Application shutting down.... timestamp: %s \n", time.Now().String())
	time.Sleep(1 * time.Second)
}
