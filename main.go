package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// TODO: Sentry
// TODO: Docker

// app will hold all the dependencies the application needs.
type app struct {
	db     interface{}
	router *mux.Router
}

// routes holds all registered routes for the app.
func (a *app) routes() {
	a.router.HandleFunc("/", a.handleIndex()).Methods(http.MethodGet)
}

func (a *app) handleIndex() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")

		_, _ = io.WriteString(writer, "Hello world!")
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("could not load env")
	}

	port := os.Getenv("PORT")
	if strings.TrimSpace(port) == "" {
		log.Fatal("port was not specified")
	}

	router := mux.NewRouter()
	a := app{router: router}
	a.routes()

	errorChan := make(chan error, 2)
	go startServer(a, errorChan, port)
	go handleInterrupt(errorChan)

	fmt.Printf("Terminated %s", <-errorChan)
}

func startServer(a app, errorChan chan error, port string) {
	log.Printf("Starting server on port %s", port)
	errorChan <- http.ListenAndServe(":"+port, a.router)
}

func handleInterrupt(errorChan chan error) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)
	errorChan <- fmt.Errorf("%s", <-signalChan)
}
