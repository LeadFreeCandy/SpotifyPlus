copy_and_apply() {
  mkdir -p "$1" 2>/dev/null
  if ! cp manifest.json index.js "$1"; then
    echo "Error copying files. Ensure manifest.json and index.js exist in the current directory."
    exit 1
  fi

  if ! spicetify config custom_apps SpotifyPlus; then
    echo "Failed to configure spicetify for SpotifyPlus. Ensure spicetify is correctly installed."
    exit 1
  fi

  if ! spicetify apply; then
    echo "Failed to apply spicetify changes. Check your spicetify installation and configuration."
    exit 1
  fi
}

# Detect OS and set target path
if [ "$(uname)" = "Darwin" ]; then
  # macOS
  TARGET_DIR="$HOME/.config/spicetify/CustomApps/SpotifyPlus"
  copy_and_apply "$TARGET_DIR"
elif [ "$(expr substr $(uname -s) 1 5)" = "Linux" ]; then
  # Assuming WSL is being used on Windows
  if [ -n "$WSL_DISTRO_NAME" ]; then
    # Convert Windows %appdata% path to WSL path
    TARGET_DIR="$(wslpath "$(wslvar APPDATA)")/spicetify/CustomApps/SpotifyPlus"
    copy_and_apply "$TARGET_DIR"
  else
    echo "This script is intended to be run on macOS or within WSL for Windows."
    exit 1
  fi
else
  echo "Unsupported operating system. This script is intended for macOS or WSL on Windows."
  exit 1
fi
