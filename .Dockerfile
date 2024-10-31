# Usa uma imagem base do Go
FROM golang:1.19

# Define o diretório de trabalho
WORKDIR /app

# Copia todos os arquivos para o container
COPY . .

# Instala as dependências e compila o binário
RUN go build -o main .

# Expõe a porta da aplicação
EXPOSE 8080

# Comando para iniciar a aplicação
CMD ["/app/main"]
