#!/bin/bash
set -e

echo "🚀 Iniciando deploy local para Minikube"

# Verifica se o Minikube está rodando
if ! minikube status | grep -q "Running"; then
  echo "⚙️ Iniciando Minikube..."
  minikube start
fi

# Configura o Docker para usar o daemon do Minikube
echo "⚙️ Configurando Docker para usar o daemon do Minikube..."
eval $(minikube docker-env)

# Constrói a imagem Docker
echo "🔨 Construindo imagem Docker..."
docker build -t demo-app-api:latest .

# Aplica os manifestos Kubernetes
echo "📦 Aplicando manifestos Kubernetes..."
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml

# Aguarda o deployment completar
echo "⏳ Aguardando o deployment completar..."
kubectl rollout status deployment/demo-app-api

# Obtém a URL do serviço
echo "🌐 Obtendo URL do serviço..."
minikube service demo-app-api --url

echo "✅ Deploy completo!" 