package handler

import (
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(`<p>5kan</p>`))
	writer.Header().Set(`Content-Type`, `text/plain`)
}
