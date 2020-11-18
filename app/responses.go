package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func writeJSONResponse(writer http.ResponseWriter, statusCode int, v interface{}) {
	writer.WriteHeader(statusCode)

	if err := json.NewEncoder(writer).Encode(v); err != nil {
		log.Errorln(err)
	}
}

func writeErrorJSONResponse(writer http.ResponseWriter, statusCode int, err error) {
	writer.WriteHeader(statusCode)

	body := fmt.Sprintf(`{"error": %q}`, err.Error())
	if _, err := io.WriteString(writer, body); err != nil {
		log.Errorln(err)
	}
}