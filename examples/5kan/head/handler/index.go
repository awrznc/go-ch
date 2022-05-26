package handler

import (
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("<p>5kan</p>"))
}
