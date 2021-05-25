/*
データベーススキーマの生成
*/
package model

import (
	//ドライバの登録
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const(
	dsn="homestead:secret@TCP(localhost:13306)/homestead"
	driver="mysql"
)

var dbConn *sql.DB

func init(){
	var err error
	
	dbConn,err=sql.Open(driver,dsn)
	if err!=nil{
		log.Fatalln("sql.Open に失敗しました．")
	}

	
}