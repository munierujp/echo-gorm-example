package model

type User struct {
	Model
	Name       string   `json:"name"`
	LanguageID uint     `json:"language_id"`
	Language   Language `json:"-"`
}
