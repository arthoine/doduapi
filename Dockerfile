# Base
FROM debian:bullseye-slim

# Mettre à jour et installer les dépendances nécessaires
RUN apt-get update && \
    apt-get install -y curl unzip ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Définir le répertoire de travail
WORKDIR /app

# Copier uniquement ce qui est nécessaire pour installer les dépendances (optionnel)
# COPY composer.json composer.lock ./

# Copier le fichier .env si tu en as besoin localement (Railway injectera ses variables)
# COPY .env /app/.env

# Installer DoduAPI directement dans /usr/local/bin
RUN curl -sSL https://github.com/DoduAPI/releases/latest/download/doduapi-linux -o /usr/local/bin/doduapi && \
    chmod +x /usr/local/bin/doduapi

# Ajouter le dossier doduapi au PATH si nécessaire (optionnel)
ENV PATH="/usr/local/bin:${PATH}"

# Copier tout le projet
COPY . /app

# Installer les dépendances de ton projet si c'est un projet PHP/Laravel
# RUN composer install --no-dev --optimize-autoloader

# Exposer le port que Railway utilisera
EXPOSE 8080

# Commande de démarrage (adapter selon ton projet, ex: Laravel Sail, PHP, Node...)
CMD ["php", "artisan", "serve", "--host=0.0.0.0", "--port=8080"]
