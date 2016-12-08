package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"site/page"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *page.Page) {
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, p)
}

func viewHandler(wr http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len(""):]
	pg, err := page.LoadPage(title)
	if err != nil {
		log.Fatal(err)
	}
	renderTemplate(wr, "logan", pg)
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
