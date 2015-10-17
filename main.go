package main


import (
    "flag"
    "github.com/shaybix/sawig/bootstrap"
    "github.com/julienschmidt/httprouter"
    "github.com/shaybix/sawig/controller"
    "net/http"
    "fmt"
    // "io/ioutil"
)

var app     = &bootstrap.Application{}

func init() {

    filename := flag.String("config", "config.json", "Path to configuration file")

    flag.Parse()
    app.Init(filename)
    // app.LoadTemplates()

}

// func IndexHandler(res http.ResponseWriter, req *http.Request ) {
//
//
// }



func main() {


    router := httprouter.New()
    router.GET("/", controller.HomePage)


    // Print on console the server is running and listen for connection
    fmt.Println("Running Server at http://localhost:8000")
    http.ListenAndServe(":8000", nil)




}
