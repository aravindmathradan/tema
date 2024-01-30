package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activate", app.activateUserHandler)

	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	protected := alice.New(app.requireActivatedUser)
	router.Handler(http.MethodGet, "/v1/projects", protected.ThenFunc(app.listProjectsHandler))
	router.Handler(http.MethodPost, "/v1/projects", protected.ThenFunc(app.createProjectHandler))
	router.Handler(http.MethodGet, "/v1/projects/:id", protected.ThenFunc(app.viewProjectHandler))
	router.Handler(http.MethodPatch, "/v1/projects/:id", protected.ThenFunc(app.updateProjectHandler))
	router.Handler(http.MethodDelete, "/v1/projects/:id", protected.ThenFunc(app.deleteProjectHandler))

	standard := alice.New(app.recoverPanic, app.rateLimit, app.authenticate)
	return standard.Then(router)
}
