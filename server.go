package main

import (
	"flag"
	"fmt"

	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/shaybix/sawig/bootstrap"
	"github.com/shaybix/sawig/controller"
	"github.com/ttacon/chalk"
)

var app = &bootstrap.Application{}
var ctrl = &controller.ControllersInit{}

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
	router.GET("/", ctrl.Pages.Index)
	router.GET("/about", ctrl.Pages.About)
	router.GET("/contact", ctrl.Pages.Contact)

	// Print on console the server is running and listen for connection
	fmt.Println("\nRunning Server at", chalk.Green, "http://localhost:8000", chalk.ResetColor)

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatalln(chalk.Red, err)

	}

}
