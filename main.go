package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/vasfvitor/WebGo/views"
	//"github.com/go-chi/chi/v5/middleware"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	t, err := views.Parse(filepath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "Error parsing the template.", http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath)

}

func getArticle(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	w.Write([]byte(fmt.Sprintf("Hello %s", name)))
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		http.Error(w, "Page not found", http.StatusNotFound)
// 	}
// }

// type Router struct{}

// func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		http.Error(w, "Page not found", http.StatusNotFound)
// 	}
// }

func main() {

	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/articles/{name}", getArticle)
	r.NotFound(
		func(w http.ResponseWriter, r *http.Request) { http.Error(w, "Page not found", http.StatusNotFound) })
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
