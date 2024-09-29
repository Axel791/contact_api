package models

type Contact struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	Type   string
	Value  string
}
