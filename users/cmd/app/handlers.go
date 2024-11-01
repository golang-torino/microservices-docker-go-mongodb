package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mmorejon/microservices-docker-go-mongodb/users/pkg/models"
)

func (app *application) all(w http.ResponseWriter, r *http.Request) {
	ctx, span := app.tracer.Start(r.Context(), "db get all users")

	// Get all user stored
	users, err := app.users.All(ctx)
	if err != nil {
		app.serverError(w, err)
	}

	span.End()

	// Convert user list into json encoding
	b, err := json.Marshal(users)
	if err != nil {
		app.serverError(w, err)
	}

	app.log.Info("Users have been listed", "count", len(users))
	app.mets.foundUsers.Add(r.Context(), int64(len(users)))

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findByID(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	ctx, span := app.tracer.Start(r.Context(), "get user by id")

	// Find user by id
	m, err := app.users.FindByID(ctx, id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.log.Info("User not found", "id", id)
			return
		}
		// Any other error will send an internal server error
		app.serverError(w, err)
	}

	span.End()

	// Convert user to json encoding
	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.log.Info("User has been found")
	app.mets.foundUsers.Add(r.Context(), 1)

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) insert(w http.ResponseWriter, r *http.Request) {
	// Define user model
	var u models.User
	// Get request information
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		app.serverError(w, err)
	}

	ctx, span := app.tracer.Start(r.Context(), "insert user")

	// Insert new user
	insertResult, err := app.users.Insert(ctx, u)
	if err != nil {
		app.serverError(w, err)
	}

	span.End()

	app.log.Info("New user have been created", "id", insertResult.InsertedID)

	app.mets.createdUsers.Add(r.Context(), 1)
}

func (app *application) delete(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete user by id
	deleteResult, err := app.users.Delete(r.Context(), id)
	if err != nil {
		app.serverError(w, err)
	}

	app.log.Info("users deleted", "count", deleteResult.DeletedCount)

	app.mets.deletedUsers.Add(r.Context(), 1)
}
