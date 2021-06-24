package config

import (
	"log"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	LogFile string

	DbProtocol    string
	DbAddress     string
	SqlDriver     string
	DbPort        string
	DbName        string
	DbUserName    string
	DbPassword    string
	TestdbName    string
	TestSqlDriver string

	AppPort string
}

var Config *ConfigList

func init() {
	cfg, err := ini.Load("/go/src/github.com/ShunyaNagashige/ca-tech-dojo-golang/config.ini")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	Config = &ConfigList{
		LogFile:       cfg.Section("log").Key("log_file").String(),
		DbProtocol:    cfg.Section("db").Key("db_protocol").String(),
		DbAddress:     cfg.Section("db").Key("db_address").String(),
		SqlDriver:     cfg.Section("db").Key("sql_driver").String(),
		DbPort:        cfg.Section("db").Key("db_port").String(),
		DbName:        cfg.Section("db").Key("db_name").String(),
		DbUserName:    cfg.Section("db").Key("db_user_name").String(),
		DbPassword:    cfg.Section("db").Key("db_password").String(),
		TestdbName:    cfg.Section("db").Key("testdb_name").String(),
		TestSqlDriver: cfg.Section("db").Key("test_sql_driver").String(),
		AppPort:       cfg.Section("app").Key("app_port").String(),
	}
}
