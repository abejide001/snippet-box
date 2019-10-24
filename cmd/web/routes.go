package main

import (
	"net/http"

	"github.com/bmizerany/pat" // New import
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/snippet/create", http.HandlerFunc(app.createSnippetForm))
	mux.Post("/snippet/create", http.HandlerFunc(app.createSnippet))
	mux.Get("/snippet/:id", http.HandlerFunc(app.showSnippet))
	mux.Get("/user/signup", http.HandlerFunc(app.signupUserForm))
	mux.Post("/user/signup", http.HandlerFunc(app.signupUser))
	mux.Get("/user/login", http.HandlerFunc(app.loginUserForm))
	mux.Post("/user/login", http.HandlerFunc(app.loginUser))
	mux.Post("/user/logout", http.HandlerFunc(app.logoutUser))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
