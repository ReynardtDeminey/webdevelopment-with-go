package controllers

import (
	"fmt"
	"net/http"

	"github.com/ReynardtDeminey/webdevelopment-with-go/models"
	"github.com/ReynardtDeminey/webdevelopment-with-go/views"
)

type Users struct {
	NewView *views.View
	us      *models.UserService
}

func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
		us:      us,
	}
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	user := models.User{
		Name:  form.Name,
		Email: form.Email,
	}
	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "User is", user)

	// fmt.Fprintln(w, "Name is", form.Name)
	// fmt.Fprintln(w, "Email is", form.Email)
	// fmt.Fprintln(w, "Password is", form.Password)
}

type SignupForm struct {
	Name     string `schema: name`
	Email    string `schema: "email"`
	Password string `schema: "password"`
}