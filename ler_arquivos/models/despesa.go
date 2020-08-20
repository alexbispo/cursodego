package models

// Despesa representa uma despesa
type Despesa struct {
	Nome  string `json:"nome"`
	Valor int    `json:"valor"`
}
