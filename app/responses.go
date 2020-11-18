package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func writeJSONResponse(writer http.ResponseWriter, statusCode int, v interface{}) {
	writer.WriteHeader(statusCode)

	if err := json.NewEncoder(writer).Encode(v); err != nil {
		log.Println(err)
	}
}

func writeErrorJSONResponse(writer http.ResponseWriter, statusCode int, err error) {
	writer.WriteHeader(statusCode)

	body := fmt.Sprintf(`{"error": %s"`, err.Error())
	if _, err := io.WriteString(writer, body); err != nil {
		log.Println(err)
	}
}