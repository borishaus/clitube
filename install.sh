#!/bin/bash

# cliTube installation script

set -e

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}Installing cliTube...${NC}"

# Build the binary
echo "Building cliTube..."
go build -o clitube

# Install binary
echo "Installing binary to /usr/local/bin..."
sudo mv clitube /usr/local/bin/

# Install man page
echo "Installing man page..."
sudo mkdir -p /usr/local/share/man/man1
sudo cp clitube.1 /usr/local/share/man/man1/
sudo mandb > /dev/null 2>&1 || true

echo -e "${GREEN}✓ Installation complete!${NC}"
echo ""
echo "Try it out:"
echo "  clitube help"
echo "  man clitube"
echo ""
echo "Get started by adding your first video:"
echo "  clitube add lofi \"https://www.youtube.com/watch?v=jfKfPfyJRdk\""
