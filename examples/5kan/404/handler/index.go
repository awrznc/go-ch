package handler

import (
	// "fmt"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	// panic(fmt.Sprintf("%#v", request))

	writer.Write([]byte("not found."))
}
