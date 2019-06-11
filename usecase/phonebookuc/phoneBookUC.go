package phonebookuc

import (
	phoneI "crudgolang/interface/phonebookinterface"
	phoneUCI "crudgolang/interface/phonebookucinterface"

	phoneModel "crudgolang/model/phonebook"
)

//phoneBookUseCase - implement PhoneBookInterface
type phoneBookUseCase struct {
	phoneI.PhoneBookI
}

//NewPhoneBookUseCase - initial phone book use case
func NewPhoneBookUseCase(phoneInterface phoneI.PhoneBookI) phoneUCI.PhoneBookUCI {
	return &phoneBookUseCase{
		phoneInterface,
	}
}

//GetAll - get data list from table phone_books
func (phoneBookUC *phoneBookUseCase) GetAll() ([]phoneModel.Datas, error) {
	var (
		list    []phoneModel.PhoneBooks
		myDatas []phoneModel.Datas
	)

	list, err := phoneBookUC.PhoneBookI.FindAll()
	if err != nil {
		return myDatas, err
	}

	for _, val := range list {
		var detail phoneModel.Datas
		detail.Name = val.Name
		detail.PhoneNumber = val.PhoneNumber
		myDatas = append(myDatas, detail)
	}

	return myDatas, nil
}

//GetByID - get data from table phone_books by id
func (phoneBookUC *phoneBookUseCase) GetByID(id int64) (phoneModel.Datas, error) {
	var (
		data   phoneModel.PhoneBooks
		myData phoneModel.Datas
	)

	data, err := phoneBookUC.PhoneBookI.FindByID(id)
	if err != nil {
		return myData, err
	}

	myData.Name = data.Name
	myData.PhoneNumber = data.PhoneNumber

	return myData, nil
}

//InsertData - insert data into table phone_books
func (phoneBookUC *phoneBookUseCase) InsertData(name, phone string) error {
	err := phoneBookUC.PhoneBookI.CreateData(name, phone)

	if err != nil {
		return err
	}

	return nil
}

//EditData - edit data from table phone_books by id
func (phoneBookUC *phoneBookUseCase) EditData(id int64, name, phone string) error {
	err := phoneBookUC.PhoneBookI.UpdateData(id, name, phone)

	if err != nil {
		return err
	}

	return nil
}

//DeleteData - delete data from table phone_books by id
func (phoneBookUC *phoneBookUseCase) DeleteData(id int64) error {
	err := phoneBookUC.PhoneBookI.DeleteData(id)
	if err != nil {
		return err
	}

	return nil
}
