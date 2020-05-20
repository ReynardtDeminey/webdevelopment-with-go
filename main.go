package main

import (
	"net/http"

	"github.com/ReynardtDeminey/webdevelopment-with-go/controllers"
	"github.com/gorilla/mux"
)

// var (
// 	homeView    *views.View
// 	contactView *views.View
// )

// func home(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type", "text/html")
// 	must(homeView.Render(w, nil))
// 	// err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// }

// func contact(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type", "text/html")
// 	must(contactView.Render(w, nil))
// 	// err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// }

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
	// homeView = views.NewView("bootstrap", "views/home.gohtml")
	// contactView = views.NewView("bootstrap", "views/contact.gohtml")
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers()

	r := mux.NewRouter()
	// r.NotFoundHandler = http.HandlerFunc(notFound)
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/faq", staticC.Faq).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	// r.HandleFunc("/faq", faq)

	http.ListenAndServe(":3000", r)
}
