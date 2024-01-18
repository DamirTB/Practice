package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"snippetbox.alexedwards.net/internal/models"
	"github.com/julienschmidt/httprouter" // New import
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Because httprouter matches the "/" path exactly, we can now remove the
	// manual check of r.URL.Path != "/" from this handler.
	snippets, err := app.snippets.Latest()
	if err != nil {
	app.serverError(w, r, err)
	return
	}
	data := app.newTemplateData(r)
	data.Snippets = snippets
	app.render(w, r, http.StatusOK, "home.html", data)
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	// When httprouter is parsing a request, the values of any named parameters
	// will be stored in the request context. We'll talk about request context
	// in detail later in the book, but for now it's enough to know that you can
	// use the ParamsFromContext() function to retrieve a slice containing these
	// parameter names and values like so:
	params := httprouter.ParamsFromContext(r.Context())
	// We can then use the ByName() method to get the value of the "id" named
	// parameter from the slice and validate it as normal.
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
	app.notFound(w)
	return
	}
	snippet, err := app.snippets.Get(id)
	if err != nil {
	if errors.Is(err, models.ErrNoRecord) {
	app.notFound(w)
	} else {
	app.serverError(w, r, err)
	}
	return
	}
	data := app.newTemplateData(r)
	data.Snippet = snippet
	app.render(w, r, http.StatusOK, "view.html", data)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display the form for creating a new snippet..."))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// Checking if the request method is a POST is now superfluous and can be
	// removed, because this is done automatically by httprouter.
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
	app.serverError(w, r, err)
	return
	}
	// Update the redirect path to use the new clean URL format.
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
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

