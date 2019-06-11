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

//Datas - struct for json
type Datas struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone"`
}
