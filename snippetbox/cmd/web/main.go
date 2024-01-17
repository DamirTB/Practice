package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"snippetbox.alexedwards.net/internal/models"
)

type application struct {
	logger *slog.Logger
	snippets *models.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	// logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil));

	db, err := OpenDB(*dsn)
	if err != nil{
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	app := &application{
		logger: logger,
		snippets: &models.SnippetModel{DB: db},
	}
	logger.Info("starting server", "addr", slog.String("addr", ":4000"));
	err = http.ListenAndServe(*addr, app.routes());
	logger.Error(err.Error());
	os.Exit(1);
	//logger.Printf("Starting server on http://localhost:%s/", *addr)
	//logger.Info("port on http://localhost:4000/")
	// err := http.ListenAndServe(":4000", mux)
	// log.Fatal(err)
}

func OpenDB(dsn string) (*sql.DB, error){
	db, err := sql.Open("mysql", dsn)
	if err != nil{
		return nil, err;
	}
	err = db.Ping()
	if err != nil{
		db.Close()
		return nil, err;
	}
	return db, nil;
}