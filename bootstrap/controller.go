package bootstrap


import (

    // "html/template"
    "fmt"
)


type Controller struct {}




func (controller *Controller) GetTemplate(templates []string) {

    app := &Application{}


    var lists []string
    lists = app.LoadTemplates()
    fmt.Println(lists)
}
