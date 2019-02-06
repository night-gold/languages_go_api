package main

import (
	"net/http"
)

// AdminHomePage is the Home page of the admin part of the site.
func AdminHomePage(w http.ResponseWriter, r *http.Request) {
	message := "voilà l'admin, voilà l'admin"
	w.Write([]byte(message))
}

// AdminGetLanguages gives the listing of all languages.
func AdminGetLanguages(w http.ResponseWriter, r *http.Request) {
	message := "Donne la liste des languages dans la langue"
	w.Write([]byte(message))
}

// AdminGetLanguage gives the form to update a language
func AdminGetLanguage(w http.ResponseWriter, r *http.Request) {
	message := "Donne le formulaire d'update d'une langue"
	w.Write([]byte(message))
}

// AdminPostLanguage updates the languages in the languages.
func AdminPostLanguage(w http.ResponseWriter, r *http.Request) {
	message := "Donne le formulaire d'update d'une langue"
	w.Write([]byte(message))
}

// AdminDeleteLanguage Deletete the language in a specific language.
func AdminDeleteLanguage(w http.ResponseWriter, r *http.Request) {
	message := "delete le nom d'une langue pour une langue spécifique"
	w.Write([]byte(message))
}
