# Etapa 1: Construção da aplicação (Builder)
FROM golang:1.20 AS builder

# Habilitar CGO
ENV CGO_ENABLED=1

# Definindo o diretório de trabalho dentro do container
WORKDIR /app

# Copiando o módulo e o arquivo de dependências para o container
COPY go.mod ./
COPY go.sum ./

# Baixando as dependências
RUN go mod download

# Instalando compiladores que são necessários para CGO
RUN apt-get update && apt-get install -y gcc

# Copiando o código-fonte do projeto para o diretório de trabalho do container
COPY . .

# Compilando o aplicativo Go
RUN go build -o main .

# Etapa 2: Criação da imagem final (Mais leve)
FROM debian:bookworm-slim

# Definindo o diretório de trabalho dentro do container final
WORKDIR /app

# Copiando o binário da etapa de construção para a imagem final
COPY --from=builder /app/main .

# Definindo a porta que a aplicação vai escutar
EXPOSE 8080

# Comando para rodar o binário
CMD ["/app/main"]
