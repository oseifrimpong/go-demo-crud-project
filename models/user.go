package models

import (
	"github.com/jinzhu/gorm"
	// "golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}
