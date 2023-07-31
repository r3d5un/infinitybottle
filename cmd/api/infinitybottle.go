package main

import (
	"fmt"
	"net/http"
)

func (app *application) listInfinityBottlesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "list infinity bottles")
}

func (app *application) createInfinityBottleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create infinity bottle")
}

func (app *application) getInfinityBottleHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "get infinity bottle %d\n", id)
}
