package main

import (
	"context"
	"fmt"
	"time"

	"temp-app/backend"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx  context.Context
	core *backend.Core
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts.
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.core = backend.NewCore(ctx)
}

// SelectDirectory opens a directory dialog and returns the selected path
func (a *App) SelectDirectory() (string, error) {
	selection, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Chọn thư mục chứa nhạc",
	})
	if err != nil {
		return "", err
	}
	return selection, nil
}

// GetTracks scans a directory and extracts metadata for all supported files
func (a *App) GetTracks(dirPath string) ([]backend.TrackMetadata, error) {
	return a.core.GetMusicFiles(dirPath)
}

// SearchCovers finds cover art for a specific track
func (a *App) SearchCovers(query string) ([]backend.CoverResult, error) {
	return a.core.SearchCovers(query)
}

// ApplyCover downloads and embeds the cover art into the file
func (a *App) ApplyCover(filePath, imageURL string, force bool) error {
	err := a.core.ApplyCover(filePath, imageURL, force)
	if err == nil {
		// Re-extract metadata to get the new cover image as base64
		meta, metaErr := backend.ExtractMetadata(filePath)
		if metaErr == nil {
			runtime.EventsEmit(a.ctx, "track_updated", meta)
		} else {
			// Fallback if extraction fails but embedding was successful
			runtime.EventsEmit(a.ctx, "track_updated", backend.TrackMetadata{
				FilePath: filePath,
				HasCover: true,
			})
		}
	}
	return err
}

// UpdateMetadata updates the metadata tags of an audio file
func (a *App) UpdateMetadata(filePath, title, artist, album, year, genre, comment string) (*backend.TrackMetadata, error) {
	err := backend.UpdateMetadataTags(filePath, title, artist, album, year, genre, comment)
	if err != nil {
		return nil, err
	}

	// Re-extract metadata to return updated values and notify UI
	meta, err := backend.ExtractMetadata(filePath)
	if err == nil {
		runtime.EventsEmit(a.ctx, "track_updated", meta)
	}
	return meta, err
}

// StartAutoTagging automatically finds and applies the first cover result for all tracks in a list
func (a *App) StartAutoTagging(tracks []backend.TrackMetadata) {
	go func() {
		for _, track := range tracks {
			if track.HasCover {
				continue
			}

			// Notify UI we are starting this track
			runtime.EventsEmit(a.ctx, "track_processing", track.FilePath)

			// 1. Search
			results, err := a.core.SearchCovers(fmt.Sprintf("%s %s", track.Artist, track.Title))
			if err != nil || len(results) == 0 {
				runtime.EventsEmit(a.ctx, "track_error", track.FilePath)
				continue
			}

			// 2. Apply Top 1
			err = a.core.ApplyCover(track.FilePath, results[0].URL, false)
			if err != nil {
				runtime.EventsEmit(a.ctx, "track_error", track.FilePath)
			} else {
				// Re-extract to get high-res base64 for UI update
				meta, metaErr := backend.ExtractMetadata(track.FilePath)
				if metaErr == nil {
					runtime.EventsEmit(a.ctx, "track_updated", meta)
				} else {
					runtime.EventsEmit(a.ctx, "track_updated", backend.TrackMetadata{
						FilePath: track.FilePath,
						HasCover: true,
					})
				}
			}

			// Small delay to be polite to Google
			time.Sleep(800 * time.Millisecond)
		}
		runtime.EventsEmit(a.ctx, "batch_finished", true)
	}()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, chào mừng anh đến với Music Cover Tagger!", name)
}
