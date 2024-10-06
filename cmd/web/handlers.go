package main

import (
	"html/template"
	"net/http"
)

func (app *app) home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./assets/templates/index.html",
		"./assets/templates/nav.html",
		"./assets/templates/home.html",
	}

	expenses, err := app.expenses.GetExpenses()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templ, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = templ.ExecuteTemplate(w, "base", expenses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
