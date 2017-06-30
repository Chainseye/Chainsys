package model

import (
	"fmt"

	"github.com/gorm"
	_ "github.com/gorm/dialects/mysql"
)

const DEFAULT_MODEL = "default"

var dbConnection *DBConnection

func init() {
	dbConnection = new(DBConnection)
	dbConnection.dbs = make(map[string]*gorm.DB)

}

type DBConnection struct {
	dbs map[string]*gorm.DB
}

func RegisterDefaultDB(connStr string) error {
	fmt.Println(connStr)
	return RegisterDB(DEFAULT_MODEL, connStr)
}

func RegisterDB(name, connStr string) error {
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		return err
	}
	dbConnection.dbs[name] = db
	return nil
}

func GetDB(name string) *gorm.DB {
	return dbConnection.dbs[name]
}

func (this *BaseModel) Using(name string) {
	this.DB = GetDB(name)
}

func Close() {
	for _, db := range dbConnection.dbs {
		db.Close()
	}
}

type BaseModel struct {
	DB *gorm.DB
}
