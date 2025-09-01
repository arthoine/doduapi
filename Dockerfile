# Base Debian
FROM debian:bullseye-slim

# Installer dépendances
RUN apt-get update && \
    apt-get install -y curl unzip ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Créer le dossier de travail
WORKDIR /app

# Installer DoduAPI
RUN curl -s https://get.dofusdu.de/doduapi | sh

# Ajouter le dossier doduapi à PATH
ENV PATH="/root/.doduapi:${PATH}"

# Exposer le port (Railway remplace $PORT)
EXPOSE 3000

# Start DoduAPI
CMD ["doduapi", "--headless", "--port", "3000"]
