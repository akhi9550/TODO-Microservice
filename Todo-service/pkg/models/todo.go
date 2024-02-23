package models

type Data struct {
	ID          int64  `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
