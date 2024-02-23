@echo off
SetLocal EnableDelayedExpansion

:: Function equivalent in batch for copy_and_apply
:copy_and_apply
mkdir "%~1" 2>nul
copy manifest.json "%~1" >nul
if errorlevel 1 (
  echo Error copying files. Ensure manifest.json and index.js exist in the current directory.
  exit /b 1
)
copy index.js "%~1" >nul
if errorlevel 1 (
  echo Error copying files. Ensure manifest.json and index.js exist in the current directory.
  exit /b 1
)

:: Attempt to configure spicetify for SpotifyPlus
spicetify config custom_apps SpotifyPlus
if errorlevel 1 (
  echo Failed to configure spicetify for SpotifyPlus. Ensure spicetify is correctly installed.
  exit /b 1
)

:: Apply spicetify changes
spicetify apply
if errorlevel 1 (
  echo Failed to apply spicetify changes. Check your spicetify installation and configuration.
  exit /b 1
)

goto :eof

:: Detect OS and set target path
for /f "tokens=2 delims==" %%i in ('wmic os get caption /value') do set OS=%%i

if "%OS%" == "Microsoft Windows 10 Pro" (
  :: Assuming this is for demonstration; adjust as necessary for your Windows version
  :: Convert %appdata% to a direct path usage in PowerShell/Command Prompt
  set TARGET_DIR=%appdata%\spicetify\CustomApps\SpotifyPlus
  call :copy_and_apply "%TARGET_DIR%"
) else (
  echo Unsupported operating system. This script is intended for Windows with PowerShell.
  exit /b 1
)
