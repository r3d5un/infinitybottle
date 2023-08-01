package main

import (
	"fmt"
	"net/http"
	"time"

	"infinitybottle.islandwind.me/internal/data"
)

func (app *application) listInfinityBottlesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "list infinity bottles")
}

func (app *application) createInfinityBottleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create infinity bottle")
}

// @Summary		Get an infinity bottle by ID
// @Description	Retrieve all information about an infinity bottle by ID
// @Param          id		path	int	true	"ID"
// @Tags			infinityBottle
// @Produce		json
// @Success		200	{object}    data.InfinityBottle
// @Failure		404	{object}    string
// @Failure		500	{object}    string
// @Router			/v1/infinitybottles/{id} [get]
func (app *application) getInfinityBottleHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
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
		app.logger.Print(err)
		http.Error(
			w,
			"The server encountered a problem and could not process your request.",
			http.StatusInternalServerError,
		)
	}
}
