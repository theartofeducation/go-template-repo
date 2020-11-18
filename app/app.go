package app

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

// App will hold all the dependencies the application needs.
type App struct {
	db     interface{}
	router *mux.Router
	log    *logrus.Logger
}

// NewApp creates a new instance of the App, registers the routes, and returns the instance.
func NewApp(router *mux.Router, log *logrus.Logger) App {
	app := App{
		router: router,
		log:    log,
	}

	app.routes()

	return app
}

// StartServer starts the HTTP server on the specified port. Any errors will be returned on the specified channel.
func (app App) StartServer(errorChan chan error, port string) {
	app.log.Infof("Starting server on port %s", port)
	errorChan <- http.ListenAndServe(":"+port, app.router)
}
