package album

import "strings"

type Album struct {
	Songs []string
	Cover string
}

func (a *Album) Add(path string) {
	if isSong(path) {
		a.Songs = append(a.Songs, path)
	}
	if isCover(path) && a.Cover == "" {
		a.Cover = path
	}
}

func (a *Album) IsEmpty() bool {
	return len(a.Songs) == 0
}

var songExtensions = []string{
	"mp3", "flac",
}

func isSong(name string) bool {
	name = strings.ToLower(name)

	for _, ext := range songExtensions {
		if strings.HasSuffix(name, "."+ext) {
			return true
		}
	}
	return false
}

var coverExtensions = []string{
	"jpg", "jpeg",
	"png",
	"gif",
	"bmp",
}

func isCover(name string) bool {
	name = strings.ToLower(name)

	for _, ext := range coverExtensions {
		if strings.HasSuffix(name, "."+ext) {
			return true
		}
	}
	return false
}
