package main

import "net/http"
import "html/template"
import "site/page"
import "log"

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
