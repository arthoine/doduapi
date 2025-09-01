RUN curl -s https://get.dofusdu.de/doduapi | sh
# Vérifier où le binaire est installé
RUN ls -la /root/.local/bin
# Ajouter au PATH si nécessaire
ENV PATH="/root/.local/bin:${PATH}"
