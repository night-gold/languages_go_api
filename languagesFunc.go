package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

// LanguagesList listing des languages pour le code
func LanguagesList(w http.ResponseWriter, r *http.Request) {
	lang := chi.URLParam(r, "lang")
	query, err := db.Query("SELECT languages_name.fk_language_name,languages_name.name FROM languages INNER JOIN languages_name WHERE languages.code LIKE languages_name.fk_language AND languages.code like ?;", lang)
	checkErr(err)
	var code string
	var name string
	var languages []Language

	defer query.Close()

	for query.Next() {
		err := query.Scan(&code, &name)
		checkErr(err)
		languages = append(languages, Language{Code: code, Name: name})
	}
	codesMarshalled, err := json.Marshal(languages)
	checkErr(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(codesMarshalled)
}

// LanguageCode Gives the code for the current language
func LanguageCode(w http.ResponseWriter, r *http.Request) {
	message := "Current language code"
	w.Write([]byte(message))
}

// LanguageName Gives the code for the current language
func LanguageName(w http.ResponseWriter, r *http.Request) {
	lang := chi.URLParam(r, "lang")
	var name string
	var resname LangName
	rname, err := db.Prepare("SELECT name FROM languages_name WHERE fk_language LIKE ? AND fk_language_name LIKE fk_language;")
	checkErr(err)
	err = rname.QueryRow(lang).Scan(&name)
	checkErr(err)

	resname = LangName{Name: name}
	codesMarshalled, err := json.Marshal(resname)
	checkErr(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(codesMarshalled)
}

// LanguageInfo Gives the current language information, code and names
func LanguageInfo(w http.ResponseWriter, r *http.Request) {
	message := "Current language informations"
	w.Write([]byte(message))
}
