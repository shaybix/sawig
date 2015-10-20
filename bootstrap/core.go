package bootstrap

import (
	"log"
)

// Application ...
type Application struct {
	Configuration *Configuration
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
