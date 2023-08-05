package main

import (
	"errors"
	"fmt"
	"net/http"

	"infinitybottle.islandwind.me/internal/data"
	"infinitybottle.islandwind.me/internal/validator"
)

type ContributionPost struct {
	InfinityBottleID int64    `json:"infinityBottleID"`
	Amount           int64    `json:"amount"`
	BrandName        string   `json:"brandName"`
	Tags             []string `json:"tags,omitempty"`
}

// @Summary		List all infinity bottle contributions
// @Description	List all infinity bottles contributions
// @Tags			contribution
// @Produce		json
//
//	@Param			brand_name	query		string	false	"brand name to search for"
//	@Param			tags	query		string	false	"tags to search for"
//
// @Success		200	{array}     data.Contribution
// @Failure		400	{object}    ErrorMessage
// @Failure		404	{object}    ErrorMessage
// @Failure		500	{object}    ErrorMessage
// @Router			/v1/contribution [get]
func (app *application) listContributionsHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		BrandName string
		Tags      []string
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.BrandName = app.readStrings(qs, "brandName", "")
	input.Tags = app.readCSV(qs, "tags", []string{})

	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_sixe", 20, v)
	input.Filters.Sort = app.readStrings(qs, "sort", "id")
	input.Filters.SortSafelist = []string{"id", "-id", "brand_name", "-brand_name"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	contributions, err := app.models.Contributions.GetAll(
		input.BrandName,
		input.Tags,
		input.Filters,
	)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, contributions, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// @Summary		Add a new contribution to an infinity bottle
// @Description	Add a new contribution to an infinity bottle
// @Tags			contribution
// @Produce		json
// @Accept     json
// @Param      Contribution	body	ContributionPost	true	"New contribution to an infinity bottle"
// @Success		201	{object}    ContributionPost
// @Failure		400	{object}    ErrorMessage
// @Failure		404	{object}    ErrorMessage
// @Failure		500	{object}    ErrorMessage
// @Router			/v1/contributions [post]
func (app *application) createContributionHandler(w http.ResponseWriter, r *http.Request) {
	contributionPost := ContributionPost{}
	err := app.readJSON(w, r, &contributionPost)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	contribution := &data.Contribution{
		InfinityBottleID: contributionPost.InfinityBottleID,
		Amount:           contributionPost.Amount,
		BrandName:        contributionPost.BrandName,
		Tags:             contributionPost.Tags,
	}

	v := validator.New()

	if data.ValidateContribution(v, contribution); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Contributions.Insert(contribution)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/contributions/%d", contribution.ID))

	err = app.writeJSON(w, http.StatusCreated, contribution, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
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

	contribution, err := app.models.Contributions.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, contribution, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// @Summary		Update an infinity bottle contribution by ID
// @Description	Update all information about an infinity bottle contribution by ID
// @Param          id		path	int	true	"ID"
// @Param      Contribution	body	ContributionPost	true	"Update contribution to an infinity bottle"
// @Tags			contribution
// @Produce		json
// @Success		200	{object}    data.Contribution
// @Failure		404	{object}    ErrorMessage
// @Failure		500	{object}    ErrorMessage
// @Router			/v1/contributions/{id} [put]
func (app *application) updateContributionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	contribution, err := app.models.Contributions.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	input := ContributionPost{}
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	contribution.InfinityBottleID = input.InfinityBottleID
	contribution.Amount = input.Amount
	contribution.BrandName = input.BrandName
	contribution.Tags = input.Tags

	v := validator.New()

	if data.ValidateContribution(v, contribution); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Contributions.Update(contribution)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, contribution, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// @Summary		Delete an infinity bottle contribution by ID
// @Description	Delete an infinity bottle contribution by ID
// @Param          id		path	int	true	"ID"
// @Tags			contribution
// @Produce		json
// @Success		200
// @Failure		404	{object}    ErrorMessage
// @Failure		500	{object}    ErrorMessage
// @Router			/v1/contributions/{id} [delete]
func (app *application) deleteContributionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Contributions.Delete(id)
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
