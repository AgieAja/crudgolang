package phonebook

import (
	"github.com/jinzhu/gorm"
)

//PhoneBooks - maping table phone_books
type PhoneBooks struct {
	gorm.Model
	Name        string
	PhoneNumber string
}
