package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// server will hold all the dependencies the application needs
type server struct {
	db     interface{} // The database connection
	router *mux.Router // The router
}

// routes holds all registered routes for the server.
func (s *server) routes() {
	s.router.HandleFunc("/", s.handleIndex())
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprint(writer, "Hello world!")
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("could not load env")
	}

	router := mux.NewRouter()

	s := server{
		router: router,
	}

	port := os.Getenv("PORT")
	if strings.TrimSpace(port) == "" {
		log.Fatal("no port was specified")
	}

	errorChan := make(chan error, 2)

	go func() {
		log.Println("Starting server at port:", port)
		errorChan <- http.ListenAndServe(":"+port, s.router)
	}()

	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT)
		errorChan <- fmt.Errorf("%s", <-signalChan)
	}()

	fmt.Printf("Terminated %s", <-errorChan)
}
