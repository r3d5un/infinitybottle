package main

import (
	"fmt"
	"net/http"
	"time"

	"infinitybottle.islandwind.me/internal/data"
	"infinitybottle.islandwind.me/internal/validator"
)

type InfinityBottlePost struct {
	BottleName string `json:"bottleName"`
}

func (app *application) listInfinityBottlesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "list infinity bottles")
}

// @Summary		Create a new infinity bottle
// @Description	Create a new infinity bottle
// @Tags			infinityBottle
// @Produce		json
// @Accept     json
// @Param      InfinityBottle	body	InfinityBottlePost	true	"New infinity bottle"
// @Success		201	{object}    InfinityBottlePost
// @Failure		400	{object}    ErrorMessage
// @Failure		404	{object}    ErrorMessage
// @Failure		500	{object}    ErrorMessage
// @Router			/v1/infinitybottles [post]
func (app *application) createInfinityBottleHandler(w http.ResponseWriter, r *http.Request) {
	infinityBottlePost := InfinityBottlePost{}
	err := app.readJSON(w, r, &infinityBottlePost)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	infinityBottle := &data.InfinityBottle{
		BottleName: infinityBottlePost.BottleName,
	}

	v := validator.New()

	if data.ValidateInfinityBottle(v, infinityBottle); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.InfinityBottles.Insert(infinityBottle)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/infinitybottles/%d", infinityBottle.ID))

	err = app.writeJSON(w, http.StatusCreated, infinityBottle, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// @Summary		Get an infinity bottle by ID
// @Description	Retrieve all information about an infinity bottle by ID
// @Param          id		path	int	true	"ID"
// @Tags			infinityBottle
// @Produce		json
// @Success		200	{object}    data.InfinityBottle
// @Failure		404	{object}    ErrorMessage
// @Failure		500	{object}    ErrorMessage
// @Router			/v1/infinitybottles/{id} [get]
func (app *application) getInfinityBottleHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	infinityBottle := data.InfinityBottle{
		ID:                    id,
		BottleName:            "Mister Smokey",
		NumberOfContributions: 23,
		CreatedAt:             time.Now(),
		UpdatedAt:             time.Now(),
		Contributions: []data.Contribution{
			{
				ID:               1,
				InfinityBottleID: 1,
				AddedAt:          time.Now(),
				Amount:           4,
				BrandName:        "Laphroaig",
				Tags:             []string{"peaty", "smokey"},
			},
			{
				ID:               2,
				InfinityBottleID: 1,
				AddedAt:          time.Now(),
				Amount:           4,
				BrandName:        "Ardbeg",
				Tags:             []string{"peaty", "smokey"},
			},
		},
	}

	err = app.writeJSON(w, http.StatusOK, infinityBottle, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
