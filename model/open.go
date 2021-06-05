package model

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ShunyaNagashige/ca-tech-dojo-golang/config"
	_ "github.com/go-sql-driver/mysql"
)

type DbError struct {
	Cmd string
	Err error
}

func (err *DbError) Error() string {
	return err.Err.Error()
}

func createDsn() string {
	return config.Config.DbUserName +
		":" + config.Config.DbPassword +
		"@" + config.Config.DbProtocol +
		"(" + config.Config.DbAddress +
		":" + config.Config.DbPort + ")" +
		"/" + config.Config.DbName
}

const (
	tableNameUsers          = "users"
	tableNameCharacters     = "characters"
	tableNameUserCharacters = "user_characters"
)

var dbConn *sql.DB

func init() {
	var err error

	fmt.Println(createDsn())

	//DBに接続(DBのオープン)
	dbConn, err = sql.Open(config.Config.SqlDriver, createDsn())
	if err != nil {
		log.Fatalf("Failed to open a database: %v", err)
	}
}
