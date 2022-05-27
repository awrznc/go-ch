package handler

import (
	"fmt"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("%#v", request)

	writer.Write([]byte("not found."))
}
