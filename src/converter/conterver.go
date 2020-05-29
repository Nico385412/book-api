package converter

import "github.com/nico385412/book-api/models"
import "github.com/nico385412/goreader/epub"
import "log"

//ConvertFileToBookModel take a ebook and extract metadata
func ConvertFileToBookModel(filename *string) *models.Book {
	rc, err := epub.OpenReader(*filename)
	if err != nil {
		log.Fatal("the epub seems not be in the good format please fix him")
	}
	defer rc.Close()

	coverBase64, err := rc.Reader.GetCoverBase64()

	book := models.Book{
		BinaryID:    filename,
		Title:       rc.Rootfiles[0].Metadata.Title,
		Language:    rc.Rootfiles[0].Metadata.Language,
		Identifier:  rc.Rootfiles[0].Metadata.Identifier,
		Creator:     rc.Rootfiles[0].Metadata.Creator,
		Contributor: rc.Rootfiles[0].Metadata.Contributor,
		Publisher:   rc.Rootfiles[0].Metadata.Publisher,
		Subject:     rc.Rootfiles[0].Metadata.Subject,
		Description: rc.Rootfiles[0].Metadata.Description,
		Cover:       coverBase64,
	}

	return &book

}
