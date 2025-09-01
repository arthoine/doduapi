#!/bin/bash
set -e

# Migration de la base
doduapi migrate up

# Lancer doduapi en mode headless pour éviter les erreurs TTY
exec doduapi --headless
