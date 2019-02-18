package main

import (
	"encoding/json"
	"net/http"
)

// LanguagesList listing des languages pour la langue
func LanguagesList(w http.ResponseWriter, r *http.Request) {
	//lang := chi.URLParam(r, "lang")
	query, err := db.Query("SELECT * FROM languages;")
	checkErr(err)
	var code string
	var codes []Code

	defer query.Close()

	for query.Next() {
		err := query.Scan(&code)
		checkErr(err)
		codes = append(codes, Code{Code: code})
	}
	codesMarshalled, err := json.Marshal(codes)
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
	message := "Current language name"
	w.Write([]byte(message))
}

// LanguageInfo Gives the current language information, code and names
func LanguageInfo(w http.ResponseWriter, r *http.Request) {
	message := "Current language informations"
	w.Write([]byte(message))
}
