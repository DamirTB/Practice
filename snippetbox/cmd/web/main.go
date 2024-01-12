package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct{
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	
	// logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil));

	app := &application{
		logger: logger,
	}
	logger.Info("starting server", "addr", slog.String("addr", ":4000"));
	err := http.ListenAndServe(*addr, app.routes());
	logger.Error(err.Error());
	os.Exit(1);
	//logger.Printf("Starting server on http://localhost:%s/", *addr)
	//logger.Info("port on http://localhost:4000/")
	// err := http.ListenAndServe(":4000", mux)
	// log.Fatal(err)
}
