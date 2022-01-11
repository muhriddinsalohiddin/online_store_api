package models

type Book struct {
	Name        string   `json:"name"`
	AuthorId    string   `json:"author_id"`
	CategoryIds []string `json:"category_id"`
}

type UpdateBook struct {
	Name        string   `json:"name"`
	AuthorId    string   `json:"author_id"`
	CategoryIds []string `json:"category_id"`
}

type BookById struct {
	ID string `json:"id"`
}

type ListBooks struct {
	Books []Book `json:"Books"`
}
