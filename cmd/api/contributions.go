package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"infinitybottle.islandwind.me/internal/data"
)

func (app *application) listContributionsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "list brand bottles")
}

func (app *application) createContributionHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		InfinityBottleID int64    `json:"infinityBottleID"`
		Amount           int64    `json:"amount"`
		BrandName        string   `json:"brandName"`
		Tags             []string `json:"tags,omitempty"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

// @Summary		Get an infinity bottle contribution by ID
// @Description	Retrieve all information about an infinity bottle contribution by ID
// @Param          id		path	int	true	"ID"
// @Tags			contribution
// @Produce		json
// @Success		200	{object}    data.Contribution
// @Failure		404	{object}    ErrorMessage
// @Failure		500	{object}    ErrorMessage
// @Router			/v1/contributions/{id} [get]
func (app *application) getContributionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
	}

	contribution := data.Contribution{
		ID:               id,
		InfinityBottleID: 1,
		AddedAt:          time.Now(),
		Amount:           4,
		BrandName:        "Laphroaig",
		Tags:             []string{"peaty", "smokey"},
	}

	err = app.writeJSON(w, http.StatusOK, contribution, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
