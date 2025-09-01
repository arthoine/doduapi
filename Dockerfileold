# Base Debian slim
FROM debian:bullseye-slim

# Installer curl et unzip
RUN apt-get update && \
    apt-get install -y curl unzip && \
    rm -rf /var/lib/apt/lists/*

# Télécharger le binaire doduapi directement depuis GitHub et le rendre exécutable
RUN curl -L https://github.com/dofusdude/doduapi/releases/latest/download/doduapi-linux-amd64 -o /usr/local/bin/doduapi && \
    chmod +x /usr/local/bin/doduapi

# Créer dossier de travail
WORKDIR /app

# Copier le script de démarrage
COPY start.sh /app/start.sh
RUN chmod +x /app/start.sh

# Exposer le port de l'API
EXPOSE 3000

# Lancer le script de démarrage
CMD ["/app/start.sh"]
