package main

import (
	"net/http"
)

type HealthCheckMessage struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

// Healthcheck godoc
//
//	@Summary		Basic healthcheck
//	@Description	Perform a basic request to check if the service is available
//	@Tags			healthcheck
//	@Produce		json
//	@Success		200	{object}    HealthCheckMessage
//	@Failure		500	{object}    ErrorMessage
//	@Router			/v1/healthcheck [get]
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	healthCheckMessage := HealthCheckMessage{
		Status:      "available",
		Environment: app.config.env,
		Version:     version,
	}

	err := app.writeJSON(w, http.StatusOK, healthCheckMessage, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
