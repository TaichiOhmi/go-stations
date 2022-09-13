package handler

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/TechBowl-japan/go-stations/model"
)

// A HealthzHandler implements health check endpoint.
type HealthzHandler struct{
	Endpoint string
}

// NewHealthzHandler returns HealthzHandler based http.Handler.
func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{Endpoint: "/healthz"}
}

// ServeHTTP implements http.Handler interface.
// JSONデータを読み込み、その結果を構造体HealthzResponseに格納する
func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	HealthzResponse := &model.HealthzResponse{Message: "OK"}
	if err := json.NewEncoder(w).Encode(HealthzResponse); err != nil {
		log.Println(err)
	}
}
