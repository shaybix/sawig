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

	files, _ := ioutil.ReadDir(application.Configuration.TemplatePath)
	var templates []string

	for _, file := range files {
		templates = append(templates, file.Name())
	}

	return templates

}


func (application *Application) Controller(templates []string) {

	application.Controller = &Controller{}
	err := application.Controller.Init(templates)

	if err != nil {

		log.Fatalln(err)
		panic(err)
	}
}
