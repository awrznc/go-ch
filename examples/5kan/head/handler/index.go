package handler

import (
	"net/http"
	"time"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(`<p>5kan</p>`))
	writer.Header().Set(`Content-Type`, `text/plain`)
	writer.Header().Set(`Last-Modified`, time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST"))
}
