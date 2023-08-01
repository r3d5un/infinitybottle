package main

import (
	"net/http"
)

// Healthcheck godoc
//
//	@Summary		Basic healthcheck
//	@Description	Perform a basic request to check if the service is available
//	@Tags			healthcheck
//	@Produce		json
//	@Success		200	{object}    ErrorMessage
//	@Failure		500	{object}    ErrorMessage
//	@Router			/v1/healthcheck [get]
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
