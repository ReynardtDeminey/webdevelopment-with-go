package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to the home page!")
}

func About(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to the about page!")
}

func main() {
	r := httprouter.New()
	r.GET("/", Home)
	r.GET("/about", About)

	log.Fatal(http.ListenAndServe(":8080", r))
}
