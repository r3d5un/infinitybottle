package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/swaggo/http-swagger"
	_ "infinitybottle.islandwind.me/docs"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)

	return router
}
