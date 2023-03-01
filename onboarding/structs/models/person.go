package models

type Person struct {
	ID       string `json:"id" xml:"id"`
	Name     string `json:"name" xml:"name"`
	LastName string `json:"last_name" xml:"last_name"`
	Age      int    `json:"age" xml:"age"`
	IsActive bool   `json:"is_active" xml:"is_active"`
}
