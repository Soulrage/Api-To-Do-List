package models

type Tasks struct {
	ID          uint    `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string `gorm:""`
	DueDate     string `gorm:"not null"`
	CreateAt    string `gorm:"not null"`
	UpdateAt    string `gorm:""`
}
