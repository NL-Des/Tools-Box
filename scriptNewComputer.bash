#!/bin/bash

echo "Lancement du script d'installation des systèmes pour un nouvel ordinateur."

# Mise à jour du système.
echo "Mise à jour générale avant installation."
sudo apt update && sudo apt upgrade -y

# Installation de programmes courants.
echo "Installation outils de bases."
sudo apt install -y git curl wget vim htop

# Installation navigateur Brave.
echo "Installtion de Brave."
sudo apt install -y curl
sudo curl -fsSLo /usr/share/keyrings/brave-browser-archive-keyring.gpg https://brave-browser-apt-release.s3.brave.com/brave-browser-archive-keyring.gpg
echo "deb [signed-by=/usr/share/keyrings/brave-browser-archive-keyring.gpg] https://brave-browser-apt-release.s3.brave.com/ stable main" | sudo tee /etc/apt/sources.list.d/brave-browser-release.list
sudo apt update
sudo apt install -y brave-browser

# Installation navigateur FireFox.
echo "Installation de FireFox."
sudo apt install -y firefox

# Installation navigateur Chrome.
echo "Installation de Chrome."
wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
sudo apt install -y ./google-chrome-stable_current_amd64.deb
rm google-chrome-stable_current_amd64.deb

# Installation Libre Office.
echo "Installation de Libre Office."
sudo apt install -y libreoffice

# Installation Gimp.
echo "Installation de Gimp."
sudo apt install -y gimp

# Installation de Flameshot.
echo "Installation de Flameshot."
sudo apt install -y flameshot

# Installation de 7zip.
echo "Installation de 7zip."
sudo apt install -y p7zip-full unrar

# Installation de VLC.
echo "Installation de VLC."
sudo apt install -y vlc

# Installation de Codecs Multimédias.
echo "Installation de Codecs Multimédias."
echo ttf-mscorefonts-installer msttcorefonts/accepted-mscorefonts-eula select true | sudo debconf-set-selections
sudo apt install -y ubuntu-restricted-extras

# Installation de Discord.
wget -O discord.deb "https://discord.com/api/download?platform=linux&format=deb"
sudo apt install -y ./discord.deb
rm discord.deb

# Installation Pilotes NVIDIA.
echo "Installation des Pilotes de NVIDIA."
if lspci | grep -i nvidia > /dev/null; then
    echo "Carte NVIDIA détectée, installation des pilotes..."
    sudo ubuntu-drivers autoinstall
else
    echo "Aucune carte NVIDIA détectée, pilotes non installés..."
fi

# Installation de VSCode.
echo "Installation de VSCode."
sudo apt install -y wget gpg
wget -qO- https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor > packages.microsoft.gpg
sudo install -o root -g root -m 644 packages.microsoft.gpg /usr/share/keyrings/
sudo sh -c 'echo "deb [arch=amd64 signed-by=/usr/share/keyrings/packages.microsoft.gpg] https://packages.microsoft.com/repos/code stable main" > /etc/apt/sources.list.d/vscode.list'
sudo apt update
sudo apt install -y code
rm packages.microsoft.gpg

# Initialisation des informations GitHub.
echo "Initialisation de vos coordonnées GitHub."
read -p "Entrez votre nom complet : " name
read -p "Entrez votre adresse mail GitHub : " email

git config --global user.name "$name"
git config --global user.email "$email"

echo "Configuration GitHub terminée !"

# Installation de NodeJS.
echo "Installation de Node JS."
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh | bash

    # Chargement du npm
    export NVM_DIR="$HOME/.nvm"
    [ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"

    # Installation depuis les données du npm
    nvm install --lts

# Installation de Python.
echo "Installation de Python."
sudo apt install -y python3-pip

# Installation de Docker et Docker Compose.
echo "Installation de Docker et Docker Compose"
sudo apt install -y ca-certificates curl gnupg lsb-release
sudo mkdir -p /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt install -y docker-ce docker-ce-cli containerd.io