package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Aghloo/bookings/config"
	"github.com/Aghloo/bookings/pkg/handlers"
	"github.com/Aghloo/bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8090"

var App config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {

	//Change this to true when in production
	App.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = App.InProduction

	App.Session = session

	tc, err := render.CreateTemplatecache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	App.TemplateCache = tc
	App.UseCahe = false
	repo := handlers.NewRepo(&App)
	handlers.NewHandler(repo)

	render.NewTemplates(&App)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	//_ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&App),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
