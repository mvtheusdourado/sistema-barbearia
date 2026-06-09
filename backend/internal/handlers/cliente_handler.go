package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mvtheusdourado/sistema-barbearia/internal/service"
)

type ClienteHandler struct {
	service *service.ClienteService
}

func NewClienteHandler(service *service.ClienteService) *ClienteHandler {
	return &ClienteHandler{service: service}
}

func (h *ClienteHandler) ListarClientes(w http.ResponseWriter, r *http.Request) {
	clientes, err := h.service.ListarClientes(r.Context())
	if err != nil {
		http.Error(w, "Erro ao buscar clientes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clientes)
}
