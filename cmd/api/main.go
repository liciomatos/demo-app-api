package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/liciomatos/demo-app-api/docs" // Importa os arquivos gerados pelo swag
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Demo App API
// @version         1.0
// @description     Uma API simples para demonstração de deploy em Minikube.
// @host            localhost:8080
// @BasePath        /
// @schemes         http https
type Response struct {
	Message string `json:"message"`
	Version string `json:"version"`
}

// @Summary      Obter informação principal
// @Description  Retorna uma mensagem de boas-vindas e a versão da API
// @Tags         info
// @Produce      json
// @Success      200  {object}  Response
// @Router       / [get]
func indexHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Bem-vindo à Demo API",
		Version: "1.0.0",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// @Summary      Verificação de saúde
// @Description  Endpoint para verificar se a API está funcionando
// @Tags         health
// @Produce      plain
// @Success      200  {string}  string  "OK"
// @Router       /health [get]
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// @title           Demo App API
// @version         1.0
// @description     Uma API simples para demonstração
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /
func main() {
	// Configuração da porta da API
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Configurar rotas
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/health", healthHandler)

	// Adicionar handler para o Swagger UI
	http.HandleFunc("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // URL relativa, não depende de hostname
	))

	// Iniciar o servidor
	fmt.Printf("Servidor rodando na porta %s...\n", port)
	fmt.Printf("Documentação Swagger disponível em /swagger/index.html\n")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
