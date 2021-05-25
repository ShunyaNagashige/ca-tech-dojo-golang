package config

import (
	"log"

	"gopkg.in/ini.v1"
)

type ConfigList struct{
	LogFile string
	DbUserName string
	DbPassword string
	DbName string
	SqlDriver string
	DbPort int
	AppPort int
}

var Config *ConfigList

func init(){
	cfg,err:=ini.Load("config.ini")
	if err!=nil{
		log.Fatalf("Failed to read file: %v",err)
	}

	Config=&ConfigList{
		LogFile: cfg.Section("log").Key("log_file").String(),
		DbUserName: cfg.Section("db").Key("user_name").String(),
		DbPassword: cfg.Section("db").Key("db_password").String(),
		DbName: cfg.Section("db").Key("db_name").String(),
		SqlDriver: cfg.Section("db").Key("sql_driver").String(),
		DbPort: cfg.Section("db").Key("db_port").MustInt(),
		AppPort: cfg.Section("app").Key("app_port").MustInt(),
	}
}