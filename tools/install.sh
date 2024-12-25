echo "Welcome! This script will install Cakeman"
export version "$(curl -fsSL 'https://github.com/beaglesoftware/cakeman/blob/main/VERSION?raw=true')"
curl -fsSL https://github.com/beaglesoftware/cakeman/releases -o ~/.cakeman/temp/cakeman.zip
mkdir ~/.cakeman/
mkdir ~/.cakeman/temp/
mkdir ~/.cakeman/bin/
unzip ~/.cakeman/temp/cakeman.zip -d ~/.cakeman/temp/cakeman
echo "Installing Cakeman..."
mv "~/.cakeman/temp/cakeman/*" ~/.cakeman/bin/

echo "Cakeman has been installed!"
echo ""
echo "Now add the following line to your .bashrc or .zshrc:"
echo "    export PATH=\$PATH:~/.cakeman/bin/"
