package routes

import (
	M "Go-Social-Blogging-App/models"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// NotFound route
func NotFound(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	id, username := M.AllSessions(r)
	renderTemplates(w, "404", &Page{"Oops!! Error • Blog App", id, username, nil, nil})
}

// Paramless route
func Paramless(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Redirect(w, r, "/404", http.StatusFound)
}

// Welcome route
func Welcome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	notLoggedIn(w, r)
	renderTemplates(w, "welcome", &Page{"Welcome to Blog App!", nil, nil, nil, nil})
}

// App route
func App(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	loggedIn(w, "/welcome", r)

	id, username := M.AllSessions(r)

	db := M.DB()
	rows, _ := db.Query("SELECT * FROM blogs ORDER BY blogID DESC")
	blogs := []interface{}{}

	for rows.Next() {
		var (
			blogID  int
			title   string
			content string
		)
		rows.Scan(&blogID, &title, &content)
		blogs = append(blogs, Blog{blogID, title, content})
	}

	renderTemplates(w, "blogs", &Page{"Go Blog App", id, username, blogs, nil})
}

// CreateBlog route
func CreateBlog(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	loggedIn(w, "", r)

	id, username := M.AllSessions(r)
	renderTemplates(w, "create-blog", &Page{"Create A New Blog", id, username, nil, nil})
}

// ViewBlog route
func ViewBlog(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	loggedIn(w, "", r)

	id, username := M.AllSessions(r)
	blog := p.ByName("blogID")

	db := M.DB()
	var (
		blogID  int
		title   string
		content string
	)
	db.QueryRow("SELECT * FROM blogs WHERE blogID=?", blog).Scan(&blogID, &title, &content)

	invalid(w, blogID, r)
	renderTemplates(w, "blog", &Page{"View Blog", id, username, nil, &Blog{blogID, title, content}})
}

// EditBlog route
func EditBlog(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	loggedIn(w, "", r)

	id, username := M.AllSessions(r)
	blog := p.ByName("blogID")

	db := M.DB()
	var (
		blogID  int
		title   string
		content string
	)
	db.QueryRow("SELECT * FROM blogs WHERE blogID=?", blog).Scan(&blogID, &title, &content)
	fmt.Println(blogID, title, content)

	invalid(w, blogID, r)
	renderTemplates(w, "edit-blog", &Page{"Edit Blog", id, username, nil, &Blog{blogID, title, content}})
}

// Profile route
func Profile(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	loggedIn(w, "", r)

	user := p.ByName("user")
	db := M.DB()

	var (
		count    int
		id       int
		username string
	)
	db.QueryRow("SELECT COUNT(id), id, username AS count FROM users WHERE id=?", user).Scan(&count, &id, &username)

	invalid(w, count, r)
	renderTemplates(w, "profile", &Page{"Profile Page", id, username, nil, nil})
}
