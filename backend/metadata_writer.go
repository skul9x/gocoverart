package backend

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bogem/id3v2/v2"
	"github.com/go-flac/flacvorbis"
	"github.com/go-flac/go-flac"
)

// UpdateMetadataTags writes new metadata to audio file
func UpdateMetadataTags(filePath, title, artist, album, year, genre, comment string) error {
	// Check file is writable
	f, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("file is locked or read-only: %w", err)
	}
	f.Close()

	ext := strings.ToLower(filepath.Ext(filePath))
	switch ext {
	case ".mp3":
		return updateMP3Metadata(filePath, title, artist, album, year, genre, comment)
	case ".flac":
		return updateFLACMetadata(filePath, title, artist, album, year, genre, comment)
	default:
		return fmt.Errorf("unsupported format: %s", ext)
	}
}

func updateMP3Metadata(filePath, title, artist, album, year, genre, comment string) error {
	tag, err := id3v2.Open(filePath, id3v2.Options{Parse: true})
	if err != nil {
		return fmt.Errorf("could not open mp3 tag: %w", err)
	}
	defer tag.Close()

	// Required field
	tag.SetTitle(title)

	// Optional fields - only set if not empty
	if artist != "" {
		tag.SetArtist(artist)
	}
	if album != "" {
		tag.SetAlbum(album)
	}
	if year != "" {
		tag.SetYear(year)
	}
	if genre != "" {
		tag.SetGenre(genre)
	}
	if comment != "" {
		// Clear existing comments first to overwrite
		tag.DeleteFrames(tag.CommonID("Comments"))
		tag.AddCommentFrame(id3v2.CommentFrame{
			Encoding:    id3v2.EncodingUTF8,
			Language:    "eng",
			Description: "",
			Text:        comment,
		})
	}

	return tag.Save()
}

func updateFLACMetadata(filePath, title, artist, album, year, genre, comment string) error {
	f, err := flac.ParseFile(filePath)
	if err != nil {
		return fmt.Errorf("could not parse flac file: %w", err)
	}

	var vorbisData *flacvorbis.MetaDataBlockVorbisComment
	var vorbisIndex int = -1

	for i, meta := range f.Meta {
		if meta.Type == flac.VorbisComment {
			temp, err := flacvorbis.ParseFromMetaDataBlock(*meta)
			if err == nil {
				vorbisData = temp
				vorbisIndex = i
			}
			break
		}
	}

	if vorbisData == nil {
		vorbisData = flacvorbis.New()
	}

	// Helper to update or add fields
	updateField := func(v *flacvorbis.MetaDataBlockVorbisComment, key, value string) {
		if value == "" {
			return
		}
		// Clear existing entries for this key
		var newComments []string
		prefix := strings.ToUpper(key) + "="
		for _, c := range v.Comments {
			if !strings.HasPrefix(strings.ToUpper(c), prefix) {
				newComments = append(newComments, c)
			}
		}
		v.Comments = newComments
		v.Add(key, value)
	}

	updateField(vorbisData, "TITLE", title)
	updateField(vorbisData, "ARTIST", artist)
	updateField(vorbisData, "ALBUM", album)
	updateField(vorbisData, "DATE", year)
	updateField(vorbisData, "GENRE", genre)
	updateField(vorbisData, "COMMENT", comment)

	res := vorbisData.Marshal()
	if vorbisIndex >= 0 {
		f.Meta[vorbisIndex] = &res
	} else {
		f.Meta = append(f.Meta, &res)
	}

	return f.Save(filePath)
}
