package phonebookinterface

import (
	phoneModel "crudgolang/model/phonebook"
)

//PhoneBookI - phone book interface
type PhoneBookI interface {
	CreateData(string, string) error
	UpdateData(int64, string, string) error
	DeleteData(int64) error
	FindAll() ([]phoneModel.PhoneBooks, error)
	FindByID(int64) (phoneModel.PhoneBooks, error)
}
