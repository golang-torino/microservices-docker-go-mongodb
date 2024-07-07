package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/mmorejon/microservices-docker-go-mongodb/users/pkg/models"
)

type userTemplateData struct {
	User  models.User
	Users []models.User
}

func (app *application) usersList(w http.ResponseWriter, r *http.Request) {

	// Get users list from API
	var utd userTemplateData
	err := app.getAPIContent(r.Context(), app.apis.users, &utd.Users)
	if err != nil {
		app.errorLog.Println(err.Error())
	}
	app.log.Info("retrieved users", "users", utd.Users)

	// Load template files
	files := []string{
		"./ui/html/users/list.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, utd)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func (app *application) usersView(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	userID := vars["id"]

	// Get users list from API
	app.log.Info("Calling users API...")
	url := fmt.Sprintf("%s/%s", app.apis.users, userID)

	var utd userTemplateData
	app.getAPIContent(r.Context(), url, &utd.User)
	app.log.Info("retrieved user", "user", utd.User)

	// Load template files
	files := []string{
		"./ui/html/users/view.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, utd.User)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
