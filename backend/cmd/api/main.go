package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá, Barbearia! O servidor está funcionando.")
	})

	fmt.Println("Servidor ligado! Acesse: http://localhost:8080/health")

	http.ListenAndServe(":8080", mux)
}
