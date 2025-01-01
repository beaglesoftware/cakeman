echo "\033[0;34m\033[0;1mINFO>\033[0;1m Welcome! This script will install Cakeman\033[0;0m"
echo "\033[0;34m\033[0;1mINFO>\033[0;1m Getting current version of Cakeman...\033[0;0m"
echo '\033[0;34m\033[0;1m===>\033[0;1m export version "$(curl -fsSL 'https://github.com/beaglesoftware/cakeman/blob/main/VERSION?raw=true')"\033[0;0m'
export version "$(curl -fsSL 'https://github.com/beaglesoftware/cakeman/blob/main/VERSION?raw=true')"
echo "\033[0;34m===>\033[0;1m mkdir ~/.cakeman/\033[0;0m"
mkdir ~/.cakeman/
echo "\033[0;34m===>\033[0;1m mkdir ~/.cakeman/temp/\033[0;0m"
mkdir ~/.cakeman/temp/
echo "\033[0;34m===>\033[0;1m mkdir ~/.cakeman/bin/\033[0;0m"
mkdir ~/.cakeman/bin/
echo "\033[0;34mINFO>\033[0;1m Downloading Cakeman...\033[0;0m"
echo "\033[0;34m===>\033[0;1m curl -fsSL https://github.com/beaglesoftware/cakeman/releases -o ~/.cakeman/temp/cakeman.zip\033[0;0m"
curl -fsSL https://github.com/beaglesoftware/cakeman/releases -o ~/.cakeman/temp/cakeman.zip
echo "\033[0;34m===>\033[0;1m unzip ~/.cakeman/temp/cakeman.zip -d ~/.cakeman/temp/cakeman\033[0;0m"
unzip ~/.cakeman/temp/cakeman.zip -d ~/.cakeman/temp/cakeman
echo "\033[0;34m===>\033[0;1m Installing Cakeman...\033[0;0m"
echo '\033[0;34m===>\033[0;1m mv "~/.cakeman/temp/cakeman/*" ~/.cakeman/bin/\033[0;0m'
mv "~/.cakeman/temp/cakeman/*" ~/.cakeman/bin/
echo "\033[0;34mINFO>\033[0;1m Let's clean the room for Cakeman"
echo "\033[0;34m===>\033[0;1m rm -rf ~/.cakeman/temp/"
rm -rf ~/.cakeman/temp/

echo "\033[0;32mCakeman has been installed!\033[0;0m"
echo ""
echo "Now add the following line to your .bashrc or .zshrc:"
echo "    export PATH=\$PATH:~/.cakeman/bin/"
