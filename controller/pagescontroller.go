package controller

import (
	"github.com/julienschmidt/httprouter"
	"github.com/shaybix/sawig/bootstrap"

	"net/http"
)

// Pages ...
type Pages struct {
	bootstrap.Templates
}

// Index ...
func (p *Pages) Index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	p.ServeView("index.html", res)
}

// About ...
func (p *Pages) About(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	p.ServeView("about.html", res)
}

// Contact ...
func (p *Pages) Contact(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	p.ServeView("contact.html", res)

}
