package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mvtheusdourado/sistema-barbearia/internal/models"
)

type AgendamentoRepository struct {
	db *pgxpool.Pool
}

func NewAgendamentoRepository(db *pgxpool.Pool) *AgendamentoRepository {
	return &AgendamentoRepository{db: db}
}

func (r *AgendamentoRepository) Criar(ctx context.Context, ag models.Agendamento) error {
	_, err := r.db.Exec(ctx, "INSERT INTO agendamentos (cliente_id, barbeiro_id, data_hora) VALUES ($1, $2, $3)", ag.ClienteID, ag.BarbeiroID, ag.DataHora)
	return err
}

func (r *AgendamentoRepository) ExisteHorarioOcupado(ctx context.Context, barbeiroID int, dataHora time.Time) (bool, error) {
	var existe bool
	err := r.db.QueryRow(ctx, "SELECT EXISTS (SELECT 1 FROM agendamentos WHERE barbeiro_id = $1 AND data_hora = $2 AND status = 'agendado')", barbeiroID, dataHora).Scan(&existe)
	return existe, err
}
func (r *AgendamentoRepository) Cancelar(ctx context.Context, id int) (bool, error) {
	tag, err := r.db.Exec(ctx, "UPDATE agendamentos SET status = 'cancelado' WHERE id = $1", id)
	if err != nil {
		return false, err
	}
	return tag.RowsAffected() > 0, nil
}
