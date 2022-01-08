package models

type Order struct {
	BookId      string `json:"book_id"`
	Description string `json:"description"`
}

type ListOrders struct {
	Orders []Order `json:"orders"`
}
