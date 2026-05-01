# Build stage
FROM golang:1.26.1-alpine AS builder

WORKDIR /app

# Descargar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar la aplicación principal
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api

# Run stage
FROM alpine:latest

WORKDIR /app

# Certificados SSL y timezones
RUN apk --no-cache add ca-certificates tzdata

# Copiar el binario desde la fase de construcción
COPY --from=builder /app/api .

# Exponer el puerto donde corre la API
EXPOSE 8090

# Comando para ejecutar la aplicación
CMD ["./api"]
