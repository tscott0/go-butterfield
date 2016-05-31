package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"regexp"
)

type Page struct {
	Title       string
	Description string
}

var Diet []string = []string{
	"Pints o’ cream",
	"Potato grids",
	"Large macs",
	"Chocolate quail’s eggs",
	"Fluffy ruffs",
	"Pasta pillows",
	"Mcfortune cookies",
	"Egg ‘n’ ham slabs",
	"Pork cylinders",
	"Artificial bacon (Facon ™)",
	"Sandwich casserole",
	"Garlic pudding",
	"Hoisin crispy owl",
	"Bonbonbonbons",
	"Discount foie gras",
	"During-dinner mints",
	"Quiches lorraine",
}

var templates = template.Must(template.ParseFiles("view.html"))
var validPath = regexp.MustCompile("^/(view|random)$")

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	title := Diet[rand.Intn(len(Diet))]
	description := title + ", part of the Butterfield diet."
	p := &Page{Title: title, Description: description}

	err := templates.ExecuteTemplate(w, "view.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}

func main() {
	http.HandleFunc("/random", makeHandler(randomHandler))

	http.ListenAndServe(":8080", nil)
}
