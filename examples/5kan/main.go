package main

import (
	"net/http"
	"regexp"

	bh "ch/examples/5kan/bbsmenu/handler"
	dh "ch/examples/5kan/dat/handler"
	ih "ch/examples/5kan/index/handler"
	sh "ch/examples/5kan/subject/handler"
	hh "ch/examples/5kan/head/handler"
	nh "ch/examples/5kan/404/handler"
)

func handler(writer http.ResponseWriter, request *http.Request) {

	if request.URL.Path == "/bbsmenu.json" {
		bh.Handler(writer, request)
	}

	matches := regexp.MustCompile(`^/(.+?)/(.+.txt|[^/]+?.dat)$`).FindStringSubmatch(request.URL.Path)
	if len(matches) == 3 {
		if matches[2] == "subject.txt" {
			sh.Handler(writer, request)
		} else if matches[2] == "head.txt" {
			hh.Handler(writer, request)
		} else {
			dh.Handler(writer, request)
		}
	} else {
		if request.URL.Path == "/" {
			ih.Handler(writer, request)
		} else {
			nh.Handler(writer, request)
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
