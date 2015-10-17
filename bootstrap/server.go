package bootstrap


import (

    "fmt"
    "net/http"
    //"github.com/julienschmidt/httprouter"
    "log"

)



func Init() {

    fmt.Println("Running Server at http://localhost:8000")

    err := http.ListenAndServe(":8000", nil)
    if err != nil {

        log.Println(err)
    }


}

func LoadTemplates() {}

func ConnectToDatabase() {}
