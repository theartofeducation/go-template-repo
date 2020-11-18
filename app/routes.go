package app

import "net/http"

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
