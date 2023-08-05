package main

import (
	"errors"
	"fmt"
	"net/http"

	"infinitybottle.islandwind.me/internal/data"
	"infinitybottle.islandwind.me/internal/validator"
)

type InfinityBottlePost struct {
	BottleName string `json:"bottleName"`
	EmptyStart bool   `json:"emptyStart,omitempty"`
}

// @Summary		List all infinity bottles
// @Description	List all infinity bottles
// @Tags			infinityBottle
// @Produce		json
//
//	@Param			bottle_name	query		string	false	"bottle name to search for"
//
// @Success		200	{array}     data.InfinityBottle
// @Failure		400	{object}    ErrorMessage
// @Failure		404	{object}    ErrorMessage
// @Failure		500	{object}    ErrorMessage
// @Router			/v1/infinitybottles [get]
func (app *application) listInfinityBottlesHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		BottleName string
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.BottleName = app.readStrings(qs, "bottleName", "")

	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
	input.Filters.Sort = app.readStrings(qs, "sort", "id")
	input.Filters.SortSafelist = []string{"id", "-id", "bottle_name", "-bottle_name"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	infinityBottles, err := app.models.InfinityBottles.GetAll(
		input.BottleName,
		input.Filters,
	)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, infinityBottles, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
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
		EmptyStart: infinityBottlePost.EmptyStart,
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

	infinityBottle, err := app.models.InfinityBottles.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, infinityBottle, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// @Summary		Update an infinity bottle by ID
// @Description	Update all information about an infinity bottle by ID
// @Param          id		path	int	true	"ID"
// @Param      InfinityBottle	body	InfinityBottlePost	true	"Update to an infinity bottle"
// @Tags			infinityBottle
// @Produce		json
// @Success		200	{object}    data.InfinityBottle
// @Failure		404	{object}    ErrorMessage
// @Failure		500	{object}    ErrorMessage
// @Router			/v1/infinitybottles/{id} [put]
func (app *application) updateInfinityBottleHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	infinityBottle, err := app.models.InfinityBottles.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	input := InfinityBottlePost{}
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	infinityBottle.BottleName = input.BottleName
	infinityBottle.EmptyStart = input.EmptyStart

	v := validator.New()

	if data.ValidateInfinityBottle(v, infinityBottle); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.InfinityBottles.Update(infinityBottle)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, infinityBottle, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// @Summary		Delete an infinity bottle by ID
// @Description	Delete an infinity bottle by ID
// @Param          id		path	int	true	"ID"
// @Tags			infinityBottle
// @Produce		json
// @Success		200
// @Failure		404	{object}    ErrorMessage
// @Failure		500	{object}    ErrorMessage
// @Router			/v1/infinitybottles/{id} [delete]
func (app *application) deleteInfinityBottleHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.InfinityBottles.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, nil, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
