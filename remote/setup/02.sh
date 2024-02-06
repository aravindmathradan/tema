#!/bin/bash
set -eu

# ==================================================================================== #
# SCRIPT LOGIC
# ==================================================================================== #

# Install nvm and install node lts version
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash
source ~/.nvm/nvm.sh
nvm install --lts

# Install pm2 globally
npm install pm2@latest -g

echo "Script complete! Rebooting..."
sudo reboot