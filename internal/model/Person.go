package model

// Person representa uma pessoa no sistema
type Person struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
