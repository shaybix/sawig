package controller

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func HomePage(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	fmt.Fprint(res, templates)
}
