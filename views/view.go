package views

import "html/template"

// NewView parses the view files
func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml")

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
	}
}

// View contains all the information of our views
type View struct {
	Template *template.Template
}
