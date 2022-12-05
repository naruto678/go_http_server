package main

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {

	dynamicMiddleWare := alice.New(app.session.Enable, app.authenticate)
	standardMiddleware := alice.New(app.recoverPanic, app.logRequests, SecureHeaders)

	mux := pat.New()
	mux.Get("/", dynamicMiddleWare.Then(http.HandlerFunc(app.HomeHandler)))
	mux.Post("/snippet/create", dynamicMiddleWare.Append(app.RequireAuthenticatedUser).Then(http.HandlerFunc(app.CreateSnippet)))
	mux.Get("/snippet/create", dynamicMiddleWare.Append(app.RequireAuthenticatedUser).Then(http.HandlerFunc(app.CreateSnippetForm)))
	mux.Get("/snippet/:id", dynamicMiddleWare.Then(http.HandlerFunc(app.ShowSnippet)))
	mux.Post("/user/login", dynamicMiddleWare.Then(http.HandlerFunc(app.UserLogin)))
	mux.Post("/user/logout", dynamicMiddleWare.Append(app.RequireAuthenticatedUser).Then(http.HandlerFunc(app.UserLogout)))
	mux.Get("/user/login", dynamicMiddleWare.Then(http.HandlerFunc(app.UserLoginForm)))
	mux.Get("/user/signup", dynamicMiddleWare.Then(http.HandlerFunc(app.UserSignupForm)))
	mux.Post("/user/signup", dynamicMiddleWare.Then(http.HandlerFunc(app.UserSignup)))
	mux.Get("/ping", http.HandlerFunc(app.PingHandler))
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	return standardMiddleware.Then(mux)
}
