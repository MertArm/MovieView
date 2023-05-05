package main

import (
	"errors"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "MovieView up and running",
		Version: "1.0.0",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.DB.AllMovies()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, movies)
}

func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	// Read JSON payload
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Validate user against DB
	user, err := app.DB.GetUserByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("Invalid credentials"), http.StatusBadRequest)
		return
	}

	// Check password
	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("Invalid credentials"), http.StatusBadRequest)
		return
	}

	// Create JWT user
	u := jwtUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	// Generate tokens
	tokens, err := app.auth.GenerateTokenPair(&u)

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	refreshCookie := app.auth.GetRefreshCookie(tokens.RefreshToken)
	http.SetCookie(w, refreshCookie)

	app.writeJSON(w, http.StatusAccepted, tokens)
}
