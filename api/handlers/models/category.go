package models

type Category struct {
	Name     string `json:"name"`
	ParentId string `json:"parent_id"`
}

type ListCategories struct {
	Categories []Category `json:"categories"`
}
