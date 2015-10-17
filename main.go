package main


import (
    "flag"
    "github.com/shaybix/sawig/bootstrap"
    "net/http"
    "fmt"
    // "io/ioutil"
)

var app     = &bootstrap.Application{}
var route   = &bootstrap.Route{}

func init() {

    filename := flag.String("config", "config.json", "Path to configuration file")

    flag.Parse()
    app.Init(filename)
    route.Init()
    // app.LoadTemplates()

}

// func IndexHandler(res http.ResponseWriter, req *http.Request ) {
//
//
// }



func main() {

    // http.HandleFunc("/", IndexHandler)
    // fs := http.FileServer(http.Dir("views"))
    // http.Handle("/", fs)

    // templates := app.LoadTemplates()


    //
    // route.Register("/", "index.html", )
    // route.Register("about", "about.html")
    //
    // for _, each := range route.All() {
    //
    // }


    route.Get("/", IndexController)


    // Print on console the server is running and listen for connection
    fmt.Println("Running Server at http://localhost:8000")
    http.ListenAndServe(":8000", nil)




}
