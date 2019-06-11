package phonebookrepo

import (
	phoneI "crudgolang/interface/phonebookinterface"
	phoneModel "crudgolang/model/phonebook"
	"errors"

	"github.com/jinzhu/gorm"
)

type dbConn struct {
	DbConn *gorm.DB
}

//NewPhoneBookRepo - initial repo block name
func NewPhoneBookRepo(DbConn *gorm.DB) phoneI.PhoneBookI {
	return &dbConn{DbConn}
}

//CreateData - create data in table phone_books
func (config *dbConn) CreateData(name, phone string) error {
	myData := phoneModel.PhoneBooks{
		Name:        name,
		PhoneNumber: phone,
	}

	err := config.DbConn.Create(&myData).Error
	if err != nil {
		return errors.New("CreateData err = " + err.Error())
	}

	return nil
}

//UpdateData - update data in table phone_books
func (config *dbConn) UpdateData(id int64, name, phone string) error {
	var myPhoneBook phoneModel.PhoneBooks
	errUpdate := config.DbConn.Model(&myPhoneBook).Where("id = ?", id).Updates(phoneModel.PhoneBooks{
		Name:        name,
		PhoneNumber: phone,
	}).Error

	if errUpdate != nil {
		return errors.New("UpdateData errUpdate = " + errUpdate.Error())
	}

	return nil
}

//DeleteData - soft delete data in table phone_books
func (config *dbConn) DeleteData(id int64) error {
	var myPhoneBook phoneModel.PhoneBooks
	errDelete := config.DbConn.Where("id = ?", id).Delete(&myPhoneBook).Error
	if errDelete != nil {
		return errors.New("DeleteData errDelete = " + errDelete.Error())
	}

	return nil
}

//FindAll - get all data in table phone_books
func (config *dbConn) FindAll() ([]phoneModel.PhoneBooks, error) {
	var list []phoneModel.PhoneBooks

	// query := `SELECT block_name_id,names FROM block_names WHERE deleted_at IS NULL`

	// rows, errQuery := config.DbConn.Query(query)

	// if errQuery != nil {
	// 	return list, errors.New("FindAll errQuery = " + errQuery.Error())
	// }

	// defer rows.Close()

	// for rows.Next() {
	// 	var data blockNameModel.ListBlockNames

	// 	errScan := rows.Scan(&data.BlockNameID, &data.Name)

	// 	if errScan != nil {
	// 		return list, errors.New("FindAll errScan = " + errScan.Error())
	// 	}

	// 	list = append(list, data)
	// }

	return list, nil
}

//FindByID - get data phone_books by id
func (config *dbConn) FindByID(id int64) (phoneModel.PhoneBooks, error) {
	var data phoneModel.PhoneBooks
	// query := `SELECT block_name_id,names FROM block_names WHERE block_name_id = ?`
	// stat, errPrepare := config.DbConn.Prepare(query)
	// if errPrepare != nil {
	// 	return data, errors.New("FindByID errPrepare = " + errPrepare.Error())
	// }

	// defer stat.Close()
	// errQueryRow := stat.QueryRow(id).Scan(&data.BlockNameID, &data.Name)
	// if errQueryRow != nil {
	// 	return data, errors.New("FindByID errQueryRow = " + errQueryRow.Error())
	// }

	return data, nil
}
