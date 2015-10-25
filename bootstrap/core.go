package bootstrap

import (
	"io/ioutil"
	"log"
)

// PublicFolders ...
type PublicFolders struct {
	Dirs map[string]string
}

// Application ...
type Application struct {
	Configuration *Configuration
}

// Init Loading - the config file
func (application *Application) Init(filename *string) {

	// instantiate the configuration
	config := &Configuration{}
	err := config.Load(*filename)

	if err != nil {
		log.Fatalln("Can't read configuration file: %s", err)
		panic(err)
	}

}

// LoadPublicFolder ...
func (application *Application) LoadPublicFolder() *PublicFolders {

	config := &Configuration{}

	files, _ := ioutil.ReadDir(config.PublicPath)

	pf := new(PublicFolders)

	pf.Dirs = make(map[string]string)

	for _, dir := range files {

		pf.Dirs[string(dir.Name())] = config.PublicPath + string(dir.Name())

	}

	return pf
}
