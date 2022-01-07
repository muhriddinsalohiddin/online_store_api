package models

type Book struct {
	ID         	 string `json:"id"`
	Name 	   	 string `json:"name"`
	AuthorId   	 string `json:"author_id"`
	CategoryId 	 string `json:"category_id"`
	CreatedAt  	 string `json:"created_at"`
	UpdatedAt  	 string `json:"updated_at"`
}

type UpdateBook struct {
	Name 	   	 string `json:"name"`
	AuthorId   	 string `json:"author_id"`
	UpdatedAt  	 string `json:"updated_at"`
}

type BookById struct {
	ID string `json:"id"`
}

type ListBooks struct {
	Books []Book `json:"Books"`
}
