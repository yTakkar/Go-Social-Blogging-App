package routes

import (
	M "Go-Social-Blogging-App/models"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

// Signup route
func Signup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	notLoggedIn(w, r)
	renderTemplates(w, "signup", &Page{"Signup to blog", nil, nil, nil, nil})
}

// Login route
func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	notLoggedIn(w, r)
	renderTemplates(w, "login", &Page{"Login to blog", nil, nil, nil, nil})
}

// Logout route
func Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	loggedIn(w, "", r)

	session := M.GetSession(r)
	delete(session.Values, "id")
	delete(session.Values, "username")
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusFound)
}

// UserSignup route for registering
func UserSignup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resp := make(map[string]interface{})

	username := strings.TrimSpace(r.PostFormValue("username"))
	email := strings.TrimSpace(r.PostFormValue("email"))
	password := strings.TrimSpace(r.PostFormValue("password"))
	passwordAgain := strings.TrimSpace(r.PostFormValue("password_again"))

	mailErr := checkmail.ValidateFormat(email)

	db := M.DB()

	var (
		userCount  int
		emailCount int
	)

	db.QueryRow("SELECT COUNT(id) AS userCount FROM users WHERE username=?", username).Scan(&userCount)
	db.QueryRow("SELECT COUNT(id) AS emailCount FROM users WHERE email=?", email).Scan(&emailCount)

	if username == "" || email == "" || password == "" || passwordAgain == "" {
		resp["mssg"] = "Some values are missing!"
	} else if len(username) < 4 || len(username) > 32 {
		resp["mssg"] = "Username should be between 4 and 32"
	} else if mailErr != nil {
		resp["mssg"] = "Invalid Format!"
	} else if password != passwordAgain {
		resp["mssg"] = "Passwords don't match"
	} else if userCount > 0 {
		resp["mssg"] = "Username already exists!"
	} else if emailCount > 0 {
		resp["mssg"] = "Email already exists!"
	} else {

		hash, hashErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if hashErr != nil {
			log.Fatal(hashErr)
		}

		rs, iErr := db.Exec(
			"INSERT INTO users(username, email, password, bio, joined) VALUES(?, ?, ?, ?, ?)",
			username,
			email,
			hash,
			"",
			M.MakeTimestamp(),
		)

		if iErr != nil {
			log.Fatal(iErr)
		}

		insertID, _ := rs.LastInsertId()
		dir, _ := os.Getwd()
		userPath := dir + "/public/users/" + strconv.FormatInt(insertID, 10)

		mkErr := os.Mkdir(userPath, 0655)
		if mkErr != nil {
			log.Fatal(mkErr)
		}

		cErr := os.Link(dir+"/public/images/golang-color-icon2.png", userPath+"/avatar.png")
		if cErr != nil {
			log.Fatal(cErr)
		}

		session := M.GetSession(r)
		session.Values["id"] = insertID
		session.Values["username"] = username
		session.Save(r, w)

		resp["success"] = true
		resp["mssg"] = "Hello, " + username

	}

	M.JSON(w, r, resp)
}

// UserLogin route
func UserLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resp := make(map[string]interface{})

	rusername := strings.TrimSpace(r.PostFormValue("username"))
	rpassword := strings.TrimSpace(r.PostFormValue("password"))

	db := M.DB()
	var (
		userCount int
		id        int
		username  string
		password  string
	)

	db.QueryRow("SELECT COUNT(id) AS userCount, id, username, password FROM users WHERE username=?", rusername).Scan(&userCount, &id, &username, &password)

	encErr := bcrypt.CompareHashAndPassword([]byte(password), []byte(rpassword))

	if rusername == "" || rpassword == "" {
		resp["mssg"] = "Some values are missing!"
	} else if userCount == 0 {
		resp["mssg"] = "Invalid username!"
	} else if encErr != nil {
		resp["mssg"] = "Invalid password!"
	} else {

		session := M.GetSession(r)
		session.Values["id"] = id
		session.Values["username"] = username
		session.Save(r, w)

		resp["mssg"] = "Hello, " + username
		resp["success"] = true

	}

	M.JSON(w, r, resp)
}
