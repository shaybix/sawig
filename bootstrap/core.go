package bootstrap


import (

    // "fmt"
    // "net/http"
    //"github.com/julienschmidt/httprouter"
     "log"
     "io/ioutil"
)

type Application struct {
	Configuration *Configuration

}



func (application *Application) Init(filename *string) {

    // instantiate the configuration
    application.Configuration = &Configuration{}
    err := application.Configuration.Load(*filename)
    if err != nil {
        log.Fatalln("Can't read configuration file: %s", err)
        panic(err)
    }

}



func (application *Application) LoadTemplates() []string {

    files, _ := ioutil.ReadDir(application.Configuration.TemplatePath)
    var templates []string

    for _, file := range files {
        templates = append(templates, file.Name())
    }

    return templates

}


func (application *Application) Route(controller interface{}, route string) interface{}{
    //
    // fn := func(c web.C, w http.ResponseWriter, r *http.Request) {
    // c.Env["Content-Type"] = "text/html"
    //
    //     methodValue := reflect.ValueOf(controller).MethodByName(route)
    //     methodInterface := methodValue.Interface()
    //     method := methodInterface.(func(c web.C, r *http.Request) (string, int))
    // 
    //     body, code := method(c, r)
    //
    //     if session, exists := c.Env["Session"]; exists {
    //         err := session.(*sessions.Session).Save(r, w)
    //         if err != nil {
    //             glog.Errorf("Can't save session: %v", err)
    //         }
    //     }
    //
    //     switch code {
    //     case http.StatusOK:
    //         if _, exists := c.Env["Content-Type"]; exists {
    //             w.Header().Set("Content-Type", c.Env["Content-Type"].(string))
    //         }
    //         io.WriteString(w, body)
    //     case http.StatusSeeOther, http.StatusFound:
    //         http.Redirect(w, r, body, code)
    //     default:
    //         w.WriteHeader(code)
    //         io.WriteString(w, body)
    //     }
    // }
    // return fn
}
