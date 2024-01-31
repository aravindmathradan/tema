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
	router.HandlerFunc(http.MethodPut, "/v1/users/password", app.updateUserPasswordHandler)

	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/activation", app.createActivationTokenHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/password-reset", app.createPasswordResetTokenHandler)

	router.Handler(http.MethodGet, "/v1/projects", app.requirePermission("projects:read", http.HandlerFunc(app.listProjectsHandler)))
	router.Handler(http.MethodPost, "/v1/projects", app.requirePermission("projects:write", http.HandlerFunc(app.createProjectHandler)))
	router.Handler(http.MethodGet, "/v1/projects/:id", app.requirePermission("projects:read", http.HandlerFunc(app.viewProjectHandler)))
	router.Handler(http.MethodPatch, "/v1/projects/:id", app.requirePermission("projects:write", http.HandlerFunc(app.updateProjectHandler)))
	router.Handler(http.MethodDelete, "/v1/projects/:id", app.requirePermission("projects:write", http.HandlerFunc(app.deleteProjectHandler)))

	standard := alice.New(app.recoverPanic, app.enableCORS, app.rateLimit, app.authenticate)
	return standard.Then(router)
}
