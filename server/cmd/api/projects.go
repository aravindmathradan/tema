package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aravindmathradan/tema/internal/data"
)

func (app *application) createProjectHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create new project")
}

func (app *application) viewProjectHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	project := data.Project{
		ID:          id,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Name:        "Casablanca",
		Description: "I am iron man",
		Status:      1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"project": project}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
