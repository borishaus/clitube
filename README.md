# cliTube

A simple CLI tool to play YouTube videos (audio or video) using memorable aliases.

## Features

- Stream audio-only from YouTube videos by default
- Optional video streaming with `-v` flag
- Save YouTube URLs with custom aliases for quick access
- Persistent storage of your favorite videos
- Simple and intuitive command-line interface

## Prerequisites

Before using cliTube, you need to install **mpv** - a media player with built-in YouTube streaming support:

```bash
# On Ubuntu/Debian
sudo apt install mpv

# On macOS
brew install mpv

# On Arch Linux
sudo pacman -S mpv
```

## Installation

### Quick Install (Recommended)

```bash
# Clone or navigate to the repository
cd cliTube

# Run the installation script
./install.sh
```

This will:
- Build the application
- Install the binary to `/usr/local/bin/`
- Install the man page to `/usr/local/share/man/man1/`

### Manual Installation

```bash
# Build the application
go build -o clitube

# Install binary
sudo mv clitube /usr/local/bin/

# Install man page
sudo mkdir -p /usr/local/share/man/man1
sudo cp clitube.1 /usr/local/share/man/man1/
sudo mandb
```

### Development Mode

```bash
# Build and run locally without installing
go build -o clitube
./clitube help
```

## Usage

### Add a video alias
```bash
clitube add lofigirl "https://www.youtube.com/watch?v=jfKfPfyJRdk"
clitube add lockpickinglawyer "https://www.youtube.com/watch?v=TZXcXW8fNYI"
```

### Play audio (default)
```bash
clitube lofigirl
```

### Play with video
```bash
clitube -v lockpickinglawyer
```

### List all saved aliases
```bash
clitube list
```

### Remove an alias
```bash
clitube remove lofigirl
# or
clitube rm lofigirl
```

### Get help
```bash
clitube help

# View full manual page (if installed)
man clitube
```

## Configuration

Video mappings are stored in `~/.config/clitube/videos.json`

Playback history (last 3 played) is stored in `~/.config/clitube/history.json`

### First Run

When you run cliTube for the first time, you'll see helpful tips and examples to get started.

### Recently Played

cliTube automatically tracks your last 3 played videos. View them with:

```bash
clitube list
```

Or run `clitube` without arguments to see your recent history in the help output.

## Examples

```bash
# Add some videos
clitube add lofi "https://www.youtube.com/watch?v=jfKfPfyJRdk"
clitube add jazz "https://www.youtube.com/watch?v=Dx5qFachd3A"

# Play audio in the background while working
clitube lofi

# Watch a video
clitube -v jazz

# See what you have saved
clitube list
```

## How it Works

cliTube uses `mpv` to play streams directly from YouTube URLs. mpv has built-in support for YouTube streaming, so no additional tools are needed. By default, cliTube streams audio-only to save bandwidth. When you use the `-v` flag, it streams both video and audio.

The tool stores your alias-to-URL mappings in a JSON file in your config directory, making it easy to quickly access your favorite videos without remembering URLs.
