# SpotifyPlus
Spotify with a better interface and a better algorithm

To create a README that explains how to build the code using the build script located in the `/extension` folder of your repository, follow this template. This guide assumes the users have basic knowledge of using terminal or command prompt, have cloned the repository, and have basic setup requirements met, such as having Git and possibly Node.js (depending on your project's requirements).

---

# Building SpotifyPlus Extension

Welcome to the SpotifyPlus extension! This guide will walk you through the steps to build and deploy the SpotifyPlus extension from the cloned repository to your local Spotify application.

## Prerequisites

Before you begin, ensure you have the following prerequisites installed and set up:

- **Git**: To clone the repository.
- **Spicetify**: Ensure Spicetify is installed and configured on your system. Spicetify is a tool to customize Spotify.
- **Windows Subsystem for Linux (WSL)** (for Windows users): Ensure WSL is installed and set up if you're on Windows, you cannot use powershell to run this script.

## Clone the Repository

If you haven't already cloned the repository, run the following command in your terminal or command prompt:

```bash
git clone [https://github.com/LeadFreeCandy/SpotifyPlus]
```

## Navigate to the Extension Directory

Change your current directory to the `/extension` folder within the cloned repository:

```bash
cd SpotifyPlus/extension
```

## Running the Build Script

To execute the build script, use the following command:

```bash
./build.sh
```

### For macOS Users

The script will automatically detect your operating system. If you're on macOS, it will copy the necessary files to `~/.config/spicetify/CustomApps/SpotifyPlus` and apply the Spicetify changes.

### For Windows Users with WSL

If you're using Windows Subsystem for Linux (WSL), the script will place the necessary files in your Windows `%appdata%\spicetify\CustomApps\SpotifyPlus` directory and then apply the Spicetify changes.

## Troubleshooting

If you encounter any errors during the build process:

- Ensure you have all the prerequisites installed.
- Verify that Spicetify is correctly set up and that you have permissions to modify its directories.
- Check the script's error messages for hints on what went wrong and adjust your environment accordingly.
