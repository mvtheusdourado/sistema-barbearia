package models

import "time"

type Agendamento struct {
	ID         int       `json:"id"`
	ClienteID  int       `json:"cliente_id"`
	BarbeiroID int       `json:"barbeiro_id"`
	DataHora   time.Time `json:"data_hora"`
	Status     string    `json:"status"`
}
