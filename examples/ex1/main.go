package main

import (
	"embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
)

//go:embed templates/* static/*
var content embed.FS

type PageData struct {
	Title   string
	Heading string
	Content string
}

func homeHDL(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.New("base").ParseFS(content, "templates/*.html")
	if err != nil {
		http.Error(w, "Template creation error", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title:   "Go Templates with Multiple Files",
		Heading: "Welcome to the Go Template Example",
		Content: "This example demonstrates using multiple templates.",
	}

	// Execute the base template (index.html) and pass in the PageData
	err = tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
		return
	}
}

func dataHDL(w http.ResponseWriter, r *http.Request) {
	b, _ := io.ReadAll(r.Body)
	w.Write(b)
}

func testHDL(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello Test"))
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	filePath := "static" + r.URL.Path[len("/static"):]
	fileContent, err := content.ReadFile(filePath)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Set appropriate content type
	w.Header().Set("Content-Type", "application/javascript")
	w.Write(fileContent)
}

func tab1HDL(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<p style="color:DodgerBlue;" >Hello Tab 1</p>`))
}

func tab2HDL(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<p style="color:MediumSeaGreen;">Hello Tab 2</p>`))
}

func tab3HDL(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<p style="color:Tomato;">Hello Tab 3</p>`))
}

func main() {
	port := "8080"
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", homeHDL)
	mux.HandleFunc("GET /static/htmx.min.js", staticHandler)
	mux.HandleFunc("GET /test", testHDL)
	mux.HandleFunc("GET /tab1", tab1HDL)
	mux.HandleFunc("GET /tab2", tab2HDL)
	mux.HandleFunc("GET /tab3", tab3HDL)
	mux.HandleFunc("POST /data", dataHDL)
	log.Printf("Starting web %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), mux))
}
