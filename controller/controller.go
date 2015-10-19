package controller

import (
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
func (t *Templates) ServeView(file string) {}
