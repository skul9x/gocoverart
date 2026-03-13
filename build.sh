#!/bin/bash

# Build script for Music Cover Tagger
# Creates .deb for Linux and .exe for Windows

set -e

echo "🚀 Starting build process..."

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Get version from CHANGELOG.md or use default
VERSION="1.0.0"
APP_NAME="music-cover-tagger"

echo -e "${BLUE}📦 Building for Linux (amd64)...${NC}"
wails build -platform linux/amd64 -clean

echo -e "${BLUE}📦 Building for Windows (amd64)...${NC}"
wails build -platform windows/amd64

# Create .deb package
echo -e "${BLUE}📦 Creating .deb package...${NC}"

DEB_DIR="build/deb"
mkdir -p ${DEB_DIR}/DEBIAN
mkdir -p ${DEB_DIR}/usr/local/bin
mkdir -p ${DEB_DIR}/usr/share/applications
mkdir -p ${DEB_DIR}/usr/share/icons/hicolor/256x256/apps

# Copy binary
cp build/bin/${APP_NAME} ${DEB_DIR}/usr/local/bin/

# Create control file
cat > ${DEB_DIR}/DEBIAN/control << EOF
Package: ${APP_NAME}
Version: ${VERSION}
Section: sound
Priority: optional
Architecture: amd64
Maintainer: skul9x <skul9x@gmail.com>
Description: Music Cover Tagger
 Desktop application to find and apply album covers to your music collection.
 Supports MP3, FLAC, WAV and other popular audio formats.
EOF

# Create desktop entry
cat > ${DEB_DIR}/usr/share/applications/${APP_NAME}.desktop << EOF
[Desktop Entry]
Name=Music Cover Tagger
Comment=Find and apply album covers to your music
Exec=/usr/local/bin/${APP_NAME}
Icon=${APP_NAME}
Terminal=false
Type=Application
Categories=Audio;AudioVideo;
EOF

# Copy icon if exists
if [ -f "build/appicon.png" ]; then
    cp build/appicon.png ${DEB_DIR}/usr/share/icons/hicolor/256x256/apps/${APP_NAME}.png
fi

# Build .deb
dpkg-deb --build ${DEB_DIR} build/bin/${APP_NAME}_${VERSION}_amd64.deb

echo -e "${GREEN}✅ Build complete!${NC}"
echo ""
echo "📦 Output files:"
echo "  - Linux: build/bin/${APP_NAME}_${VERSION}_amd64.deb"
echo "  - Windows: build/bin/${APP_NAME}.exe"
echo ""
echo "🎉 Ready to upload to GitHub Releases!"
