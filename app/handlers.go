package app

import (
	"io"
	"net/http"
)

func (app *App) handleIndex() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)

		_, _ = io.WriteString(writer, "Hello world!")
	}
}
