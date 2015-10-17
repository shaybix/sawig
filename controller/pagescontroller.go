package controller


import (
    "github.com/julienschmidt/httprouter"
    "fmt"
    "net/http"
)






func HomePage(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

    fmt.Fprint(res, "Hit The HomePage")
}
