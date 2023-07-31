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
//	@Success		200	{object}    string
//	@Router			/v1/healthcheck [get]
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}
