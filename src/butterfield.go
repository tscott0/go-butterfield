package main

import (
	"fmt"
	"github.com/gorilla/mux"
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
	"20 Cheese Omlette",
	"Artificial bacon (Facon ™)",
	"Birthday Pie",
	"Bonbonbonbons",
	"Chocolate quail’s eggs",
	"Discount foie gras",
	"During-dinner mints",
	"Egg ‘n’ ham slabs",
	"Fluffy ruffs",
	"Garlic pudding",
	"Hoisin crispy owl",
	"Large macs",
	"Mcfortune cookies",
	"Pasta pillows",
	"Pints o’ cream",
	"Pork cylinders",
	"Potato grids",
	"Quiches lorraine",
	"Sandwich casserole",
}

var templates = template.Must(template.ParseGlob("templates/*.html"))
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
		fmt.Println(r.URL.Path)
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/random", makeHandler(randomHandler))

	http.ListenAndServe(":8080", r)
}
