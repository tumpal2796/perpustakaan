package model

type Filter struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

type BookResponse struct {
	Book
	Author []Author `json:"author"`
	Publisher
}
