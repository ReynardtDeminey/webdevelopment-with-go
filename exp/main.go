package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := struct {
		FirstName string
		LastName  string
		Age       int
		Hobbies   []string
	}{
		FirstName: "John",
		LastName:  "Smith",
		Age:       42,
		Hobbies:   []string{"Golf", "Fishing"},
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
