package models

//Book model
type Book struct {
	BinaryID    interface{} `json:"binary_id"`
	Title       string      `json:"title"`
	Language    string      `json:"languague"`
	Identifier  string      `json:"identifier"`
	Creator     string      `json:"creator"`
	Contributor string      `json:"contributor"`
	Publisher   string      `json:"publisher"`
	Subject     string      `json:"subject"`
	Description string      `json:"description"`
	Cover       string      `json:"cover"`
}
