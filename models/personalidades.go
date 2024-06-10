package models

type Personalidade struct {
	ID int `json:"id"`
	Nome string `json:"nome"`
	Descricao string `json:"descricao"`
}

var Personalidades []Personalidade
