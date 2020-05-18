package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello from my web app!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch contact us at <a href=\"abc@gmail.com\">abc@gmail.com</a>")
	} else {
		fmt.Fprint(w, "Sorry, but this page could not be found")
	}

}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
