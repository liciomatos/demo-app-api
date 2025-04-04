package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Message string `json:"message"`
	Version string `json:"version"`
}

func main() {
	// Configuração da porta da API
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Handler para a rota principal
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := Response{
			Message: "Bem-vindo à Demo API",
			Version: "1.0.0",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// Handler para verificação de saúde
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Iniciar o servidor
	fmt.Printf("Servidor rodando na porta %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
