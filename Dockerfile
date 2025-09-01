# Utiliser une image Ubuntu de base
FROM ubuntu:22.04

# Installer curl et bash
RUN apt-get update && \
    apt-get install -y curl bash ca-certificates && \
    apt-get clean

# Créer le dossier de travail
WORKDIR /app

# Copier le fichier .env
COPY .env /app/.env

# Installer DoduAPI
RUN curl -s https://get.dofusdu.de/doduapi | sh

# Exposer le port Railway
EXPOSE 3000

# Lancer DoduAPI
CMD ["doduapi", "--headless", "--port", "3000"]
