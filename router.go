package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func adminRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", AdminHomePage)
	r.Get("/languages", AdminGetLanguages)
	r.Get("/languages/{lang2}", AdminGetLanguage)
	r.Post("/languages/{lang2}", AdminPostLanguage)
	r.Delete("/languages/{lang2}", AdminDeleteLanguage)
	return r
}

func languagesRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", LanguagesList)
	r.Get("/info", LanguageInfo)
	r.Get("/code", LanguageCode)
	r.Get("/name", LanguageName)
	return r
}
