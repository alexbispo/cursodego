package models

// Expense represents a expense
type Expense struct {
	ID    string `json:"_id"`
	Name  string `json:"name"`
	Value int    `json:"amount"`
}
