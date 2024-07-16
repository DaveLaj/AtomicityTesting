package models

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type CreateUser struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
