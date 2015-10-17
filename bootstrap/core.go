package bootstrap


import (

    // "fmt"
    // "net/http"
    // "github.com/shaybix/sawig/controller"
    "log"
    "io/ioutil"
)

type Application struct {
	Configuration *Configuration

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
