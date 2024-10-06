package main

import (
	"flag"
	"fmt"
	"go-breeders/configuration"
	"go-breeders/models"
	"html/template"
	"log"
	"net/http"
	"time"
)

const port = ":4000"

type application struct {
	templateMap map[string]*template.Template
	config      appConfig
	Models      models.Models
	App         *configuration.Application
}

type appConfig struct {
	useCache bool
	dsn      string
}

func main() {
	app := application{
		templateMap: make(map[string]*template.Template),
	}

	// parseTime: makes go time type work well with time values or date values inside of the database
	// Collation: specify the character set for a string alues UTF-8 m
	flag.StringVar(
		&app.config.dsn,
		"dsn",
		`mariadb:myverysecretpassword@tcp(localhost:3306)/breeders?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s`,
		"DSN",
	)
	flag.BoolVar(&app.config.useCache, "cache", false, "Use template cache")
	flag.Parse()

	// Get database
	db, err := initMySQLDB(app.config.dsn)

	if err != nil {
		log.Panic(err)
		return
	}
	
	app.App = configuration.New(db)

	srv := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	fmt.Println("Starting web application on port", port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

func (app *application) TestPatterns(w http.ResponseWriter, r *http.Request) {
	app.render(w, "test.page.gohtml", nil)
}
