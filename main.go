package main

import (
	"fmt"
	"net/http"

	"github.com/ReynardtDeminey/webdevelopment-with-go/controllers"
	"github.com/ReynardtDeminey/webdevelopment-with-go/models"
	"github.com/gorilla/mux"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "photo_blog_dev"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// func faq(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type", "text/html")
// 	fmt.Fprint(w, "Here we answer all your questions")
// }

// func notFound(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type", "text/html")
// 	w.WriteHeader(http.StatusNotFound)
// 	fmt.Fprint(w, "<h1>This page was not found</h1>")
// }

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.AutoMigrate()
	// homeView = views.NewView("bootstrap", "views/home.gohtml")
	// contactView = views.NewView("bootstrap", "views/contact.gohtml")
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(us)

	r := mux.NewRouter()
	// r.NotFoundHandler = http.HandlerFunc(notFound)
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/faq", staticC.Faq).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.Handle("/login", usersC.LoginView).Methods("GET")
	r.HandleFunc("/login", usersC.Login).Methods("POST")
	// r.HandleFunc("/faq", faq)

	http.ListenAndServe(":3000", r)
}
