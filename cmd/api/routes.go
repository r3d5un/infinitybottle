package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/swaggo/http-swagger"
	_ "infinitybottle.islandwind.me/docs"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/infinitybottles", app.listInfinityBottlesHandler)
	router.HandlerFunc(http.MethodPost, "/v1/infinitybottles", app.createInfinityBottleHandler)
	router.HandlerFunc(http.MethodGet, "/v1/infinitybottles/:id", app.getInfinityBottleHandler)
	router.HandlerFunc(http.MethodPut, "/v1/infinitybottles/:id", app.updateInfinityBottleHandler)
	router.HandlerFunc(
		http.MethodDelete,
		"/v1/infinitybottles/:id",
		app.deleteInfinityBottleHandler,
	)

	router.HandlerFunc(http.MethodGet, "/v1/contributions", app.listContributionsHandler)
	router.HandlerFunc(http.MethodPost, "/v1/contributions", app.createContributionHandler)
	router.HandlerFunc(http.MethodGet, "/v1/contributions/:id", app.getContributionHandler)
	router.HandlerFunc(http.MethodPut, "/v1/contributions/:id", app.updateContributionHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/contributions/:id", app.deleteContributionHandler)

	router.HandlerFunc(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)

	return router
}
