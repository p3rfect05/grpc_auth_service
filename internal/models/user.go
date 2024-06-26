package models

type User struct {
	ID           int64  `gorn:"primaryKey"`
	Email        string `gorm:"primaryKey"`
	PasswordHash []byte
}
