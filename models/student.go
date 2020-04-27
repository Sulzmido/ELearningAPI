package models

// Student model
type Student struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    string `json:"age"`
}
