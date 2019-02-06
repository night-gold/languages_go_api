package main

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var db *sql.DB

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	message := "Welcome on this test API"
	w.Write([]byte(message))
}

func newRouter() *chi.Mux {
	r := chi.NewRouter()
	r.HandleFunc("/", homePage)
	r.HandleFunc("/code", LanguageCode)
	r.Mount("/{lang}/admin", adminRouter())
	r.Mount("/{lang}", languagesRouter())
	return r
}

func handleRequests() {
	var err error
	r := newRouter()
	r.Use(middleware.Recoverer)

	db, err = sql.Open("mysql", "langage:Password123@/languages?charset=utf8mb4")
	checkErr(err)
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}

	defer db.Close()
}

func main() {
	handleRequests()
}