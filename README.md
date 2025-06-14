![Cakeman Poster](assets/Poster.png)
# CForge
A no-nonsense C/C++ package manager

> [!WARNING]
> CForge is not ready for production use yet. Please wait until final release

## Installation
### Windows
Open PowerShell **as an administrator**

Run this comnand:
```powershell
irm "https://github.com/beaglesoftware/cakeman/blob/main/tools/install.ps1?raw=true" | iex
```

If you're on Windows 11 24H2, you're lucky. Open a **non-elevated** PowerShell window and run this command:
```powershell
irm "https://github.com/beaglesoftware/cakeman/blob/main/tools/install.ps1?raw=true" | sudo iex
```
It will ask for permission, click on Yes

### Mac or GNU/Linux
Open your terminal (Example: Mac's default terminal, iTerm2, GNOME Terminal, Konsole, Terminator, Warp or anything)

Run this command:
```shell
sh -c "$(curl -fsSL https://github.com/beaglesoftware/cakeman/blob/main/tools/install.sh?raw=true)"
```

## Build from source
Install [Go](https://go.dev/dl/)

### Mac/Linux
Run these commands:
```shell
./build.sh
```

The path of binary is `dist/{YOUR OS}/{YOUR ARCHITECTURE}/cforge`

### Windows
Run these commands:
```powershell
$ARCH = $env:PROCESSOR_ARCHITECTURE
go build main.go -o "dist/windows/$ARCH/cforge.exe "
./dist/windows/$ARCH/cforge.exe 
```

## Features

**Note:** "No" doesn't mean that it won't be available in the future. It may be available in the future

| Feature               | Supported OS/OSes | Available | It is working? |
|-----------------------|-------------------|-----------|----------------|
| Supports dependencies | Cross-platform     | ✅ Yes     | ✅ Yes          |
| Have a build system   | N/A               | ❌ No      | N/A            |
