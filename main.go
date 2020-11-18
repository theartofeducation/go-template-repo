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

// server will hold all the dependencies the application needs.
type server struct {
	db     interface{} // The database connection
	router *mux.Router // The router
}

// routes holds all registered routes for the server.
func (s *server) routes() {
	s.router.HandleFunc("/", s.handleIndex()).Methods(http.MethodGet)
}

// handleIndex handles the "/" route.
func (s *server) handleIndex() http.HandlerFunc {
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
		log.Fatal("no port was specified")
	}

	router := mux.NewRouter()
	s := server{router: router}
	s.routes()

	errorChan := make(chan error, 2)

	go func() {
		log.Printf("Starting server at port %s", port)
		errorChan <- http.ListenAndServe(":"+port, s.router)
	}()

	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT)
		errorChan <- fmt.Errorf("%s", <-signalChan)
	}()

	fmt.Printf("Terminated %s", <-errorChan)
}
