package controller

// import (
// 	"database/sql"
// 	"fmt"
// )

// 数据库信息
// const (
// 	DB_Admin     = "root"
// 	DB_Password  = "a123123"
// 	DB_ConnectDB = "GoLang_DB"
// )

// func Mysql_Driver() {
// 	db, err := sql.Open("mysql", DB_Admin+":"+DB_Password+"@/"+DB_ConnectDB+"?charset=utf8")
// 	CheckErr(err)
// 	rows, err := db.Query("select * from person")
// 	CheckErr(err)
// 	for rows.Next() {
// 		var ID int
// 		var Name string
// 		var Password string
// 		err = rows.Scan(&ID, &Name, &Password)
// 		CheckErr(err)
// 		fmt.Println(ID)
// 		fmt.Println(Name)
// 		fmt.Println(Password)
// 	}
// 	db.Close()
// }
