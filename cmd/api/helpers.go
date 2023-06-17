package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1048576

	// Get the body
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	// create a decoder
	decorder := json.NewDecoder(r.Body)

	err := decorder.Decode(data)

	if err != nil {
		return err
	}

	err = decorder.Decode(&struct{}{})

	if err != io.EOF {
		return errors.New("body must have only a single json value")
	}

	return nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")

	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(out)

	if err != nil {
		return err
	}

	return nil
}
