# Use uma imagem base com suporte para Go
FROM golang:1.20

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie os arquivos go.mod e go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Baixe as dependências do Go
RUN go mod download

# Copie todo o código fonte para o diretório de trabalho
COPY . .

# Compile a aplicação
RUN go build -o main .

# Execute a aplicação quando o contêiner for iniciado
CMD ["./main"]
