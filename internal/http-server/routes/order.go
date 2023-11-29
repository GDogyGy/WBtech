package routes

import (
	"WbTech0/internal/lib/filterOrder"
	"WbTech0/internal/model"
	"WbTech0/internal/storage/postgres"
	"html/template"
	"log"
	"log/slog"
	"net/http"
)

func OrderRoutes(log *slog.Logger, storage *postgres.Storage, data *[]model.Order) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		UID := r.FormValue("uid_id")
		s := filterOrder.GetOrderById(*data, UID)
		tplView(w, s)
	}
}

func tplView(w http.ResponseWriter, data model.Order) {
	var tpl *template.Template
	tpl = template.Must(template.ParseGlob("internal/http-server/templates/*"))
	if err := tpl.ExecuteTemplate(w, "index.gohtml", data); err != nil {
		log.Printf("Template error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
