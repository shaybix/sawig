package controller

import (
	// "fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"html/template"
	"log"

)


type Page struct {}



func (p *Page) Index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {


	// Initialise a new template, and parse the files
	t := template.New("index.html")
	t, _ = t.ParseFiles("./views/index.html")

	// Takes as argument a writer and an interfacee of 'data' to be passed
	err := t.Execute(res, nil)

	// Process any errors
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}


}	

func About(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {


	// Initialise a new template, and parse the files
	t := template.New("about.html")
	t, _ = t.ParseFiles("./views/about.html")

	// Takes as argument a writer and an interfacee of 'data' to be passed
	err := t.Execute(res, nil)

	// Process any errors
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}


}
