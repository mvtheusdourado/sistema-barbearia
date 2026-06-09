package service

import (
	"context"

	"github.com/mvtheusdourado/sistema-barbearia/internal/models"
	"github.com/mvtheusdourado/sistema-barbearia/internal/repository"
)

type ClienteService struct {
	repo *repository.ClienteRepository
}

func NewClienteService(repo *repository.ClienteRepository) *ClienteService {
	return &ClienteService{repo: repo}
}

func (s *ClienteService) ListarClientes(ctx context.Context) ([]models.Cliente, error) {
	return s.repo.ListarTodos(ctx)
}
