# Demo App API

Uma API simples em Go para demonstração de deploy em Minikube utilizando CI/CD com GitHub Actions.

## Requisitos

- Go 1.20+
- Docker
- Minikube
- kubectl

## Estrutura do Projeto

```
├── cmd/
│   └── api/        # Código principal da API
├── docs/           # Documentação Swagger
├── k8s/            # Manifestos Kubernetes
├── pkg/            # Pacotes reutilizáveis
├── scripts/        # Scripts auxiliares
└── Dockerfile      # Arquivo para build da imagem Docker
```

## Como Executar Localmente

1. Clone o repositório:
   ```
   git clone https://github.com/liciomatos/demo-app-api.git
   cd demo-app-api
   ```

2. Teste a aplicação localmente:
   ```
   go run cmd/api/main.go
   ```

3. Use o script de deploy para Minikube:
   ```
   ./scripts/deploy-local.sh
   ```

## Documentação da API (Swagger)

A API inclui documentação interativa usando Swagger UI:

1. Quando a API estiver em execução, acesse:
   ```
   http://localhost:8080/swagger/index.html
   ```

2. A documentação Swagger permite testar os endpoints diretamente na interface.

3. Para atualizar a documentação após modificações na API:
   ```
   swag init -g cmd/api/main.go
   ```

## CI/CD

O projeto utiliza GitHub Actions para CI/CD:

- **CI:** Executa testes e build automaticamente em cada push/PR para a branch main.
- **Local Deployment:** Configurado para executar em um ambiente Minikube local.

## Como Testar a API

1. Após o deploy, obtenha a URL do serviço:
   ```
   minikube service demo-app-api --url
   ```

2. Teste os endpoints:
   ```
   # Endpoint principal
   curl http://<service-url>/
   
   # Verificação de saúde
   curl http://<service-url>/health
   
   # Documentação Swagger
   http://<service-url>/swagger/index.html
   ```

## Customização

Você pode personalizar os parâmetros da aplicação através de variáveis de ambiente:

- `PORT`: Porta em que a API vai escutar (padrão: 8080) 