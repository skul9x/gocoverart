package backend

import (
	"context"
)

// TrackMetadata holds music file info
type TrackMetadata struct {
	FilePath  string `json:"filePath"`
	Title     string `json:"title"`
	Artist    string `json:"artist"`
	Album     string `json:"album"`
	HasCover  bool   `json:"hasCover"`
	CoverData string `json:"coverData"`
}

// CoverResult holds found image info
type CoverResult struct {
	URL      string `json:"url"`
	Preview  string `json:"preview"`
	Rank     int    `json:"rank"`
}

// Core handles business logic for the app
type Core struct {
	ctx context.Context
}

// NewCore initializes the backend logic
func NewCore(ctx context.Context) *Core {
	return &Core{ctx: ctx}
}

// GetMusicFiles scans a directory and extracts metadata for all supported files
func (c *Core) GetMusicFiles(dirPath string) ([]TrackMetadata, error) {
	paths, err := ScanMusicFiles(dirPath)
	if err != nil {
		return nil, err
	}

	var results []TrackMetadata
	for _, p := range paths {
		meta, err := ExtractMetadata(p)
		if err != nil {
			continue // Skip errors for individual files
		}
		results = append(results, *meta)
	}

	return results, nil
}

// SearchCovers finds cover art for a specific track
func (c *Core) SearchCovers(query string) ([]CoverResult, error) {
	return SearchGoogleImage(query)
}

// ApplyCover downloads and embeds the cover art into the file
// Updated to accept a force flag to allow overwriting existing artwork.
func (c *Core) ApplyCover(filePath, imageURL string, force bool) error {
	return EmbedCoverArt(filePath, imageURL, force)
}
