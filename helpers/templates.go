package helpers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

// Helper ...
type Helper struct {
}

// Templates ...
type Templates struct {
	FileName string
	FilePath string
}

// LoadTemplates - Loading the templates int he ./Views folder and returning
// a struct with FileName and FilePath
func (tpls *Templates) LoadTemplates(req *http.Request) *Templates {

	tpl := new(Templates)

	if req.URL.Path == "/" {
		tpl.FileName = "index.html"
	} else {

		tpl.FileName = func() string {
			path := strings.Join(strings.Split(req.URL.Path, "/")[1:], "-") + ".html"
			fmt.Println(path)
			return path

		}()
	}

	tpl.FilePath = "./views/" + tpl.FileName

	return tpl

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
