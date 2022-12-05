package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/server-practice/pkg/models"
)

func (app *application) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	return
}

func (app *application) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
	return
}

func (app *application) NotFound(w http.ResponseWriter) {
	app.ClientError(w, http.StatusNotFound)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, templateData *TemplateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.ServerError(w, fmt.Errorf("The template %s does not exist", name))
		return
	}

	buf := new(bytes.Buffer)

	err := ts.Execute(buf, app.AddDefaultData(templateData, r))
	if err != nil {
		app.ServerError(w, err)
		return
	}

	buf.WriteTo(w)

}

func (app *application) AddDefaultData(td *TemplateData, r *http.Request) *TemplateData {
	if td == nil {
		td = &TemplateData{}
	}

	td.AuthenticatedUser = app.authenticatedUser(r)

	td.CurrentYear = time.Now().Year()
	td.Flash = app.session.PopString(r, "flash")
	return td
}

func (app *application) IsAuthenticated(r *http.Request) int {
	return app.session.GetInt(r, "userID")
}

func (app *application) authenticatedUser(r *http.Request) *models.User {
	user, err := r.Context().Value(context_key).(*models.User)
	if !err {
		return nil
	}
	return user
}
