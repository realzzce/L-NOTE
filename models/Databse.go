package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var GlobalDB *gorm.DB

// InitConnectDB 初始化db
func InitConnectDB() (*gorm.DB, error) {

	mysqlInfo := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"", "", "localhost", 3306, "LNOTE")

	globalDB, err := gorm.Open("mysql", mysqlInfo)

	if err != nil {
		panic("Fail connect to DB")
	}

	GlobalDB = globalDB
	GlobalDB.SingularTable(true)

	GlobalDB.AutoMigrate(&UserModel{})
	GlobalDB.AutoMigrate(&ClaimsModel{})
	GlobalDB.AutoMigrate(&NotesModel{})

	return GlobalDB, nil
}
