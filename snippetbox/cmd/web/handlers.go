package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"log"
)

func home(w http.ResponseWriter, r *http.Request) {
	// Check if the requested URL path is "/"
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Define the file paths for the templates
	files := []string{
		"./ui/html/base.html",
		"./ui/html/home.html",
	}

	// Parse the templates
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute the "home.html" template
	err = ts.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	// w.Header().Set("Content-type", "application/json")
	// w.Write([]byte(`{"name":"Alex"}`))
	//w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1{
		http.NotFound(w, r);
		return;
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)	
	//w.Write([]byte("Display this snippet"));
}

func snippetCreate(w http.ResponseWriter, r *http.Request){
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
