package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

type Error struct {
	Code   int
	Reason string
}

const (
	templatesPrefix = "templates"
)

func LoadPage(name string) (*Page, error) {
	filename := templatesPrefix + "/" + name + ".html"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{
		Title: name,
		Body:  body,
	}, nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file := r.URL.Path[1:]
		p, err := LoadPage(file)
		if err != nil {
			t, _ := template.ParseFiles(templatesPrefix + "/error.html")
			errObj := &Error{
				Reason: err.Error(),
				Code:   404,
			}
			t.Execute(w, errObj)
		} else {
			fmt.Fprintf(w, "%s", p.Body)
		}
	})

	http.ListenAndServe(":8080", nil)
}
