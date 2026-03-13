#!/bin/bash

# GitHub Release Script
# Usage: ./release.sh v1.0.0 "Release notes here"

set -e

VERSION=$1
NOTES=$2

if [ -z "$VERSION" ]; then
    echo "❌ Error: Version required"
    echo "Usage: ./release.sh v1.0.0 \"Release notes\""
    exit 1
fi

if [ -z "$NOTES" ]; then
    NOTES="Release $VERSION"
fi

echo "🚀 Creating GitHub Release: $VERSION"

# Check if gh CLI is installed
if ! command -v gh &> /dev/null; then
    echo "❌ GitHub CLI (gh) not installed"
    echo "Install: sudo apt install gh"
    exit 1
fi

# Create git tag
git tag -a $VERSION -m "$NOTES"
git push origin $VERSION

# Create GitHub release and upload files
gh release create $VERSION \
    build/bin/music-cover-tagger_1.0.0_amd64.deb \
    build/bin/music-cover-tagger.exe \
    --title "$VERSION" \
    --notes "$NOTES"

echo "✅ Release created successfully!"
echo "🔗 https://github.com/skul9x/gocoverart/releases/tag/$VERSION"
