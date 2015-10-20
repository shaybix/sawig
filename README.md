# sawig [Simple Application Written In Go]


This application is an excercise allowing me to learn and explore the language called Golang (or simply put Go). The idea is not to delve too much in the design of the application but rather on getting things done and working. In the process I am forced to read the code and application of other peoples code and see how they have managed to implement a particular functionality.



======================================================================


#### TODO Tasks (Functionalities)


- [x] HTTP Server
- [x] HTTP Server that accepts incoming traffic
- [x] Read from a json configuration file and parse settings
- [x] Router handles requests then passes it over to a controller
- [x] Serve HTML Files in Views folder
- [x] Load Templates properly
- [ ] Load assets in public folder
- [ ] Pass on data to the Templates
- [ ] Load templates dynamically
- [ ] Connect to the Database
- [ ] Load Models
- [ ] Explore GoRoutines



======================================================================

TODO: Currently the views get served from the controller in this fashion. However
ideally we want to be able to load all of the templates, and their file paths
into a map of some kind, and using convention map a url path to a template.

``` go

func (p *Page) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {


tpl := p.Templates.LoadTemplates(req)
p.Templates.ServeView(res, tpl)

}

```

After many trial error of writing up my own router, I finally gave up and decided
to settle with julienschmidt's httprouter, and it works just fine for me now.
I call it in this fashion.


``` go

var ctrl = &controller.Controller{}

func main() {

    // instantiate the router
    router := httprouter.New()

    // The application's routes
    router.GET("/", ctrl.Index)

}

```
