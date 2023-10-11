package main

import (
	"errors"
	"fmt"
	"net/http"
	"snippetbox.disconico/internal/models"
	"strconv"
	"time"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFoundError(w)
		return
	}

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	for i := range snippets {
		snippets[i].CreatedInDays = int(time.Now().Sub(snippets[i].Created).Hours() / 24)
	}

	data := app.newTemplateData()
	data.Snippets = snippets

	app.render(w, r, http.StatusOK, "home.tmpl.html", data)
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFoundError(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFoundError(w)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	snippet.CreatedInDays = int(time.Now().Sub(snippet.Created).Hours() / 24)
	snippet.ExpiresInDays = int(snippet.Expires.Sub(time.Now()).Hours() / 24)

	data := app.newTemplateData()
	data.Snippet = snippet

	app.render(w, r, http.StatusOK, "view.tmpl.html", data)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	var (
		title   = "0 snail"
		content = "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
		expires = 7
	)

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
