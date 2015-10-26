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
- [x] Load assets in public folder
- [ ] Load the routes from a seperate routes.go file
- [ ] Pass on data to the Templates
- [ ] Load templates dynamically
- [ ] Connect to the Database
- [ ] Load Models
- [ ] Explore GoRoutines



======================================================================



The controller serving a view has been simplified dramatically.



``` go

func (p *Page) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {


p.ServeView("index.html", w)

}

```



After many trial error of writing up my own router, I finally gave up and decided
to settle with julienschmidt's httprouter, and it works just fine for me now.
I call it in this fashion.




``` go

var ctrl = &controller.ControllersInit{}

func main() {

    // instantiate the router
    router := httprouter.New()

    // The application's routes
    router.GET("/", ctrl.Pages.Index)

}

```


=========================================================




##### How to load the templates dynamically?

The templates currently are loaded in the controller inside the bootstrap folder.

Possible ways to load it dynamically:

- [xyz]


##### Convenient way to set up the routes

Currently the routes are in the main package loaded in the main function.
This however would serve to make things more complicated especially when the
application grows dramatically and you have dozens of routes. Although I don't
have it entirely figured it out, but having a routes.go file seems more ideal
at the moment.
