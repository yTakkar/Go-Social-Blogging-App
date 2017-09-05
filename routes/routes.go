package routes

import (
	M "Go-Blog-App/models"
	"html/template"
	"log"
	"net/http"
)

// Page struct
type Page struct {
	Title    string
	UserID   interface{}
	Username interface{}
	Blogs    interface{}
	Blog     interface{}
}

// Blog struct
type Blog struct {
	BlogID  int
	Title   string
	Content string
}

var templates = template.Must(template.ParseGlob("views/*.html"))

func renderTemplates(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, "Oops! Page not found!!", http.StatusInternalServerError)
		log.Fatal(err)
	}
}

func loggedIn(w http.ResponseWriter, urlRedirect string, r *http.Request) {
	var URL string
	if urlRedirect == "" {
		URL = "/login"
	} else {
		URL = urlRedirect
	}
	id, _ := M.AllSessions(r)
	if id == nil {
		http.Redirect(w, r, URL, http.StatusFound)
	}
}

func notLoggedIn(w http.ResponseWriter, r *http.Request) {
	id, _ := M.AllSessions(r)
	if id != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func invalid(w http.ResponseWriter, what int, r *http.Request) {
	if what == 0 {
		http.Redirect(w, r, "/404", http.StatusFound)
	}
}
