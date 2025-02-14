package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// HomeHandler is used to render the homepage
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/base.html", "templates/order_form.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error loading templates: %v", err), http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "base", nil) // Executes the "base" template
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
		return
	}
}

// OrderFormHandler dynamically renders the form based on order type
func OrderFormHandler(w http.ResponseWriter, r *http.Request) {
	orderType := r.URL.Query().Get("orderType")
	log.Printf("Received orderType: %s", orderType) // Debugging log

	var templateFile string
	if orderType == "local" {
		templateFile = "templates/local_form.html"
	} else if orderType == "global" {
		templateFile = "templates/global_form.html"
	} else {
		w.Write([]byte("<div></div>")) // Empty div for no selection
		return
	}

	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error loading template: %v", err), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
	}
}

func main() {
	mux := http.NewServeMux()

	// Serve static files (CSS, JS, etc.)
	fileServer := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Routes
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/order-form", OrderFormHandler)

	log.Println("Server running on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
