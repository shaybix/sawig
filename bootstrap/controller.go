package bootstrap

import (
	"html/template"
	"log"
	"net/http"
)

// TODO: Need to load and parse the template filename, and return an error

// Templates ...
type Templates struct {
	FileName string
	FilePath string
}

// LoadTemplate - Loading the templates in the ./Views
func (tpls *Templates) LoadTemplate(FileName string) *template.Template {

	tpls.FileName = FileName
	tpls.FilePath = "views/" + tpls.FileName

	// Initialise a new template, and parse the files
	t := template.New(tpls.FileName)
	template, err := t.ParseFiles(tpls.FilePath)
	// Process any errors
	if err != nil {

		log.Println(err)
	}

	return template

}

//ServeView ...
func (tpls Templates) ServeView(view string, res http.ResponseWriter) {

	tpl := tpls.LoadTemplate(view)

	// Takes as argument a writer and an interfacee of 'data' to be passed
	err := tpl.Execute(res, nil)

	// Process any errors
	if err != nil {
		log.Println(err)
	}

}
