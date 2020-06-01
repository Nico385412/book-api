package service

import (
	"log"

	"github.com/nico385412/book-api/models"
	"github.com/nico385412/goreader/epub"
)

//ConvertFileToBookModel take a ebook and extract metadata
func ConvertFileToBookModel(filename *string, coverId *string) *models.Book {
	rc, err := epub.OpenReader(*filename)
	if err != nil {
		log.Fatal("the epub seems not be in the good format please fix him")
	}
	defer rc.Close()

	book := models.Book{
		ID:          filename,
		Title:       rc.Rootfiles[0].Metadata.Title,
		Language:    rc.Rootfiles[0].Metadata.Language,
		Identifier:  rc.Rootfiles[0].Metadata.Identifier,
		Creator:     rc.Rootfiles[0].Metadata.Creator,
		Contributor: rc.Rootfiles[0].Metadata.Contributor,
		Publisher:   rc.Rootfiles[0].Metadata.Publisher,
		Subject:     rc.Rootfiles[0].Metadata.Subject,
		Description: rc.Rootfiles[0].Metadata.Description,
		CoverID:     *coverId,
	}

	return &book
}

func GetCover(filename *string) ([]byte, error) {
	rc, err := epub.OpenReader(*filename)
	if err != nil {
		log.Fatal("the epub seems not be in the good format please fix him")
	}
	defer rc.Close()

	return rc.Reader.GetCoverBytes()
}
