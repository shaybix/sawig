package bootstrap

import (

	"io/ioutil"
	"log"
)


type Application struct {

	Configuration *Configuration
	Controller *Controller
}



func (application *Application) Init(filename *string) {

	// instantiate the configuration
	application.Configuration = &Configuration{}
	err := application.Configuration.Load(*filename)
	if err != nil {
		log.Fatalln("Can't read configuration file: %s", err)
		panic(err)
	}

}



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


func (application *Application) LoadControllers(templates []string) {

	application.Controller = &Controller{}

	//application.Controller.GetTemplates(templates)

}
