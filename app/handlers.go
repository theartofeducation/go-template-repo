package app

import (
	"net/http"
)

func (app *App) handleIndex() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writeJSONResponse(writer, http.StatusOK, "Hello world!")
	}
}
