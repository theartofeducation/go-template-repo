package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/getsentry/sentry-go"

	"github.com/theartofeducation/go-template-repo/app"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	log "github.com/sirupsen/logrus"
)

// TODO: Docker
// TODO: Logging
// TODO: Tests (black box)

// env variables
var (
	port string // The HTTP port the server will run on.
	dsn  string // The Sentry DSN.
)

func main() {
	loadEnvVariables()

	if strings.TrimSpace(dsn) != "" {
		loadSentry(dsn)
		defer sentry.Flush(time.Second * 2)
	}

	router := mux.NewRouter()
	a := app.NewApp(router)

	errorChan := make(chan error, 2)
	go a.StartServer(errorChan, port)
	go handleInterrupt(errorChan)

	err := <-errorChan
	sentry.CaptureMessage(err.Error())
	log.Errorln(err)
}

func loadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		log.Infoln("could not load env file")
	}

	port = os.Getenv("PORT")
	if strings.TrimSpace(port) == "" {
		log.Fatal("port was not specified")
	}

	dsn = os.Getenv("SENTRY_DSN")
	if strings.TrimSpace(dsn) == "" {
		log.Infoln("Sentry DSN not specified")
	}
}

func loadSentry(dsn string) {
	if err := sentry.Init(sentry.ClientOptions{Dsn: dsn}); err != nil {
		log.Errorln("failed to connect to Sentry:", err)
	} else {
		log.Infoln("connected to Sentry")
	}
}

func handleInterrupt(errorChan chan error) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)
	errorChan <- fmt.Errorf("%s", <-signalChan)
}
