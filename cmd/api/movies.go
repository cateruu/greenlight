package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cateruu/greenlight/internals/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create a movie...\n")
}

func (app *application) getMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "spider-man",
		Runtime:   100,
		Genres:    []string{"action", "romance"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
