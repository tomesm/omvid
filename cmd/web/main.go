package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tomesm/virtd/pkg/models/mysql"
)

// application holds application-wide dependencies for the web app.
type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	courses       *mysql.CourseModel
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "admin:pass@/virtd?parseTime=true", "MySQL data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}
	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		courses:       &mysql.CourseModel{DB: db},
		templateCache: templateCache,
	}
	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	infoLog.Printf("Starting server on %s\n", *addr)
	err = server.ListenAndServe()
	errorLog.Fatal(err)
}

// Wrapper function for sql.Open() and returns sql.DB connection pool
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
