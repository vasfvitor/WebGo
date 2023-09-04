package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vasfvitor/WebGo/controllers"
	"github.com/vasfvitor/WebGo/templates"
	"github.com/vasfvitor/WebGo/views"
	//"github.com/go-chi/chi/v5/middleware"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	w.Write([]byte(fmt.Sprintf("Hello %s", name)))
}

func main() {
	r := chi.NewRouter()

	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	// dynamic
	r.Get("/hello/{name}", getHello)

	r.NotFound(
		func(w http.ResponseWriter, r *http.Request) { http.Error(w, "Page not found", http.StatusNotFound) })
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
