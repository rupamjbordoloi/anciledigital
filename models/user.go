package models

type User struct {
	IDColumn
	Name     string  `gorm:"not null" json:"name"`
	Email    string  `gorm:"not null;unique" json:"email"`
	MobileNo *string `json:"mobile_no"`
	DateColumn
}
