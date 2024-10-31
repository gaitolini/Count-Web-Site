# Visão Geral🚀

Este repositório contém o backend de um contador de visitas, desenvolvido em Go (Golang), hospedado utilizando o Fly.io. Ele incrementa e exibe o número de visitas únicas ao site de portfólio de Anderson Gaitolini: [https://anderson.gaitolini.com.br/](https://anderson.gaitolini.com.br/). O backend gerencia o contador através de uma API REST, que é consumida pela aplicação frontend, permitindo exibir e armazenar contagens de visitas de maneira segura e eficiente.

Repositório do backend: [Count-Web-Site](https://github.com/gaitolini/Count-Web-Site)

## Tecnologias Utilizadas🎈
- **Go (Golang)**: Utilizado para implementar o servidor backend e a API REST.
- **PostgreSQL**: Banco de dados para armazenar o contador de visitas de maneira persistente.
- **Fly.io**: Utilizado para hospedagem do backend e do banco de dados PostgreSQL.
- **Docker**: Para criar imagens containerizadas do backend.

## Configuração e Deploy no Fly.io✈️🛩️🪽

### Instruções para Deploy
1. Primeiro, precisamos criar o ambiente no Fly.io:
   ```sh
   flyctl launch
   ```
   Durante o processo, selecione o tipo de aplicativo, defina a região, escolha 256MB de RAM e configure um banco PostgreSQL.

2. Com o ambiente criado, faça o deploy do backend:
   ```sh
   flyctl deploy
   ```

3. Caso precise monitorar logs:
   ```sh
   flyctl logs -a nome_do_app
   ```

4. Para rodar o servidor localmente durante o desenvolvimento:
   ```sh
   go run main.go
   ```
   Certifique-se de configurar a string de conexão com o banco corretamente para conectar ao PostgreSQL local.

### Configuração do Dockerfile🐳
O Dockerfile utilizado é responsável por compilar o backend e preparar a imagem para o deploy:
```dockerfile
# Usando uma imagem base do Go
FROM golang:1.20

# Definindo o diretório de trabalho
WORKDIR /app

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
```

### Fly.io
O backend foi hospedado utilizando [Fly.io](https://fly.io/), configurado com PostgreSQL para armazenamento das visitas.

## Referências ao Frontend
O frontend que consome este backend está disponível no seguinte repositório: [Gaitolini-WebPage](https://github.com/gaitolini/Gaitolini-WebPage). 

## Habilidades Demonstradas
- **Go (Golang)**: Construção de APIs REST eficientes.
- **Infraestrutura em Cloud**: Configuração e deploy em Fly.io.
- **Integração de Serviços**: Uso de PostgreSQL com Docker e configuração de Cloudflare Tunnel para garantir a segurança e performance do serviço.


## Se você quiser saber mais sobre este projeto ou discutir futuras colaborações, entre em contato comigo:

- LinkedIn: [Anderson Gaitolini](https://www.linkedin.com/in/andersongaitolini/)
- WhatsApp: [Entre em contato](https://wa.me/qr/CFND4RGOJHHUN1)
<div align="center">
  <img src="https://github.com/user-attachments/assets/a019d3e6-5a04-4124-b42b-b5824d1f92d5" alt="image" width="300"/>
</div>



🚀 Feito com muito Go, Dockere Cloud por Anderson Gaitolini❤️.

## Agradecimentos 🙌
Um grande obrigado a todos que contribuíram para o desenvolvimento deste projeto e ajudaram a aprimorar a automação e o deploy contínuo! 🎉

## Fique à vontade para contribuir, testar e sugerir melhorias! 😄
