package controller

import (
	// "fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"html/template"
	"log"

)



func Index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	t := template.New("index.html")
	t, _ = t.ParseFiles("./views/index.html")
	err := t.Execute(res, nil)

	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
	// Initialise a new template, and then execute the template
	// t := template.New()
	//
	// Takes as argument a writer and an interfacee of 'data'
	// t.Execute()

}
