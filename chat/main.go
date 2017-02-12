package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"sync"
)

//go:generate go-bindata -o templates.go ./templates

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		if b, err := Asset(path.Join("templates", t.filename)); err != nil {
			log.Fatal(err)
		} else {
			t.templ, _ =
				template.New("complex").Parse(string(b))
		}
	})
	if err := t.templ.Execute(w, nil); err != nil {
		log.Fatal("ServeHTTP:", err)
	}
}

func main() {
	http.Handle("/", &templateHandler{filename: "chat.html"})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
