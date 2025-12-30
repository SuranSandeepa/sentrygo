package handlers

import (
	"html/template"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/SuranSandeepa/sentrygo/internal/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

// This tells Go where to find your HTML files
var tmpl = template.Must(template.ParseGlob("web/templates/*.html"))

func Dashboard(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Fetch data from Postgres
		services, err := db.GetAllServices(pool)
		if err != nil {
			http.Error(w, "Failed to fetch services", http.StatusInternalServerError)
			return
		}

		// Inject the data into the HTML template
		tmpl.ExecuteTemplate(w, "index.html", services)
	}
}

func AddService(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. Get values from the form
		name := r.FormValue("name")
		url := r.FormValue("url")

		// 2. Save to Database
		err := db.CreateService(pool, name, url)
		if err != nil {
			http.Error(w, "Failed to save service", http.StatusInternalServerError)
			return
		}

		// 3. HTMX trick: Redirect back to the dashboard
		w.Header().Set("HX-Redirect", "/")
	}
}

func DeleteService(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id") // Get the ID from the URL
		db.DeleteService(pool, id)
		
		// Sending nothing back tells HTMX to remove the element from the UI
		w.WriteHeader(http.StatusOK)
	}
}