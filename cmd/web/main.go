package main

import (
	"fmt"
	"log"
	"time"

	"net/http"

	"github.com/AbilashKriz/bookings/pkg/config"
	"github.com/AbilashKriz/bookings/pkg/handlers"
	"github.com/AbilashKriz/bookings/pkg/renders"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

// Creating a new vaiable for the struct we created in config folder
var app config.AppConfig

var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()

	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	//Running the cretingTemplate function from rendres package
	tc, err := renders.CreatingTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	//assiging the output from the cretingtemplate function to the app variable from config folder

	app.TempCache = tc

	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)
	//In render function now we do not have to run the createtemplate function as we running it here , all we have to do is to create a pointer to the app variable so
	//so that we can use the values directly

	renders.NewTemplate(&app)

	srv := http.Server{
		Addr:    portNumber,
		Handler: Route(&app),
	}
	fmt.Println("listening on port", portNumber)
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
