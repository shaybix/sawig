package controller

import (
	// "fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Page ...
type Page struct{}

// Consider exporting this code below to another function, and simply
// execute to serve the template, like this:
//
// ServeView("file.html")
//
//
// TODO: refactor all of this code
//
//

// Index ...
func (p *Page) Index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	// // Initialise a new template, and parse the files
	// t := template.New("index.html")
	// t, _ = t.ParseFiles("./views/index.html")
	//
	// // Takes as argument a writer and an interfacee of 'data' to be passed
	// err := t.Execute(res, nil)
	//
	// // Process any errors
	// if err != nil {
	// 	log.Fatalln(err)
	// 	panic(err)
	// }

	tpl := new(Templates)
	tpl.FileName = "index.html"
	tpl.FilePath = "./views/index.html"
	tpl.ServeView(res, tpl)

}

// About ...
func (p *Page) About(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	// // Initialise a new template, and parse the files
	// t := template.New("about.html")
	// t, _ = t.ParseFiles("./views/about.html")
	//
	// // Takes as argument a writer and an interfacee of 'data' to be passed
	// err := t.Execute(res, nil)
	//
	// // Process any errors
	// if err != nil {
	// 	log.Fatalln(err)
	// 	panic(err)
	// }

	t := new(Templates)
	t.FileName = "about.html"
	t.FilePath = "./views/about.html"
	t.ServeView(res, t)

}
