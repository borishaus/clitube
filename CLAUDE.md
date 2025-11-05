# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

cliTube is a CLI application written in Go that allows users to stream YouTube videos (audio or video) using memorable aliases. Users can assign custom names to YouTube URLs and quickly play them from the command line.

## Architecture

The codebase is organized into three main files:

- **main.go**: CLI interface and command routing. Handles argument parsing, command dispatch (add/list/remove/play), the `-v` video flag, first-run detection, and help/history display
- **storage.go**: Persistence layer using JSON storage. Manages:
  - VideoMapping: Alias-to-URL mappings in `~/.config/clitube/videos.json`
  - PlaybackHistory: Last 3 played items in `~/.config/clitube/history.json`
  - Tracks timestamps and playback mode (audio vs video)
- **player.go**: Media playback via mpv. Handles dependency checking and process execution with appropriate flags (`--no-video` for audio-only mode)

The application relies on **mpv** as the only external dependency. mpv has built-in YouTube streaming support, so no yt-dlp or similar tools are needed.

### Key Features

- **First-run experience**: Detects when running for the first time and displays helpful tips
- **History tracking**: Automatically tracks the last 3 played items with timestamps and playback mode
- **Recent history display**: Shows recent plays in the `list` command and when running without arguments

## Building and Running

```bash
# Build the application
go build -o clitube

# Run directly during development
go run main.go [commands]

# Example commands
./clitube add lofi "https://www.youtube.com/watch?v=jfKfPfyJRdk"
./clitube list                    # Shows recently played + all aliases
./clitube lofi                    # Audio-only (default)
./clitube -v lofi                 # With video
./clitube remove lofi             # Or use 'rm' alias
```

## Installation

```bash
# Quick install (installs binary + man page)
./install.sh

# Manual installation
go build -o clitube
sudo mv clitube /usr/local/bin/
sudo mkdir -p /usr/local/share/man/man1
sudo cp clitube.1 /usr/local/share/man/man1/
sudo mandb
```

## Configuration Storage

### Video Mappings
`~/.config/clitube/videos.json`:
```json
{
  "aliases": {
    "lofigirl": "https://www.youtube.com/watch?v=...",
    "jazz": "https://www.youtube.com/watch?v=..."
  }
}
```

### Playback History
`~/.config/clitube/history.json`:
```json
{
  "recent": [
    {
      "alias": "lofi",
      "url": "https://www.youtube.com/watch?v=...",
      "played_at": "2025-11-04T21:17:39.391000546-05:00",
      "video_mode": true
    }
  ]
}
```

The config directory is created automatically if it doesn't exist (with 0755 permissions). Only the last 3 played items are kept in history.

## External Dependencies

The application requires `mpv` to be installed on the system. The player.go file checks for mpv availability before attempting playback.

## Documentation Files

- **README.md**: User-facing documentation with installation and usage instructions
- **clitube.1**: Man page in standard man format (section 1 for user commands)
- **install.sh**: Installation script that builds the binary and installs both the executable and man page
- **CLAUDE.md**: This file - architecture and development guidance
