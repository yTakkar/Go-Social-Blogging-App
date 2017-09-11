package main

import (
	R "Go-Social-Blogging-App/routes"
	"os"

	"github.com/joho/godotenv"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

func init() {
	godotenv.Load()
}

func main() {
	router := httprouter.New()

	router.GET("/", R.App)
	router.GET("/signup", R.Signup)
	router.GET("/login", R.Login)
	router.GET("/logout", R.Logout)
	router.GET("/404", R.NotFound)
	router.GET("/welcome", R.Welcome)
	router.GET("/create-blog", R.CreateBlog)
	router.GET("/blog/:blogID", R.ViewBlog)
	router.GET("/blog", R.Paramless)
	router.GET("/edit-blog/:blogID", R.EditBlog)
	router.GET("/edit-blog", R.Paramless)
	router.GET("/profile/:user", R.Profile)
	router.GET("/profile", R.Paramless)

	router.POST("/user/signup", R.UserSignup)
	router.POST("/user/login", R.UserLogin)

	router.POST("/api/create-new-blog", R.CreateNewBlog)
	router.POST("/api/delete-blog", R.DeleteBlog)
	router.POST("/api/edit-blog", R.EditYourBlog)

	server := negroni.Classic()
	server.UseHandler(router)

	server.Run(os.Getenv("PORT"))

}
