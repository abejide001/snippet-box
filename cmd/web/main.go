package main

import (
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

var (
	errorLog = log.New(os.Stderr, "ERROR \t", log.Ldate|log.Ltime)
	infoLog  = log.New(os.Stdout, "INFO \t", log.Ldate|log.Ltime)
)

func main() {
	// Initialize a new instance of application containing the dependencies.
	app := application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	app.infoLog.Println("Starting on port 4000")
	err := http.ListenAndServe(":4000", app.routes())
	app.errorLog.Fatal(err)
}
