package handlers

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/base.html", "templates/order_form.html")
	tmpl.ExecuteTemplate(w, "base", nil)
}

func OrderFormHandler(w http.ResponseWriter, r *http.Request) {
	orderType := r.URL.Query().Get("orderType")

	var templateFile string
	if orderType == "local" {
		templateFile = "templates/local_form.html"
	} else if orderType == "global" {
		templateFile = "templates/global_form.html"
	} else {
		w.Write([]byte("<div></div>")) // Empty div for no selection
		return
	}

	tmpl, _ := template.ParseFiles(templateFile)
	tmpl.Execute(w, nil)
}
