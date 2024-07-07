package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"text/template"

	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.log.Error(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.log.Error(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func (app *application) getAPIContent(ctx context.Context, url string, templateData interface{}) error {
	_, span := app.tracer.Start(ctx, "getAPIContent")
	defer span.End()

	span.SetAttributes(semconv.HTTPURL(url))
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	json.Unmarshal(bodyBytes, templateData)
	return nil
}

func (app *application) static(dir string) http.Handler {
	dirCleaned := filepath.Clean(dir)
	return http.StripPrefix("/static/", http.FileServer(http.Dir(dirCleaned)))
}
