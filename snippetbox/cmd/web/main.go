package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil));

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home);
	mux.HandleFunc("/snippet/view", snippetView);
	mux.HandleFunc("/snippet/create", snippetCreate);
	mux.HandleFunc("/snippet/delete", snippetDelete);

	logger.Info("starting server", "addr", slog.String("addr", ":4000"));
	err := http.ListenAndServe(*addr, mux);
	logger.Error(err.Error());
	os.Exit(1);
	//logger.Printf("Starting server on http://localhost:%s/", *addr)
	//logger.Info("port on http://localhost:4000/")
	// err := http.ListenAndServe(":4000", mux)
	// log.Fatal(err)
}
