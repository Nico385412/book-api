package models

//Book model
type Book struct {
	ID          interface{} `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string      `json:"title"`
	Language    string      `json:"language"`
	Identifier  string      `json:"identifier"`
	Creator     string      `json:"creator"`
	Contributor string      `json:"contributor"`
	Publisher   string      `json:"publisher"`
	Subject     string      `json:"subject"`
	Description string      `json:"description"`
	CoverID     string      `json:"cover"`
}
