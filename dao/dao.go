package dao

import (
	"gin-course/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	Db, err = gorm.Open(mysql.Open(config.MysqlDSN), &gorm.Config{})

	if err != nil {
		log.Panic("mysql connect error: ", err)
	}
	if Db.Error != nil {
		log.Panic("database error: ", Db.Error)
	}
}
