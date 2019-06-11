package phonebook

import (
	"github.com/jinzhu/gorm"
)

//PhoneBooks - maping table phone_books
type PhoneBooks struct {
	gorm.Model
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}
