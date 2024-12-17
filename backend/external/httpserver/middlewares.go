package httpserver

import (
	"fmt"
	"log"
	"net/http"
)

func answerError(w http.ResponseWriter, statusCode int, format string, v ...any) {
	err := fmt.Errorf(format, v...)
	log.Print(err)
	w.WriteHeader(statusCode)
	w.Write([]byte(err.Error()))
}
