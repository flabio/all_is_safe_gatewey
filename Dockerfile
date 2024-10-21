# Construir la aplicación
FROM golang:1.19 AS build

WORKDIR /gatewey

# Copiar los archivos de Go y descargar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del código fuente
COPY . .

# Compilar la aplicación
RUN go build -o main .

# Imagen final para ejecutar la aplicación
FROM golang:1.19

WORKDIR /gatewey

# Copiar el binario construido en la etapa anterior
COPY --from=build /gatewey/main .

# Establecer la variable de entorno si es necesario
ENV PORT=8080

# Exponer el puerto utilizado por la aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]
