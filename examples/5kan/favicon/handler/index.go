package handler

import (
	"net/http"
	"time"
	_ "embed"
)

var (
	//go:embed favicon.ico
	favicon []byte
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	writer.Write(favicon)
	writer.Header().Set(`Content-Type`, `image/x-icon`)
	writer.Header().Set(`Accept-Ranges`, `bytes`)
	writer.Header().Set(`Last-Modified`, time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST"))
}
