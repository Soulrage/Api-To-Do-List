package models


type Users struct {
	ID       uint    `gorm:"primaryKey"`
	Login    string  `gorm:"not null"`
	Password string  `gorm:"not null"`
	Email    string  `gorm:""`
	Tasks    []Tasks `gorm:"foreignKey:UserID"`
}

