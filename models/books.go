package models

type Book struct {
    ID     uint   `json:"id" gorm:"primary_key"`
    Title  string `json:"title"`
    Author string `json:"author"`
    Rating float32 `json:"rating"`
}

type CreateBookInput struct {
    Title string `json:"title" binding:"required"`
    Author string `json:"author" binding:"required"`
    Rating float32 `json:"rating,omitempty"`
}

type UpdateBookInput struct {
    Title string `json:"title"`
    Author string `json:"author"`
    Rating float32 `json:"rating"`
}
