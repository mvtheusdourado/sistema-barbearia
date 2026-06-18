package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mvtheusdourado/sistema-barbearia/internal/models"
	"github.com/mvtheusdourado/sistema-barbearia/internal/service"
)

type AgendamentoHandler struct {
	service *service.AgendamentoService
}

func NewAgendamentoHandler(service *service.AgendamentoService) *AgendamentoHandler {
	return &AgendamentoHandler{service: service}
}

func (h *AgendamentoHandler) CriarAgendamento(w http.ResponseWriter, r *http.Request) {
	var ag models.Agendamento

	err := json.NewDecoder(r.Body).Decode(&ag)
	if err != nil {
		http.Error(w, "Dados inválidos no corpo da requisição", http.StatusBadRequest)
		return
	}

	err = h.service.CriarAgendamento(r.Context(), ag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]string{"mensagem": "Agendamento criado com sucesso!"})
}
