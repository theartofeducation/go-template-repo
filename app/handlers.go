package app

import (
	"errors"
	"net/http"
)

func (app *App) handleIndex() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writeJSONResponse(writer, http.StatusOK, "Hello world!")
	}
}

func (app *App) handleTestError() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writeErrorJSONResponse(writer, http.StatusInternalServerError, errors.New("test error"))
	}
}
