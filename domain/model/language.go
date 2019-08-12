package model

type Language struct {
	Model
	Code string `json:"code"`
	Name string `json:"name"`
}
