package upspin

import (
	"io"
	"strings"
	"time"

	"github.com/gildasch/upspin-music/album"
	"github.com/pkg/errors"
	"upspin.io/upspin"
)

type Accesser struct {
	upspin.Client
}

func (a *Accesser) List(path string) ([]*album.Album, error) {
	pattern := createPattern(path)

	entries, err := a.Glob(pattern)
	if err != nil {
		return nil, errors.Wrapf(err, "could not Glob pattern %q", pattern)
	}

	albums := []*album.Album{}
	currentAlbum := &album.Album{}

	for _, entry := range entries {
		if !entry.IsDir() {
			if !a.canRead(entry) {
				continue
			}
			currentAlbum.Add(string(entry.Name))
			continue
		}

		subAlbums, err := a.List(string(entry.Name))
		if err != nil {
			continue
		}
		albums = append(albums, subAlbums...)
	}

	if !currentAlbum.IsEmpty() {
		albums = append(albums, currentAlbum)
	}

	return albums, nil
}

func (a *Accesser) canRead(entry *upspin.DirEntry) bool {
	return len(entry.Blocks) > 0
}

func (a *Accesser) Get(path string) (io.ReadSeeker, string, time.Time, error) {
	upath := formatFilePath(path)

	de, err := a.Lookup(upath, true)
	if err != nil {
		return nil, "", time.Time{}, errors.Wrapf(err, "could not Lookup path %q", path)
	}

	f, err := a.Open(upath)
	if err != nil {
		return nil, "", time.Time{}, errors.Wrapf(err, "could not Open path %q", path)
	}

	return f, filename(de.Name), de.Time.Go(), nil
}

func createPattern(path string) string {
	pattern := strings.TrimPrefix(path, "/")

	if strings.Contains(pattern, "*") {
		return pattern
	}
	return strings.TrimSuffix(pattern, "/") + "/*"
}

func formatFilePath(path string) upspin.PathName {
	path = strings.TrimPrefix(path, "/")
	return upspin.PathName(path)
}

func filename(path upspin.PathName) string {
	splitted := strings.Split(string(path), "/")
	return splitted[len(splitted)-1]
}
