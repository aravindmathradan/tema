package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"

	"github.com/aravindmathradan/tema/internal/validator"
)

type errorCode string

const (
	ENOTFOUND           errorCode = "not_found"
	EINTERNAL           errorCode = "internal"
	EMETHODNOTALLOWED   errorCode = "method_not_allowed"
	EFAILEDVALIDATION   errorCode = "validation_failed"
	EEDITCONFLICT       errorCode = "edit_conflict"
	ERATELIMITEXCEEDED  errorCode = "rate_limit_exceeded"
	EINVALIDCREDENTIALS errorCode = "invalid_credentials"
	EINVALIDTOKEN       errorCode = "invalid_token"
	EBADREQUEST         errorCode = "bad_request"
	EAUTHREQUIRED       errorCode = "authentication_required"
	EINACTIVEACCOUNT    errorCode = "inactive_account"
	ENOTPERMITTED       errorCode = "not_permitted"
)

type errorRes struct {
	code    errorCode
	message string
	fields  map[string]validator.FieldError
}

func (app *application) logError(r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = string(debug.Stack())
	)

	app.logger.Error(
		err.Error(),
		slog.String("method", method),
		slog.String("uri", uri),
		slog.String("trace", trace),
	)
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message errorRes) {
	err := app.writeJSON(w, status, envelope{"error": message}, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	res := errorRes{
		code:    EINTERNAL,
		message: "the server encountered a problem and could not process your request",
		fields:  map[string]validator.FieldError{},
	}
	app.errorResponse(w, r, http.StatusInternalServerError, res)
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	res := errorRes{
		code:    ENOTFOUND,
		message: "the requested resource could not be found",
		fields:  map[string]validator.FieldError{},
	}
	app.errorResponse(w, r, http.StatusNotFound, res)
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	res := errorRes{
		code:    EMETHODNOTALLOWED,
		message: fmt.Sprintf("the %s method is not supported for this resource", r.Method),
		fields:  map[string]validator.FieldError{},
	}
	app.errorResponse(w, r, http.StatusMethodNotAllowed, res)
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	res := errorRes{
		code:    EBADREQUEST,
		message: err.Error(),
		fields:  map[string]validator.FieldError{},
	}
	app.errorResponse(w, r, http.StatusBadRequest, res)
}

func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]validator.FieldError) {
	res := errorRes{
		code:    EFAILEDVALIDATION,
		message: "input validation failed. please check input fields.",
		fields:  errors,
	}
	app.errorResponse(w, r, http.StatusUnprocessableEntity, res)
}

func (app *application) editConflictResponse(w http.ResponseWriter, r *http.Request) {
	res := errorRes{
		code:    EEDITCONFLICT,
		message: "unable to update the record due to an edit conflict, please try again",
		fields:  map[string]validator.FieldError{},
	}
	app.errorResponse(w, r, http.StatusConflict, res)
}

func (app *application) rateLimitExceededResponse(w http.ResponseWriter, r *http.Request) {
	res := errorRes{
		code:    ERATELIMITEXCEEDED,
		message: "rate limit exceeded",
		fields:  map[string]validator.FieldError{},
	}
	app.errorResponse(w, r, http.StatusTooManyRequests, res)
}

func (app *application) invalidCredentialsResponse(w http.ResponseWriter, r *http.Request) {
	res := errorRes{
		code:    EINVALIDCREDENTIALS,
		message: "invalid authentication credentials",
		fields:  map[string]validator.FieldError{},
	}
	app.errorResponse(w, r, http.StatusUnauthorized, res)
}

func (app *application) invalidAuthenticationTokenResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("WWW-Authenticate", "Bearer")

	res := errorRes{
		code:    EINVALIDTOKEN,
		message: "invalid or missing authentication token",
		fields:  map[string]validator.FieldError{},
	}
	app.errorResponse(w, r, http.StatusUnauthorized, res)
}

func (app *application) authenticationRequiredResponse(w http.ResponseWriter, r *http.Request) {
	res := errorRes{
		code:    EAUTHREQUIRED,
		message: "you must be authenticated to access this resource",
		fields:  map[string]validator.FieldError{},
	}
	app.errorResponse(w, r, http.StatusUnauthorized, res)
}

func (app *application) inactiveAccountResponse(w http.ResponseWriter, r *http.Request) {
	res := errorRes{
		code:    EINACTIVEACCOUNT,
		message: "your account must be activated to access this resource",
		fields:  map[string]validator.FieldError{},
	}
	app.errorResponse(w, r, http.StatusForbidden, res)
}

func (app *application) notPermittedResponse(w http.ResponseWriter, r *http.Request) {
	res := errorRes{
		code:    ENOTPERMITTED,
		message: "your user account doesn't have the necessary permissions to access this resource",
		fields:  map[string]validator.FieldError{},
	}
	app.errorResponse(w, r, http.StatusForbidden, res)
}
