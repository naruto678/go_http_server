package main

import (
	"fmt"
	"github.com/server-practice/pkg/forms"
	"github.com/server-practice/pkg/models"
	"net/http"
	"strconv"
)

func (app *application) CreateSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.html", &TemplateData{
		Form: forms.New(nil),
	})
}

func (app *application) UserLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.ClientError(w, http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)

	id, err := app.users.Authenticate(form.Get("email"), form.Get("password"))
	if err == models.ErrInvalidCredential {
		form.Errors.Add("generic", "Email or password is incorrect")
		app.render(w, r, "login.page.html", &TemplateData{Form: form})
		return
	} else if err != nil {
		app.ServerError(w, err)
		return
	}
	app.session.Put(r, "userID", id)
	http.Redirect(w, r, "/snippet/create", http.StatusSeeOther)

}

func (app *application) UserLogout(w http.ResponseWriter, r *http.Request) {
	app.session.Remove(r, "userID")
	app.session.Put(r, "flash", "You've been successfully logged out")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) UserLoginForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.html", &TemplateData{Form: forms.New(nil)})
}

func (app *application) UserSignupForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "signup.page.html", &TemplateData{
		Form: forms.New(nil),
	})
}

func (app *application) UserSignup(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.ClientError(w, http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)
	form.MatchesPattern("email", forms.EmailRgx)
	form.MinLength("password", 2)
	form.Required("name", "email", "password")

	if !form.IsValid() {
		app.render(w, r, "signup.page.html", &TemplateData{Form: form})
		return
	}

	err = app.users.InsertUser(form.Get("name"), form.Get("email"), form.Get("password"))
	if err == models.ErrDuplicateMail {
		form.Errors.Add("Email", "Email already in use")
		app.render(w, r, "signup.page.html", &TemplateData{Form: form})
	} else if err != nil {
		app.ServerError(w, err)
		return
	}
	app.session.Put(r, "flash", "Your signup was successful , Please log in")

	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (app *application) HomeHandler(w http.ResponseWriter, r *http.Request) {

	snippets, err := app.snippets.GetAll()
	if err != nil {
		app.ServerError(w, err)
		return
	}

	templateData := &TemplateData{
		Snippets: snippets,
	}

	app.render(w, r, "home.page.html", templateData)

}

func (app *application) ShowSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}
	snippet, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.NotFound(w)
		return
	}

	templateData := &TemplateData{
		Snippet: snippet,
	}

	app.render(w, r, "show.page.html", templateData)

}

func (app *application) CreateSnippet(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		app.ClientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLen("title", 100)
	form.PermittedValues("expires", "365", "7", "1")

	if !form.IsValid() {
		app.render(w, r, "create.page.html", &TemplateData{
			Form: form,
		})
		return
	}

	app.render(w, r, "create.page.html", &TemplateData{
		Form: forms.New(nil),
	})

	id, err := app.snippets.Insert(form.Get("title"), form.Get("content"), form.Get("expires"))
	if err != nil {
		app.ServerError(w, err)
		return
	}

	app.session.Put(r, "flash", "Snippet successfully created")
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}

func (app *application) GetLatestSnippets(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		app.ClientError(w, http.StatusMethodNotAllowed)
	}

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.ServerError(w, err)
		return
	}
	for _, snippet := range snippets {
		fmt.Fprintf(w, "%v\n", snippet)
	}
}

func (app *application) PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
