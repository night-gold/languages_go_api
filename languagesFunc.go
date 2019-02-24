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

// LanguageCode Lists all the codes available.
func LanguageCode(w http.ResponseWriter, r *http.Request) {
	var code string
	var codes []Code
	query, err := db.Query("SELECT * FROM languages;")
	checkErr(err)
	defer query.Close()

	for query.Next() {
		err := query.Scan(&code)
		checkErr(err)
		codes = append(codes, Code{Code: code})
	}

	codesMarshalled, err := json.Marshal(codes)
	checkErr(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(codesMarshalled))
}

// LanguageName Gives the code for the current language
func LanguageName(w http.ResponseWriter, r *http.Request) {
	lang := chi.URLParam(r, "lang")
	var name string
	var resname LangName
	rname, err := db.Prepare("SELECT name FROM languages_name WHERE fk_language LIKE ? AND fk_language_name LIKE fk_language;")
	checkErr(err)
	defer rname.Close()
	err = rname.QueryRow(lang).Scan(&name)
	checkErr(err)

	resname = LangName{Name: name}
	codesMarshalled, err := json.Marshal(resname)
	checkErr(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(codesMarshalled)
}
