package router

import (
	"database/sql"
	"net/http"
	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/service"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()
	s := service.NewTODOService(todoDB)
	h := handler.NewTODOHandler(s)
	mux.HandleFunc("/todos", h.ServeHTTP)
	return mux
}
