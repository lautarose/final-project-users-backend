# Etapa de construcción para compilar la aplicación Go
FROM golang:alpine AS builder

WORKDIR /app

# Copia los archivos de tu aplicación Go al contenedor
COPY . .

# Inicializar el módulo Go (solo por fines demostrativos, no suele ir en el Dockerfile)
RUN go mod init user-backend

# Actualizar las dependencias y los módulos (solo por fines demostrativos, no suele ir en el Dockerfile)
RUN go mod tidy

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
