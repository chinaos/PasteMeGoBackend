/*
@File: model.go
@Contact: lucien@lucien.ink
@Licence: (C)Copyright 2019 Lucien Shui

@Modify Time      @Author    @Version    @Description
------------      -------    --------    -----------
2019-06-23 14:03  Lucien     1.0         None
*/
package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/wonderivan/logger"
	"os"
)

var (
	username = "username"
	password = "password"
	network  = "tcp"
	server   = "mysql"
	port     = 3306
	database = "pasteme"
)

func format(
	username string,
	password string,
	network string,
	server string,
	port int,
	database string) string {
	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s?parseTime=True&loc=Local", username, password, network, server, port, database)
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", format(username, password, network, server, port, database))
	if err != nil {
		logger.Fatal("Connect to MySQL failed: " + err.Error())
	} else {
		logger.Info("MySQL connected")
		if os.Getenv("PASTEMED_RUNTIME") == "debug" {
			logger.Warn("Running in debug mode, database execute will be displayed")
			db = db.Debug()
		}
		if !db.HasTable(&Permanent{}) {
			logger.Warn("Table permanents not found, start creating")
			if err := db.Set(
				"gorm:table_options",
				"ENGINE=Innodb DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=100",
			).CreateTable(&Permanent{}).Error; err != nil {
				logger.Fatal("Create table permanents failed: " + err.Error())
			}
		}

		if !db.HasTable(&Temporary{}) {
			logger.Warn("Table temporaries not found, start creating")
			if err := db.Set(
				"gorm:table_options",
				"ENGINE=Innodb DEFAULT CHARSET=utf8mb4",
			).CreateTable(&Temporary{}).Error; err != nil {
				logger.Fatal("Create table temporaries failed: " + err.Error())
			}
		}
	}
}

func FixAutoIncrement(autoIncrement uint64) error {
	return db.Exec(fmt.Sprintf("ALTER TABLE `permanents` AUTO_INCREMENT=%d", autoIncrement)).Error
}
