package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mvtheusdourado/sistema-barbearia/internal/models"
)

type ClienteRepository struct {
	db *pgxpool.Pool
}

func NewClienteRepository(db *pgxpool.Pool) *ClienteRepository {
	return &ClienteRepository{db: db}
}

func (r *ClienteRepository) ListarTodos(ctx context.Context) ([]models.Cliente, error) {
	linhas, err := r.db.Query(ctx, "SELECT id, nome, telefone FROM clientes")
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var clientes []models.Cliente
	for linhas.Next() {
		var c models.Cliente
		err := linhas.Scan(&c.ID, &c.Nome, &c.Telefone)
		if err != nil {
			return nil, err
		}
		clientes = append(clientes, c)
	}
	return clientes, nil
}
