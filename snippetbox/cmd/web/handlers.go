package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	_"log"
)

func(app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/base.html",
		"./ui/html/home.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	// w.Header().Set("Content-type", "application/json")
	// w.Write([]byte(`{"name":"Alex"}`))
	//w.Write([]byte("Hello from Snippetbox"))
}

func(app *application) snippetView(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1{
		http.NotFound(w, r);
		return;
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)	
	//w.Write([]byte("Display this snippet"));
}

func(app *application) snippetCreate(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/snippet/create"{
		http.NotFound(w, r);
		return;
	}
	if r.Method != http.MethodPost{
		w.WriteHeader(405);
		http.Error(w, "only POST Method is allowed", http.StatusMethodNotAllowed);
		//w.Write([]byte("Method now allowed"));
		return;
	}
	w.Write([]byte("Create a new snippet"));
}

func(app *application) snippetDelete(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/snippet/delete"{
		http.NotFound(w, r);
		return;
	}
	if r.Method != http.MethodDelete{
		w.WriteHeader(405);
		http.Error(w, "only DELETE Method is allowed", http.StatusMethodNotAllowed);
		//w.Write([]byte("Method now allowed"));
		return;
	}
	w.Write([]byte("Delete a new snippet"));
}

