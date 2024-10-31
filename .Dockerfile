# Usando uma imagem base do Go
FROM golang:1.20

# Habilitar CGO (opcional, caso necessário)
ENV CGO_ENABLED=1

# Definindo o diretório de trabalho
WORKDIR /app

# Instalando GCC
RUN apt-get update && apt-get install -y gcc

# Copiando o módulo e o arquivo de dependências para o container
COPY go.mod ./
COPY go.sum ./

# Baixando as dependências
RUN go mod download

# Copiando o código do projeto para o diretório de trabalho do container
COPY . .

# Compilando o aplicativo Go
RUN go build -o main .

# Definindo a porta que a aplicação vai escutar
EXPOSE 8080

# Comando para rodar o binário
CMD ["/app/main"]
