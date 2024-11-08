package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID    string `json:"userid" gorm:"unique;not null;min=15;max=60"`
	Password  string `json:"password" gorm:"unique;not null"`
	PublicKey string `json:"public_key" gorm:"type:text;not null"`
}
