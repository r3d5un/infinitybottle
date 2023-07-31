package main

import (
	"fmt"
	"net/http"
)

func (app *application) listBrandBottlesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "list brand bottles")
}

func (app *application) createBrandBottleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create brand bottle")
}

func (app *application) getBrandBottleHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "get brand bottle %d\n", id)
}
