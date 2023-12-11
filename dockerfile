# Etapa de construcción para compilar la aplicación Go
FROM golang:alpine AS builder

WORKDIR /app

# Copia los archivos de tu aplicación Go al contenedor
COPY . .

# Compila tu aplicación Go
RUN go build -o user-backend

# Etapa de producción con Alpine
FROM alpine:latest

WORKDIR /app

# Copia el ejecutable compilado desde la etapa de construcción
COPY --from=builder /app/user-backend .

# Puerto en el que escucha tu aplicación Go
EXPOSE 8080

# Comando para ejecutar tu aplicación Go
CMD ["./user-backend"]