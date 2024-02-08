package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/aravindmathradan/tema/internal/data"
	"github.com/aravindmathradan/tema/internal/validator"
)

func (app *application) createAuthenticationTokenHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email                 string `json:"email"`
		Password              string `json:"password"`
		Scope                 string `json:"scope"`
		RefreshTokenPlainText string `json:"refresh_token"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	data.ValidateTokenScope(v, input.Scope)

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	if input.Scope == data.ScopeAuthentication {
		data.ValidateEmail(v, input.Email)
		data.ValidatePasswordPlaintext(v, input.Password)

		if !v.Valid() {
			app.failedValidationResponse(w, r, v.Errors)
			return
		}

		user, err := app.models.Users.GetByEmail(input.Email)
		if err != nil {
			switch {
			case errors.Is(err, data.ErrRecordNotFound):
				app.invalidCredentialsResponse(w, r)
			default:
				app.serverErrorResponse(w, r, err)
			}
			return
		}

		match, err := user.Password.Matches(input.Password)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		if !match {
			app.invalidCredentialsResponse(w, r)
			return
		}

		refreshToken, err := app.models.Tokens.New(user.ID, 30*24*time.Hour, data.ScopeRefresh)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		authToken, err := app.models.Tokens.New(user.ID, 24*time.Hour, data.ScopeAuthentication)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		err = app.writeJSON(w, http.StatusCreated, envelope{"authentication_token": authToken, "refresh_token": refreshToken}, nil)
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
	} else if input.Scope == data.ScopeRefresh {
		if data.ValidateTokenPlaintext(v, input.RefreshTokenPlainText); !v.Valid() {
			app.failedValidationResponse(w, r, v.Errors)
			return
		}

		user, err := app.models.Users.GetForToken(data.ScopeRefresh, input.RefreshTokenPlainText)
		if err != nil {
			switch {
			case errors.Is(err, data.ErrRecordNotFound):
				v.AddError("refresh_token", validator.EINVALIDTOKEN, "invalid or expired refresh token")
				app.failedValidationResponse(w, r, v.Errors)
			default:
				app.serverErrorResponse(w, r, err)
			}
			return
		}

		authToken, err := app.models.Tokens.New(user.ID, 24*time.Hour, data.ScopeAuthentication)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		err = app.writeJSON(w, http.StatusCreated, envelope{"authentication_token": authToken}, nil)
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
	}
}

func (app *application) createPasswordResetTokenHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email string `json:"email"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	if data.ValidateEmail(v, input.Email); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	user, err := app.models.Users.GetByEmail(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			v.AddError("email", validator.ENOTFOUND, "no matching email address found")
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	if !user.Activated {
		v.AddError("email", validator.EACCOUNTINACTIVE, "user account must be activated by verifying the email")
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	token, err := app.models.Tokens.New(user.ID, 15*time.Minute, data.ScopePasswordReset)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.background(func() {
		data := map[string]any{
			"passwordResetToken": token.Plaintext,
			"userName":           user.Name,
		}

		app.mailer.Send(user.Email, "password_reset_token.tmpl", data)
		if err != nil {
			// Importantly, if there is an error sending the email then we use the
			// app.logger.Error() helper to manage it, instead of the app.serverErrorResponse().
			// This is because by the time we encounter the errors, the client will probably
			// have already been sent a 202 Accepted response by our writeJSON() helper.
			app.logger.Error(err.Error())
		}
	})

	env := envelope{"message": "an email will be sent to you containing the password reset instructions"}

	err = app.writeJSON(w, http.StatusAccepted, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createActivationTokenHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email string `json:"email"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	if data.ValidateEmail(v, input.Email); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	user, err := app.models.Users.GetByEmail(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			v.AddError("email", validator.ENOTFOUND, "no matching email address found")
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	if user.Activated {
		v.AddError("email", validator.EALREADYACTIVE, "account has already been activated")
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	token, err := app.models.Tokens.New(user.ID, 3*24*time.Hour, data.ScopeActivation)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.background(func() {
		data := map[string]any{
			"activationToken": token.Plaintext,
			"userName":        user.Name,
		}

		err = app.mailer.Send(user.Email, "activation_token.tmpl", data)
		if err != nil {
			app.logger.Error(err.Error())
		}
	})

	env := envelope{"message": "an email will be sent containing the instructions to activate your account"}

	err = app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteAuthenticationTokenHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)

	err := app.models.Tokens.DeleteAllForUser(data.ScopeAuthentication, user.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.models.Tokens.DeleteAllForUser(data.ScopeRefresh, user.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "auth token deleted from database"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
