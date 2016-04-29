package main

import "net/http"
import "html/template"
import "site/page"

func renderTemplate(w http.ResponseWriter, tmpl string, p *page.Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := page.LoadPage(title)
	renderTemplate(w, "view", p)
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
