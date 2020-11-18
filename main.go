package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/theartofeducation/go-template-repo/app"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// TODO: Sentry
// TODO: Docker
// TODO: Logging

// env variables
var (
	port string
)

func main() {
	loadEnvVariables()

	router := mux.NewRouter()
	a := app.NewApp(router)

	errorChan := make(chan error, 2)
	go a.StartServer(errorChan, port)
	go handleInterrupt(errorChan)

	fmt.Printf("Terminated %s", <-errorChan)
}

func loadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		log.Println("could not load env")
	}

	port = os.Getenv("PORT")
	if strings.TrimSpace(port) == "" {
		log.Fatal("port was not specified")
	}
}

func handleInterrupt(errorChan chan error) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)
	errorChan <- fmt.Errorf("%s", <-signalChan)
}
