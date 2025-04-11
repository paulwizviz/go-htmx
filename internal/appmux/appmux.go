package appmux

import (
	"embed"
	"net/http"
	"text/template"
)

//go:embed static/* templates/*
var web embed.FS

func staticHandler(rw http.ResponseWriter, r *http.Request) {
	staticContent, err := web.ReadFile("static/htmx.min.js")
	if err != nil {
		http.NotFound(rw, r)
		return
	}
	rw.Header().Set("Content-Type", "application/javascript")
	rw.Write(staticContent)
}

func homeHandler(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("base").ParseFS(web, "templates/*.html"))
	data := map[string]string{"Title": "Home", "Welcome": "Hello, world!"}
	err := t.ExecuteTemplate(rw, "homePage", data)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func clickedHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	rw.Write([]byte(`<strong>Button clicked!</strong>`))
}

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", homeHandler)
	mux.HandleFunc("GET /static/htmx.min.js", staticHandler)
	mux.HandleFunc("GET /clicked", clickedHandler)
	return mux
}
