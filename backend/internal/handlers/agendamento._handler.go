package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

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
func (h *AgendamentoHandler) CancelarAgendamento(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = h.service.CancelarAgendamento(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"mensagem": "Agendamento cancelado com sucesso!"})
}
func (h *AgendamentoHandler) ListarHorariosDisponiveis(w http.ResponseWriter, r *http.Request) {
	barbeiroID, err := strconv.Atoi(r.PathValue(("id")))
	if err != nil {
		http.Error(w, "ID do barbeiro inválido", http.StatusBadRequest)
		return
	}

	dataStr := r.URL.Query().Get("data")
	dia, err := time.Parse("2006-01-02", dataStr)
	if err != nil {
		http.Error(w, "Data inválida (use o formato AAAA-MM-DD)", http.StatusBadRequest)
		return
	}

	disponiveis, err := h.service.ListarHorariosDisponiveis(r.Context(), barbeiroID, dia)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(disponiveis)
}
func (h *AgendamentoHandler) ListarAgendamentos(w http.ResponseWriter, r *http.Request) {
	clienteID, err := strconv.Atoi(r.URL.Query().Get("cliente_id"))
	if err != nil {
		http.Error(w, "cliente_id inválido", http.StatusBadRequest)
		return
	}

	agendamentos, err := h.service.ListarAgendamentos(r.Context(), clienteID)
	if err != nil {
		http.Error(w, "Erro ao buscar agendamentos", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(agendamentos)
}
