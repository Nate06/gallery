package main

import (
	"fmt"
	"gallery/controllers"
	"gallery/models"
	"gallery/views"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "yuhuiyuan"
	password = ""
	dbname   = "gallery"
)

var homeView *views.View
var contactView *views.View
var notFoundView *views.View

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s"+
		" password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.DestructiveReset()

	usersC := controllers.NewUsers(us)
	staticC := controllers.NewStatic()
	galleriesC := controllers.NewGalleries()

	r := mux.NewRouter()
	r.NotFoundHandler = staticC.NotFound

	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.HandleFunc("/galleries/new", galleriesC.New).Methods("GET")
	r.Handle("/login", usersC.LoginView).Methods("GET")
	r.HandleFunc("/login", usersC.Login).Methods("POST")
	http.ListenAndServe(":8080", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
