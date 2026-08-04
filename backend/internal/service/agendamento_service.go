package service

import (
	"context"
	"errors"
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
