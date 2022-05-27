package handler

import (
	"net/http"
	"time"
	"compress/gzip"
	"strings"
	"bytes"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	var result string
	contents := `<p>5kan</p>`
	if strings.Contains(request.Header.Get("Accept-Encoding"), "gzip") {
		result = gzipping(contents)
		writer.Header().Set("Content-Encoding", "gzip")
	} else {
		result = contents
	}
	writer.Write([]byte(result))
	writer.Header().Set(`Content-Type`, `text/plain`)
	writer.Header().Set(`Accept-Ranges`, `bytes`)
	writer.Header().Set(`Last-Modified`, time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST"))
}

func gzipping(target string) string {
	var buffer bytes.Buffer
	writer := gzip.NewWriter(&buffer)
	writer.Write([]byte(target))
	writer.Close()
	return string(buffer.Bytes())
}
