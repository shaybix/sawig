package bootstrap

import (
	"io/ioutil"
	"log"
)

// Application ...
type Application struct {
	Configuration *Configuration
	Controller    *Controller
}

// Init Loading - the config file
func (application *Application) Init(filename *string) {

	// instantiate the configuration
	application.Configuration = &Configuration{}
	err := application.Configuration.Load(*filename)
	if err != nil {
		log.Fatalln("Can't read configuration file: %s", err)
		panic(err)
	}

}

// LoadTemplates - Loading the templates int he ./Views folder and returning
// a slice of the file name.
//
// Perhaps should return a map of route path as key
// and template value.
func (application *Application) LoadTemplates() []string {

	// Read the template directory specified in Application.Configuration{}
	files, _ := ioutil.ReadDir(application.Configuration.TemplatePath)

	// Instantiate the templates slice which will hold all the template filenames
	var templates []string
	for _, file := range files {
		templates = append(templates, file.Name())
	}

	// return a slice of template filenames
	return templates

}

// LoadControllers - Prepares and passes on the templates on to the Controller
// struct, which then makes it available for other methods of the Controller.
func (application *Application) LoadControllers(templates []string) {

	application.Controller = &Controller{}

	// application.Controller.GetTemplates(templates)

}
