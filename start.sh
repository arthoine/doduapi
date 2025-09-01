#!/bin/bash
set -e

# Lancer les migrations
doduapi migrate up

# Lancer doduapi en mode headless
exec doduapi --headless
