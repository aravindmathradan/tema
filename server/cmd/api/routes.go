package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/projects", app.listProjectsHandler)
	router.HandlerFunc(http.MethodPost, "/v1/projects", app.createProjectHandler)
	router.HandlerFunc(http.MethodGet, "/v1/projects/:id", app.viewProjectHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/projects/:id", app.updateProjectHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/projects/:id", app.deleteProjectHandler)

	router.HandlerFunc(http.MethodPost, "/v1/users", app.signupUserHandler)

	return app.recoverPanic(app.rateLimit(router))
}
