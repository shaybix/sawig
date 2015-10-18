package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/shaybix/sawig/bootstrap"
	"github.com/shaybix/sawig/controller"
)

var app = &bootstrap.Application{}


func init() {

	// Load the config.json file
	filename := flag.String("config", "config.json", "Path to configuration file")

	// Parse the arguments
	flag.Parse()

	// initialising the config.json
	app.Init(filename)

}


func main() {

	// instantiate the router
	router := httprouter.New()


	// The application's routes
	router.GET("/", controller.Index)



	// Print on console the server is running and listen for connection
	fmt.Println("Running Server at http://localhost:8000")
	http.ListenAndServe(":8000", router)


}
