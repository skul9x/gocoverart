package backend

import (

	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/bogem/id3v2/v2"
	"github.com/go-flac/flacpicture"
	"github.com/go-flac/go-flac"
)

// EmbedCoverArt downloads image and embeds it into the audio file
// force controls whether existing artwork should be replaced when present
func EmbedCoverArt(filePath, imageURL string, force bool) error {
	// Pre-flight: Check if file is writable
	f, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("file is locked or read-only: %w", err)
	}
	f.Close()

	// 1. Download image
	resp, err := http.Get(imageURL)
	if err != nil {
		return fmt.Errorf("failed to download image: %w", err)
	}
	defer resp.Body.Close()

	imgData, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read image data: %w", err)
	}

	ext := strings.ToLower(filepath.Ext(filePath))
	switch ext {
	case ".mp3":
		return embedMP3(filePath, imgData, force)
	case ".flac":
		return embedFLAC(filePath, imgData, force)
	default:
		return fmt.Errorf("unsupported format: %s", ext)
	}
}

func embedMP3(filePath string, imgData []byte, force bool) error {
	tag, err := id3v2.Open(filePath, id3v2.Options{Parse: true})
	if err != nil {
		return fmt.Errorf("could not open mp3 tag: %w", err)
	}
	defer tag.Close()

	if force {
		// Delete existing attached picture frames (APIC / "Attached picture")
		// id3v2 stores APIC frames; DeleteFrames accepts frame IDs - use the CommonID
		tag.DeleteFrames(tag.CommonID("Attached picture"))
	}

	pic := id3v2.PictureFrame{
		Encoding:    id3v2.EncodingUTF8,
		MimeType:    "image/jpeg",
		PictureType: id3v2.PTFrontCover,
		Description: "Front Cover",
		Picture:     imgData,
	}
	// Always add the new picture
	tag.AddAttachedPicture(pic)

	return tag.Save()
}

func embedFLAC(filePath string, imgData []byte, force bool) error {
	f, err := flac.ParseFile(filePath)
	if err != nil {
		return fmt.Errorf("could not parse flac file: %w", err)
	}

	picture, err := flacpicture.NewFromImageData(
		flacpicture.PictureTypeFrontCover,
		"Front Cover",
		imgData,
		"image/jpeg",
	)
	if err != nil {
		return fmt.Errorf("could not create flac picture frame: %w", err)
	}

	pictureMeta := picture.Marshal()

	if force {
		// Replace any existing picture metadata blocks of type FrontCover
		var newMeta []*flac.MetaDataBlock
		for _, m := range f.Meta {
			// naive detection: skip picture-like meta blocks
			// Ideally check block type; here we filter out blocks that contain "PICTURE" in their string
			if strings.Contains(strings.ToUpper(string(m.Marshal(false))), "PICTURE") {
				continue
			}
			newMeta = append(newMeta, m)
		}
		newMeta = append(newMeta, &pictureMeta)
		f.Meta = newMeta
	} else {
		f.Meta = append(f.Meta, &pictureMeta)
	}

	return f.Save(filePath)
}
