# Etapa 1: build
FROM golang:1.25.1-alpine AS builder
WORKDIR /app

# Cache de módulos y build para acelerar reconstrucciones
COPY go.mod go.sum ./
COPY foundation/go.mod foundation/go.sum ./foundation/
RUN go mod download

RUN pwd

COPY . .

# Compilar el paquete de cmd/api y forzar salida del binario
# Ajusta GOARCH si necesitas otra arch.
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main cmd/api/main.go

FROM alpine:3.20

COPY --from=builder /app/main .

EXPOSE 9000
CMD ["./main"]


