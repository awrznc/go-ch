package main

import (
	"net/http"

	bh "ch/examples/5kan/bbsmenu/handler"
	ih "ch/examples/5kan/index/handler"
)

func main() {
	http.HandleFunc("/", ih.Handler)
	http.HandleFunc("/bbsmenu.json", bh.Handler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
