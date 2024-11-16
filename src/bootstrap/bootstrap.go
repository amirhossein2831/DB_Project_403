package bootstrap

import (
	"DB_Project/src/api/http"
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

	// Initialize Env variable
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("ENV Service: Failed to  loading .env file.    timestamp: %s", time.Now().String())
	}
	log.Printf("ENV Service: Env variable initial successfully.    timestamp: %s \n", time.Now().String())

	//Initialize http server
	go func() {
		log.Printf("HTTP Service: HTTP server initial successfully.    timestamp: %s \n", time.Now().String())

		err = http.Init()
		if err != nil {
			log.Fatalf("HTTP Service: Failed to Initialize. %v.    timestamp: %s \n", err, time.Now().String())
		}
	}()

	time.Sleep(50 * time.Millisecond)
	log.Printf("Application is now running.Press CTRL-C to exit.    timestamp: %s \n", time.Now().String())
	<-sc

	log.Printf("Application shutting down....    timestamp: %s \n", time.Now().String())
	err = http.ShutdownServer()
	if err != nil {
		log.Fatalf("HTTP Service: Failed to shutdown server. %v.    timestamp: %s \n", err, time.Now().String())
	}
	log.Printf("HTTP Service: server shutdwon sucessfully. %v.    timestamp: %s \n", err, time.Now().String())
	time.Sleep(1 * time.Second)
}
