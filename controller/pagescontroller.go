package controller

import (
	"github.com/julienschmidt/httprouter"
	"github.com/shaybix/sawig/helpers"
	"github.com/ttacon/chalk"

	"fmt"
	"net/http"
)

// Page ...
type Page struct {
	Templates *helpers.Templates
}

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

	// fmt.Println(chalk.Yellow, req.URL.Path, chalk.ResetColor)

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

	tpl := p.Templates.LoadTemplates(req)
	p.Templates.ServeView(res, tpl)

}

// About ...
func (p *Page) About(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	fmt.Println(chalk.Yellow, req.URL.Path, chalk.ResetColor)

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

	tpl := p.Templates.LoadTemplates(req)
	p.Templates.ServeView(res, tpl)

}

// Contact ...
func (p *Page) Contact(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	tpl := p.Templates.LoadTemplates(req)
	p.Templates.ServeView(res, tpl)

}
