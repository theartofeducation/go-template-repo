package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// App will hold all the dependencies the application needs.
type App struct {
	db     interface{}
	router *mux.Router
}

// NewApp creates and returns a new App instance.
func NewApp(router *mux.Router) App {
	app := App{
		router: router,
	}

	return app
}

// StartServer starts the HTTP server on the specified port. Any errors will be returned on the specified channel.
func (app App) StartServer(errorChan chan error, port string) {
	log.Printf("Starting server on port %s", port)
	errorChan <- http.ListenAndServe(":"+port, app.router)
}
