package backend

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dhowden/tag"
)

// ExtractMetadata reads Artist, Title, Album and checks if cover art already exists
func ExtractMetadata(filePath string) (*TrackMetadata, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	m, err := tag.ReadFrom(f)
	title := m.Title()
	if title == "" {
		// Fallback to filename without extension
		title = filepath.Base(filePath)
		ext := filepath.Ext(title)
		title = strings.TrimSuffix(title, ext)
	}

	artist := m.Artist()
	if artist == "" {
		artist = "N/A"
	}

	album := m.Album()
	if album == "" {
		album = "N/A"
	}

	hasCover := false
	coverData := ""
	if p := m.Picture(); p != nil {
		hasCover = true
		base64Data := base64.StdEncoding.EncodeToString(p.Data)
		coverData = fmt.Sprintf("data:%s;base64,%s", p.MIMEType, base64Data)
	}

	return &TrackMetadata{
		FilePath: filePath,
		Title:    title,
		Artist:   artist,
		Album:    album,
		HasCover:  hasCover,
		CoverData: coverData,
	}, nil
}
