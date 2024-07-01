package main

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func (app *application) routes() *mux.Router {
	handleAndCountVisits := func(r *mux.Router, p string, h http.HandlerFunc) {
		r.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
			app.measures.requests.Add(context.Background(), 1,
				metric.WithAttributes(semconv.HTTPRoute(p)),
			)
			app.errorLog.Print("HERE")
			h(w, r)
		})
	}

	// Register handler functions.
	r := mux.NewRouter()
	r.Use(otelmux.Middleware("my-server"))

	handleAndCountVisits(r, "/", app.home)
	handleAndCountVisits(r, "/users/list", app.usersList)
	handleAndCountVisits(r, "/users/view/{id}", app.usersView)
	handleAndCountVisits(r, "/movies/list", app.moviesList)
	handleAndCountVisits(r, "/movies/view/{id}", app.moviesView)
	handleAndCountVisits(r, "/showtimes/list", app.showtimesList)
	handleAndCountVisits(r, "/showtimes/view/{id}", app.showtimesView)
	handleAndCountVisits(r, "/bookings/list", app.bookingsList)
	handleAndCountVisits(r, "/bookings/view/{id}", app.bookingsView)

	// This will serve files under http://localhost:8000/static/<filename>
	r.PathPrefix("/static/").Handler(app.static("./ui/static/"))
	return r
}
