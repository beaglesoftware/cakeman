Write-Host "Welcome! This script will install Cakeman."
Write-Host "Let's download the latest version of Cakeman"
Invoke-WebRequest $myDownloadUrl -OutFile "$env:USERPROFILE/Downloads/cakeman-temp/Cakeman.zip"
Write-Host "Extracting Cakeman"
Expand-Archive -Path "$env:USERPROFILE/Downloads/cakeman-temp/Cakeman.zip" -DestinationPath "$env:USERPROFILE/Downloads/cakeman-temp/cakeman"
Write-Host "Copying Cakeman to the installation directory"
Copy-Item -Path "$env:USERPROFILE/Downloads/cakeman-temp/cakeman" -Destination "$env:USERPROFILE/.cakeman/bin/" -Recurse

Write-Host "Cleaning up temporary files"
Remove-Item -Path "$env:USERPROFILE/Downloads/cakeman-temp" -Recurse
Write-Host "Adding Cakeman to PATH"
$currentpath = [System.Environment]::GetEnvironmentVariable("PATH", "User")
[System.Environment]::SetEnvironmentVariable("PATH", "$currentpath;$env:USERPROFILE/.cakeman/bin", "System")
Write-Host "Installation complete!"
