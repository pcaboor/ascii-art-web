# Utiliser une image de base officielle de Go
FROM golang:1.22-alpine

# Définir le répertoire de travail dans le conteneur
WORKDIR /app

# Copier les fichiers go.mod et go.sum, et installer les dépendances
COPY go.mod ./
RUN go mod download

# Copier tous les fichiers du projet dans le conteneur
COPY . .

# Construire l'application
RUN go build -o main .

# Exposer le port sur lequel l'application écoute
EXPOSE 8080

# Ajouter des labels pour fournir des métadonnées à l'image
LABEL maintainer="Pierre Caboor pierre.caboor59@gmail.com"
LABEL description="Image Docker in Go 🐳"
LABEL version="1.0"
LABEL com.example.department="IT"

# Commande pour démarrer l'application
CMD ["./main"]


