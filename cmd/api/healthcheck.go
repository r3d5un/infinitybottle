package main

import (
	"fmt"
	"net/http"
)

// Healthcheck godoc
//
//	@Summary		Basic healthcheck
//	@Description	Perform a basic request to check if the service is available
//	@Tags			healthcheck
//	@Produce		json
//	@Success		200	{object}    string
//	@Router			/v1/healthcheck [get]
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	js := `{"status": "available", "environment": %q, "version": %q}`
	js = fmt.Sprintf(js, app.config.env, version)

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(js))
}
