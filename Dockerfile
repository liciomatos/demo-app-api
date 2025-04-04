FROM golang:1.20-alpine AS builder

WORKDIR /app

# Copiar arquivos de dependências
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copiar o código da aplicação
COPY . .

# Compilar o aplicativo
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api

# Imagem de produção
FROM alpine:3.18

WORKDIR /app

# Copiar o binário compilado
COPY --from=builder /app/api .

# Porta que a aplicação vai escutar
EXPOSE 8080

# Executar a aplicação
CMD ["./api"] 