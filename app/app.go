package app

import (
	"io"
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

// Routes holds all registered Routes and universal middleware for the App.
func (app *App) Routes() {
	app.router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Content-Type", "application/json")

			next.ServeHTTP(writer, request)
		})
	})

	app.router.HandleFunc("/", app.handleIndex()).Methods(http.MethodGet)
}

func (app *App) handleIndex() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)

		_, _ = io.WriteString(writer, "Hello world!")
	}
}

// StartServer starts the HTTP server on the specified port. Any errors will be returned on the specified channel.
func (app App) StartServer(errorChan chan error, port string) {
	log.Printf("Starting server on port %s", port)
	errorChan <- http.ListenAndServe(":"+port, app.router)
}
