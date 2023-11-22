package models

type Book struct {
	ID     int    `json:"id"`
	Genre  string `json:"genre"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
