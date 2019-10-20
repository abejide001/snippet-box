package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"alexedwards.net/snippetbox/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *mysql.SnippetModel
}

var (
	errorLog = log.New(os.Stderr, "ERROR \t", log.Ldate|log.Ltime)
	infoLog  = log.New(os.Stdout, "INFO \t", log.Ldate|log.Ltime)
)

func main() {
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()
	flag.Parse()

	// Initialize a new instance of application containing the dependencies.
	app := application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &mysql.SnippetModel{DB: db},
	}

	app.infoLog.Println("Starting on port 4000")
	err = http.ListenAndServe(":4000", app.routes())
	app.errorLog.Fatal(err)

}

// The openDB() function wraps sql.Open() and returns a sql.DB connection pool
// for a given DSN.
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
