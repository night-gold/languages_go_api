package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var db *sql.DB

func checkErr(err error) {
	if err != nil {
		fmt.Errorf("pkg: %v", err)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	message := "Welcome on this test API"
	w.Write([]byte(message))
}

func newRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	//Statics file server
	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "statics")
	FileServer(r, "/statics", http.Dir(filesDir))
	//Other routes
	r.HandleFunc("/", homePage)
	r.HandleFunc("/code", LanguageCode)
	r.Mount("/{lang}/admin", adminRouter())
	r.Mount("/{lang}", languagesRouter())
	return r
}

func handleRequests() {
	var err error
	r := newRouter()

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

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
