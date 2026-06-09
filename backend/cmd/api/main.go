package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"github.com/mvtheusdourado/sistema-barbearia/internal/handlers"
	"github.com/mvtheusdourado/sistema-barbearia/internal/repository"
	"github.com/mvtheusdourado/sistema-barbearia/internal/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env: ", err)
	}

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

	fmt.Println("Conectado ao banco de dados com sucesso!")

	clienteRepository := repository.NewClienteRepository(pool)
	clienteService := service.NewClienteService(clienteRepository)
	clienteHandler := handlers.NewClienteHandler(clienteService)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá, Barbearia! O servidor está funcionando.")
	})

	mux.HandleFunc("GET /clientes", clienteHandler.ListarClientes)

	fmt.Println("Servidor ligado! Acesse: http://localhost:8080/health")

	http.ListenAndServe(":8080", mux)
}
