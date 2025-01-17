package handler

import (
	"context"
	"net/http"
	"encoding/json"
	"log"
	"github.com/TechBowl-japan/go-stations/model"
	"github.com/TechBowl-japan/go-stations/service"
)

// A TODOHandler implements handling REST endpoints.
type TODOHandler struct {
	svc *service.TODOService
}

// NewTODOHandler returns TODOHandler based http.Handler.
func NewTODOHandler(svc *service.TODOService) *TODOHandler {
	return &TODOHandler{
		svc: svc,
	}
}

func (h *TODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		
		request := &model.CreateTODORequest{}
		// log.Print("request: ",request)
		if err := json.NewDecoder(r.Body).Decode(request); err != nil {
			log.Println(err)
			return
		}
		// log.Print("decoded request: ",request)
		if request.Subject == "" {
			w.WriteHeader(http.StatusBadRequest)
			return 
		}	
		todos, err := h.svc.CreateTODO(r.Context(), request.Subject, request.Description)
		// log.Printf("%T\n", todos)
		// log.Println("res",todos)
		if err != nil {
			log.Println(err)
			return
		}

		response := model.CreateTODOResponse{TODO:*todos}

		// log.Println(response)
		// log.Printf("%v, %T\n", response.TODO.ID,response.TODO.ID)

		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Println(err)
			return
		}

		// log.Printf("%v, %T\n", response.TODO.ID,response.TODO.ID)
		// log.Println(response)
	}
	
}

// Create handles the endpoint that creates the TODO.
func (h *TODOHandler) Create(ctx context.Context, req *model.CreateTODORequest) (*model.CreateTODOResponse, error) {
	_, _ = h.svc.CreateTODO(ctx, "", "")
	return &model.CreateTODOResponse{}, nil
}

// Read handles the endpoint that reads the TODOs.
func (h *TODOHandler) Read(ctx context.Context, req *model.ReadTODORequest) (*model.ReadTODOResponse, error) {
	_, _ = h.svc.ReadTODO(ctx, 0, 0)
	return &model.ReadTODOResponse{}, nil
}

// Update handles the endpoint that updates the TODO.
func (h *TODOHandler) Update(ctx context.Context, req *model.UpdateTODORequest) (*model.UpdateTODOResponse, error) {
	_, _ = h.svc.UpdateTODO(ctx, 0, "", "")
	return &model.UpdateTODOResponse{}, nil
}

// Delete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) Delete(ctx context.Context, req *model.DeleteTODORequest) (*model.DeleteTODOResponse, error) {
	_ = h.svc.DeleteTODO(ctx, nil)
	return &model.DeleteTODOResponse{}, nil
}
