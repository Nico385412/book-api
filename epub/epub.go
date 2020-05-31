package epub

import (
	"archive/zip"
	"io"
)

type Item struct {
	ID        string `xml:"id,attr"`
	HREF      string `xml:"href,attr"`
	MediaType string `xml:"media-type,attr"`
	f         *zip.File
}

func (item *Item) Open(r io.ReadCloser, err error) {
	if item.f == nil {
		return nil, ErrBadManifest
	}

	return item.f.Open()
}
