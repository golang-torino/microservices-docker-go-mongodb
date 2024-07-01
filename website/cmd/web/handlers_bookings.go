package main

import (
	"context"
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/mmorejon/microservices-docker-go-mongodb/bookings/pkg/models"
	modelsShowTime "github.com/mmorejon/microservices-docker-go-mongodb/showtimes/pkg/models"
	modelsUser "github.com/mmorejon/microservices-docker-go-mongodb/users/pkg/models"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type bookingTemplateData struct {
	Booking      models.Booking
	Bookings     []models.Booking
	BookingData  bookingData
	BookingsData []bookingData
}

type bookingData struct {
	ID           string
	UserFullName string
	ShowTimeDate string
}

func (app *application) loadBookingData(ctx context.Context, btd *bookingTemplateData, isList bool) {
	ctx, span := app.tracer.Start(ctx, "load booking data")
	defer span.End()

	// Clean booking data
	btd.BookingsData = []bookingData{}
	btd.BookingData = bookingData{}

	span.SetAttributes(attribute.Bool("list", isList))

	// Load booking data
	if isList {
		for _, b := range btd.Bookings {
			// Load user data
			userURL := fmt.Sprintf("%s/%s", app.apis.users, b.UserID)
			var user modelsUser.User
			err := app.getAPIContent(ctx, userURL, &user)
			if err != nil {
				app.errorLog.Println(err.Error())
			}

			// Load showtime data
			showtimeURL := fmt.Sprintf("%s/%s", app.apis.showtimes, b.ShowtimeID)
			var showtime modelsShowTime.ShowTime
			err = app.getAPIContent(ctx, showtimeURL, &showtime)
			if err != nil {
				app.errorLog.Println(err.Error())
			}

			bookingData := bookingData{
				ID:           b.ID.Hex(),
				UserFullName: fmt.Sprintf("%s %s", user.Name, user.LastName),
				ShowTimeDate: showtime.Date,
			}
			btd.BookingsData = append(btd.BookingsData, bookingData)
			app.infoLog.Println(b.UserID)
		}
	} else {
		b := btd.Booking

		// Load user data
		userURL := fmt.Sprintf("%s/%s", app.apis.users, b.UserID)
		var user modelsUser.User
		err := app.getAPIContent(ctx, userURL, &user)
		if err != nil {
			app.errorLog.Println(err.Error())
		}

		// Load showtime data
		showtimeURL := fmt.Sprintf("%s/%s", app.apis.showtimes, b.ShowtimeID)
		var showtime modelsShowTime.ShowTime

		err = app.getAPIContent(ctx, showtimeURL, &showtime)
		if err != nil {
			app.errorLog.Println(err.Error())
		}

		btd.BookingData = bookingData{
			ID:           b.ID.Hex(),
			UserFullName: fmt.Sprintf("%s %s", user.Name, user.LastName),
			ShowTimeDate: showtime.Date,
		}
	}
}

func (app *application) bookingsList(w http.ResponseWriter, r *http.Request) {
	// Get bookings list from API
	var td bookingTemplateData
	app.infoLog.Println("Calling bookings API...")

	err := app.getAPIContent(r.Context(), app.apis.bookings, &td.Bookings)
	if err != nil {
		app.errorLog.Println(err.Error())
	}
	app.infoLog.Println(td.Bookings)
	app.infoLog.Println(td)

	app.loadBookingData(r.Context(), &td, true)

	// Load template files
	files := []string{
		"./ui/html/bookings/list.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	if err = renderTemplates(app.tracer, r.Context(), files, td, w); err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func (app *application) bookingsView(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	bookingID := vars["id"]

	// Get bookings list from API
	var td bookingTemplateData
	app.infoLog.Println("Calling bookings API...")
	url := fmt.Sprintf("%s/%s", app.apis.bookings, bookingID)

	err := app.getAPIContent(r.Context(), url, &td.Booking)
	if err != nil {
		app.errorLog.Println(err.Error())
	}
	app.infoLog.Println(td.Booking)
	app.infoLog.Println(url)

	app.loadBookingData(r.Context(), &td, false)

	// Load template files
	files := []string{
		"./ui/html/bookings/view.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	if err = renderTemplates(app.tracer, r.Context(), files, td, w); err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func renderTemplates(t trace.Tracer, ctx context.Context, files []string, td bookingTemplateData, w http.ResponseWriter) error {
	_, span := t.Start(ctx, "render template")
	defer span.End()

	ts, err := template.ParseFiles(files...)
	if err != nil {
		return err
	}

	err = ts.Execute(w, td)
	if err != nil {
		return err
	}

	return nil
}
