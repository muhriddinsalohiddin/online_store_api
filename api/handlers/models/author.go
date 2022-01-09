package models

type Author struct {
	Name string `json:"name"`
}

type ListAuthors struct {
	Authors []Author `json:"authors"`
}
