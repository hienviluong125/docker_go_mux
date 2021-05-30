package models

type Post struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
