package models

type Request struct {
	Nums []int `json:"nums"`
}

type IdRequest struct {
	Id int `json:"Id"`
}

type Book struct {
	Id           int    `json:"Id"`
	Author_name  string `json:"author_name"`
	Book_title   string `json:"book_title"`
	Release_year int    `json:"release_year"`
}
