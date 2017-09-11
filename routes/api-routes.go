package routes

import (
	M "Go-Social-Blogging-App/models"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// CreateNewBlog route
func CreateNewBlog(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res := make(map[string]interface{})

	db := M.DB()

	title := strings.TrimSpace(r.PostFormValue("title"))
	content := strings.TrimSpace(r.PostFormValue("content"))

	_, err := db.Exec("INSERT INTO blogs(title, content) VALUES(?, ?)", title, content)
	if err != nil {
		log.Fatal(err)
	}

	res["mssg"] = "Blog created!!"
	M.JSON(w, r, res)
}

// DeleteBlog route
func DeleteBlog(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res := make(map[string]interface{})
	blog, _ := strconv.Atoi(r.PostFormValue("blog"))

	db := M.DB()
	db.Exec("DELETE FROM blogs WHERE blogID=?", blog)

	res["mssg"] = "Blog Deleted!!"
	M.JSON(w, r, res)
}

// EditYourBlog route
func EditYourBlog(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res := make(map[string]interface{})

	blogID := r.PostFormValue("blogID")
	title := r.PostFormValue("title")
	content := r.PostFormValue("content")

	db := M.DB()
	db.Exec("UPDATE blogs SET title=?, content=? WHERE blogID=?", title, content, blogID)

	res["mssg"] = "Blog Edited!!"
	M.JSON(w, r, res)
}
