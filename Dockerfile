# Usamos una imagen base de Go para la compilación
FROM golang:1.20-alpine AS builder

# Establecemos el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiamos los archivos go.mod y go.sum (esto permite cachear dependencias)
COPY go.mod go.sum ./

# Instalamos las dependencias necesarias
RUN go mod download

# Copiamos todo el código fuente al contenedor
COPY . .

# Compilamos la aplicación Go, generando el binario "main"
RUN go build -o main .

# Usamos una imagen más ligera para ejecutar la aplicación
FROM alpine:latest

# Establecemos el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiamos el binario "main" desde la etapa de construcción
COPY --from=builder /app/main .

# Establecemos permisos de ejecución en el binario "main"
RUN chmod +x main

# Ejecutamos el binario "main" cuando el contenedor se inicie
CMD ["./main"]