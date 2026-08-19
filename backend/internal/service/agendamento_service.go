package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/mvtheusdourado/sistema-barbearia/internal/models"
	"github.com/mvtheusdourado/sistema-barbearia/internal/repository"
)

type AgendamentoService struct {
	repo *repository.AgendamentoRepository
}

func NewAgendamentoService(repo *repository.AgendamentoRepository) *AgendamentoService {
	return &AgendamentoService{repo: repo}
}

func (s *AgendamentoService) CriarAgendamento(ctx context.Context, ag models.Agendamento) error {
	if ag.DataHora.Before(time.Now()) {
		return errors.New("não é possível agendar em uma data no passado")
	}

	ocupado, err := s.repo.ExisteHorarioOcupado(ctx, ag.BarbeiroID, ag.DataHora)
	if err != nil {
		return err
	}
	if ocupado {
		return errors.New("este horário já está ocupado para este barbeiro")
	}

	return s.repo.Criar(ctx, ag)
}
func (s *AgendamentoService) CancelarAgendamento(ctx context.Context, id int) error {
	cancelado, err := s.repo.Cancelar(ctx, id)
	if err != nil {
		return err
	}
	if !cancelado {
		return errors.New("agendamento não encontrado")
	}
	return nil
}
func (s *AgendamentoService) ListarHorariosDisponiveis(ctx context.Context, barbeiroID int, dia time.Time) ([]string, error) {
	ocupados, err := s.repo.HorariosOcupados(ctx, barbeiroID, dia)
	if err != nil {
		return nil, err
	}

	horasOcupadas := make(map[int]bool)
	for _, o := range ocupados {
		horasOcupadas[o.Hour()] = true
	}

	disponiveis := []string{}
	for hora := 9; hora < 18; hora++ {
		if !horasOcupadas[hora] {
			disponiveis = append(disponiveis, fmt.Sprintf("%02d:00", hora))
		}
	}
	return disponiveis, nil
}
func (s *AgendamentoService) ListarAgendamentos(ctx context.Context, clienteID int) ([]models.Agendamento, error) {
	return s.repo.ListarPorCliente(ctx, clienteID)
}
