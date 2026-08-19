package service

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/mvtheusdourado/sistema-barbearia/internal/models"
)

func TestEstaNoPassado(t *testing.T) {
	// Caso 1: uma data no passado deve retornar true
	passado := time.Date(2020, 1, 1, 10, 0, 0, 0, time.UTC)
	if !estaNoPassado(passado) {
		t.Errorf("esperava TRUE para uma data no passado, mas veio false")
	}

	// Caso 2: uma data no futuro deve retornar false
	futuro := time.Date(2030, 1, 1, 10, 0, 0, 0, time.UTC)
	if estaNoPassado(futuro) {
		t.Errorf("esperava FALSE para uma data no futuro, mas veio true")
	}
}
func TestCalcularHorariosDisponiveis(t *testing.T) {
	casos := []struct {
		nome     string
		ocupados []time.Time
		esperado []string
	}{
		{
			nome:     "dia vazio retorna todos os 9 slots",
			ocupados: []time.Time{},
			esperado: []string{"09:00", "10:00", "11:00", "12:00", "13:00", "14:00", "15:00", "16:00", "17:00"},
		},
		{
			nome:     "com as 10h ocupadas, o 10:00 some",
			ocupados: []time.Time{time.Date(2030, 1, 1, 10, 0, 0, 0, time.UTC)},
			esperado: []string{"09:00", "11:00", "12:00", "13:00", "14:00", "15:00", "16:00", "17:00"},
		},
	}

	for _, caso := range casos {
		resultado := calcularHorariosDisponiveis(caso.ocupados)
		if !reflect.DeepEqual(resultado, caso.esperado) {
			t.Errorf("%s: esperava %v, mas veio %v", caso.nome, caso.esperado, resultado)
		}
	}
}

type fakeRepo struct {
	ocupado      bool // controla o que o ExisteHorarioOcupado devolve
	criarChamado bool // registra se o Criar foi chamado
	cancelou     bool
}

func (f *fakeRepo) Criar(ctx context.Context, ag models.Agendamento) error {
	f.criarChamado = true
	return nil
}
func (f *fakeRepo) ExisteHorarioOcupado(ctx context.Context, barbeiroID int, dataHora time.Time) (bool, error) {
	return f.ocupado, nil
}
func (f *fakeRepo) Cancelar(ctx context.Context, id int) (bool, error) {
	return f.cancelou, nil
}
func (f *fakeRepo) HorariosOcupados(ctx context.Context, barbeiroID int, dia time.Time) ([]time.Time, error) {
	return nil, nil
}
func (f *fakeRepo) ListarPorCliente(ctx context.Context, clienteID int) ([]models.Agendamento, error) {
	return nil, nil
}
func TestCriarAgendamento(t *testing.T) {
	futuro := time.Date(2030, 1, 1, 10, 0, 0, 0, time.UTC)
	passado := time.Date(2020, 1, 1, 10, 0, 0, 0, time.UTC)

	// Caso 1: horário livre + futuro → deve criar, sem erro
	fake := &fakeRepo{ocupado: false}
	s := NewAgendamentoService(fake)
	err := s.CriarAgendamento(context.Background(), models.Agendamento{DataHora: futuro})
	if err != nil {
		t.Errorf("esperava sucesso, veio erro: %v", err)
	}
	if !fake.criarChamado {
		t.Errorf("esperava que o Criar fosse chamado, mas não foi")
	}

	// Caso 2: horário ocupado → deve dar erro e NÃO criar
	fakeOcup := &fakeRepo{ocupado: true}
	s2 := NewAgendamentoService(fakeOcup)
	err = s2.CriarAgendamento(context.Background(), models.Agendamento{DataHora: futuro})
	if err == nil {
		t.Errorf("esperava erro de horário ocupado, veio nil")
	}
	if fakeOcup.criarChamado {
		t.Errorf("NÃO deveria chamar Criar com horário ocupado")
	}

	// Caso 3: data no passado → deve dar erro
	fake3 := &fakeRepo{ocupado: false}
	s3 := NewAgendamentoService(fake3)
	err = s3.CriarAgendamento(context.Background(), models.Agendamento{DataHora: passado})
	if err == nil {
		t.Errorf("esperava erro de data no passado, veio nil")
	}
}
func TestCancelarAgendamento(t *testing.T) {
	// Caso 1: existe → cancela sem erro
	fake := &fakeRepo{cancelou: true}
	s := NewAgendamentoService(fake)
	if err := s.CancelarAgendamento(context.Background(), 1); err != nil {
		t.Errorf("esperava sucesso, veio erro: %v", err)
	}

	// Caso 2: não existe → erro "não encontrado"
	fake2 := &fakeRepo{cancelou: false}
	s2 := NewAgendamentoService(fake2)
	if err := s2.CancelarAgendamento(context.Background(), 999); err == nil {
		t.Errorf("esperava erro de não encontrado, veio nil")
	}
}

func TestListarHorariosDisponiveis(t *testing.T) {
	fake := &fakeRepo{} // HorariosOcupados devolve nil → nada ocupado
	s := NewAgendamentoService(fake)
	horarios, err := s.ListarHorariosDisponiveis(context.Background(), 1, time.Now())
	if err != nil {
		t.Errorf("veio erro: %v", err)
	}
	if len(horarios) != 9 {
		t.Errorf("esperava 9 horários livres, veio %d", len(horarios))
	}
}
