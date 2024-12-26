package bootstrap

import (
	"DB_Project/src/api/http"
	"DB_Project/src/database/connection/pgx"
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
		log.Fatalf("ENV Service: Failed to  loading .env file. %v.    timestamp: %s", err, time.Now().String())
	}
	log.Printf("ENV Service: Env variable initial successfully.    timestamp: %s \n", time.Now().String())

	// Initialize pgx:Database
	err = pgx.Init()
	if err != nil {
		log.Fatalf("Database:pgx Service: Failed to Initialize: %v.    timestamp: %s", err, time.Now().String())
	}
	log.Printf("Database:pgx Service: Database  initial successfully.    timestamp: %s \n", time.Now().String())

	//Initialize http server
	go func() {
		log.Printf("HTTP Service: HTTP server initial successfully.    timestamp: %s \n", time.Now().String())

		err = http.Init()
		if err != nil {
			log.Fatalf("HTTP Service: Failed to Initialize. %v.    timestamp: %s \n", err, time.Now().String())
		}
	}()

	// app started ...
	time.Sleep(50 * time.Millisecond)
	log.Printf("Application is now running.Press CTRL-C to exit.    timestamp: %s \n", time.Now().String())
	<-sc

	// Shutting down application
	log.Printf("Application shutting down....    timestamp: %s \n", time.Now().String())

	// Close Database:pgx
	err = pgx.Close()
	if err != nil {
		log.Fatalf("Databasde Service: Failed to close database. %v.    timestamp: %s \n", err, time.Now().String())
	}
	log.Printf("Databasde Service: database close sucessfully.    timestamp: %s \n", time.Now().String())

	// shutdown httpserver
	err = http.ShutdownServer()
	if err != nil {
		log.Fatalf("HTTP Service: Failed to shutdown server. %v.    timestamp: %s \n", err, time.Now().String())
	}
	log.Printf("HTTP Service: server shutdwon sucessfully.    timestamp: %s \n", time.Now().String())

	// wait for every thing down
	time.Sleep(1 * time.Second)
}
