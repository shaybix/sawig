package controller

import (
	"html/template"
	"log"
	"net/http"

	"github.com/shaybix/sawig/bootstrap"
)

// TODO: Need to load and parse the template filename, and return an error
//
//
// Please look at PagesController to get a better idea as to what is the
// desire result.
//

// Controller ...
type Controller struct {
	bootstrap.Controller
	Page
}

// Templates ...
type Templates struct {
	FileName string
	FilePath string
}

//ServeView ...
func (tpls *Templates) ServeView(res http.ResponseWriter, tpl *Templates) {

	// Initialise a new template, and parse the files
	t := template.New(tpl.FileName)
	template, err := t.ParseFiles(tpl.FilePath)

	// Process any errors
	if err != nil {

		log.Println(err)
	}

	// Takes as argument a writer and an interfacee of 'data' to be passed
	err = template.Execute(res, nil)

	// Process any errors
	if err != nil {
		log.Println(err)
	}

}
