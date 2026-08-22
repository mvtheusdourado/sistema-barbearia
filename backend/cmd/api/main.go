package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"github.com/mvtheusdourado/sistema-barbearia/internal/handlers"
	"github.com/mvtheusdourado/sistema-barbearia/internal/repository"
	"github.com/mvtheusdourado/sistema-barbearia/internal/service"
)

func main() {
	godotenv.Load()

	databaseURL := os.Getenv("DATABASE_URL")

	pool, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		log.Fatal("Não foi possível conectar-se ao banco: ", err)
	}
	defer pool.Close()

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal("O banco não respondeu ao ping: ", err)
	}

	log.Println("✅ Conectado ao banco de dados com sucesso!")

	clienteRepository := repository.NewClienteRepository(pool)
	clienteService := service.NewClienteService(clienteRepository)
	clienteHandler := handlers.NewClienteHandler(clienteService)

	agendamentoRepository := repository.NewAgendamentoRepository(pool)
	agendamentoService := service.NewAgendamentoService(agendamentoRepository)
	agendamentoHandler := handlers.NewAgendamentoHandler(agendamentoService)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá, Barbearia! O servidor está funcionando.")
	})

	mux.HandleFunc("GET /clientes", clienteHandler.ListarClientes)

	mux.HandleFunc("POST /agendamentos", agendamentoHandler.CriarAgendamento)

	mux.HandleFunc("PATCH /agendamentos/{id}/cancelar", agendamentoHandler.CancelarAgendamento)

	mux.HandleFunc("GET /barbeiros/{id}/horarios", agendamentoHandler.ListarHorariosDisponiveis)

	mux.HandleFunc("GET /agendamentos", agendamentoHandler.ListarAgendamentos)

	porta := os.Getenv("PORT")
	if porta == "" {
		porta = "8080"
	}

	log.Println("Servidor ligado na porta " + porta)

	http.ListenAndServe(":"+porta, comLog(comCORS(mux)))
}
func comCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
func comLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		inicio := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s — levou %v", r.Method, r.URL.Path, time.Since(inicio))
	})
}
