package main

import (
	"fmt"
	"free-mentors-api/internal/driver"
	"log"
	"net/http"
	"os"
)

type config struct {
	port int
}

// shared configuration properties
type application struct {
	config   config
	infoLog  *log.Logger // prints info to the console
	errorLog *log.Logger
	db       *driver.DB
}

func main() {
	var config config
	config.port = 9800

	// Define infoLog and errorLog structure output in the console
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	dsn := os.Getenv("DSN")
	database, err := driver.ConnectPostgres(dsn)

	if err != nil {
		log.Fatal("Cannot connect to db ", err)
	}
	defer database.SQL.Close()

	// create a variable app that refers to the application struct
	app := &application{
		config:   config,
		infoLog:  infoLog,
		errorLog: errorLog,
		db:       database,
	}

	// start the web server
	err = app.serve()
	if err != nil {
		log.Fatal(err)
	}
}

func (app *application) serve() error {
	app.infoLog.Printf("Server running on port: %v ", app.config.port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
	}
	return srv.ListenAndServe()
}
