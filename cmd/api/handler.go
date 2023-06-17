package main

import "net/http"

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	type credentials struct {
		Username string `json:"email"`
		Password string `json:"password"`
	}

	var payload jsonResponse

	var creds credentials

	err := app.readJSON(w, r, &creds)

	if err != nil {
		app.errorLog.Println(err)
		payload.Error = true
		payload.Message = "Invalid json supplied, or missing json fields"

		_ = app.writeJSON(w, http.StatusBadRequest, payload)
	}

	err = app.writeJSON(w, http.StatusOK, payload)

	if err != nil {
		app.errorLog.Println(err)
	}
}
