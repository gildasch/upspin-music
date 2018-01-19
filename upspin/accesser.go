package upspin

import (
	"io"
	"strings"

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
	f, err := a.Open(entry.Name)
	if err != nil {
		// we cannot access this file
		return false
	}
	f.Close()
	return true
}

func (a *Accesser) Get(path string) (io.Reader, error) {
	upath := formatFilePath(path)

	f, err := a.Open(upath)
	if err != nil {
		return nil, errors.Wrapf(err, "could not Open path %q", path)
	}

	return f, nil
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
