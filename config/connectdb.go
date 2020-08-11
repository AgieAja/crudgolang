package config

// coba commit ( abbas )
import (
	"fmt"

	"github.com/jinzhu/gorm"

	//driver mysql database
	_ "github.com/go-sql-driver/mysql"
)

var (
	mydb   *gorm.DB
	sqlErr error
)

//InitConnMySQLDB - preparation connection database mysql
func InitConnMySQLDB() {
	dbHost := "127.0.0.1"
	dbPort := "3306"
	dbUser := "root"
	dbPass := "tiara1!@#"
	dbName := "test"

	desc := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	mydb, sqlErr = createConnMySQL(desc)
}

//GetMySQLDB - get connection db mysql
func GetMySQLDB() (*gorm.DB, error) {
	return mydb, sqlErr
}

//createConnMySQL - create connection database mysql
func createConnMySQL(desc string) (*gorm.DB, error) {
	mydb, sqlErr = gorm.Open("mysql", desc)
	if sqlErr != nil {
		return nil, sqlErr
	}

	sqlErr = mydb.DB().Ping()
	if sqlErr != nil {
		return nil, sqlErr
	}

	mydb.DB().SetMaxIdleConns(10)
	mydb.DB().SetMaxOpenConns(10)
	return mydb, nil
}
