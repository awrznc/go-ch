package handler

import (
	_ "embed"
	"html/template"
	"net/http"
)

type Template struct {
	Title string
	Body  template.HTML
}

var (
	//go:embed index.html
	indexHtml string
)

func Handler(writer http.ResponseWriter, request *http.Request) {

	html, err := template.New("").Parse(indexHtml)
	if err != nil {
		panic(err)
	}

	template := Template{`BBS menu`, template.HTML(`<p><a href="bbsmenu.json">bbsmenu.json</a></p>`)}
	err = html.Execute(writer, template)
	if err != nil {
		panic(err)
	}
}
