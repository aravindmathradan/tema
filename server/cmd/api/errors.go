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
	Code    errorCode                       `json:"code"`
	Message string                          `json:"message"`
	Fields  map[string]validator.FieldError `json:"fields"`
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

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, res errorRes) {
	err := app.writeJSON(w, status, envelope{"error": res}, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	res := errorRes{
		Code:    EINTERNAL,
		Message: "the server encountered a problem and could not process your request",
		Fields:  map[string]validator.FieldError{},
	}
	app.errorResponse(w, r, http.StatusInternalServerError, res)
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	res := errorRes{
		Code:    ENOTFOUND,
		Message: "the requested resource could not be found",
		Fields:  map[string]validator.FieldError{},
	}
	app.errorResponse(w, r, http.StatusNotFound, res)
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	res := errorRes{
		Code:    EMETHODNOTALLOWED,
		Message: fmt.Sprintf("the %s method is not supported for this resource", r.Method),
		Fields:  map[string]validator.FieldError{},
	}
	app.errorResponse(w, r, http.StatusMethodNotAllowed, res)
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	res := errorRes{
		Code:    EBADREQUEST,
		Message: err.Error(),
		Fields:  map[string]validator.FieldError{},
	}
	app.errorResponse(w, r, http.StatusBadRequest, res)
}

func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]validator.FieldError) {
	res := errorRes{
		Code:    EFAILEDVALIDATION,
		Message: "input validation failed",
		Fields:  errors,
	}
	app.errorResponse(w, r, http.StatusUnprocessableEntity, res)
}

func (app *application) editConflictResponse(w http.ResponseWriter, r *http.Request) {
	res := errorRes{
		Code:    EEDITCONFLICT,
		Message: "unable to update the record due to an edit conflict, please try again",
		Fields:  map[string]validator.FieldError{},
	}
	app.errorResponse(w, r, http.StatusConflict, res)
}

func (app *application) rateLimitExceededResponse(w http.ResponseWriter, r *http.Request) {
	res := errorRes{
		Code:    ERATELIMITEXCEEDED,
		Message: "rate limit exceeded",
		Fields:  map[string]validator.FieldError{},
	}
	app.errorResponse(w, r, http.StatusTooManyRequests, res)
}

func (app *application) invalidCredentialsResponse(w http.ResponseWriter, r *http.Request) {
	res := errorRes{
		Code:    EINVALIDCREDENTIALS,
		Message: "invalid credentials",
		Fields:  map[string]validator.FieldError{},
	}
	app.errorResponse(w, r, http.StatusUnauthorized, res)
}

func (app *application) invalidAuthenticationTokenResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("WWW-Authenticate", "Bearer")

	res := errorRes{
		Code:    EINVALIDTOKEN,
		Message: "invalid or missing authentication token",
		Fields:  map[string]validator.FieldError{},
	}
	app.errorResponse(w, r, http.StatusUnauthorized, res)
}

func (app *application) authenticationRequiredResponse(w http.ResponseWriter, r *http.Request) {
	res := errorRes{
		Code:    EAUTHREQUIRED,
		Message: "you must be authenticated to access this resource",
		Fields:  map[string]validator.FieldError{},
	}
	app.errorResponse(w, r, http.StatusUnauthorized, res)
}

func (app *application) inactiveAccountResponse(w http.ResponseWriter, r *http.Request) {
	res := errorRes{
		Code:    EINACTIVEACCOUNT,
		Message: "your account must be activated to access this resource",
		Fields:  map[string]validator.FieldError{},
	}
	app.errorResponse(w, r, http.StatusForbidden, res)
}

func (app *application) notPermittedResponse(w http.ResponseWriter, r *http.Request) {
	res := errorRes{
		Code:    ENOTPERMITTED,
		Message: "your user account doesn't have the necessary permissions to access this resource",
		Fields:  map[string]validator.FieldError{},
	}
	app.errorResponse(w, r, http.StatusForbidden, res)
}
