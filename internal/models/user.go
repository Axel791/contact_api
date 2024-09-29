package models

type User struct {
	ID          uint `gorm:"primaryKey"`
	FirstName   string
	LastName    string
	Age         int
	Description string
	Contacts    []Contact `gorm:"foreignKey:UserID"`
}
