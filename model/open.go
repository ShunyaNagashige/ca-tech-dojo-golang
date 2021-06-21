package model

import (
	"database/sql"
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

func NewDbError(cmd string, err error) *DbError {
	return &DbError{cmd, err}
}

func CreateDsn() string {
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

var DbConn *sql.DB

func Open() {
	var err error

	//DBに接続(DBのオープン)
	DbConn, err = sql.Open(config.Config.SqlDriver, CreateDsn())
	if err != nil {
		log.Fatalf("Failed to open a database: %v", err)
	}
}
