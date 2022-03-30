package dao

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type dbAccess struct {
	db *gorm.DB
}

func (d *dbAccess) DB() *gorm.DB {
	return d.db
}

func SQLite(dsn string, target any) *dbAccess {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.WithError(err).Fatalln("failed to connect database")
	}
	err = db.AutoMigrate(target)
	if err != nil {
		log.WithError(err).Fatalln("failed to migrate schema")
	}
	return &dbAccess{
		db: db,
	}
}

func MySQL(user string, pass string, ip string, port int, dbName string, target any) *dbAccess {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, ip, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.WithError(err).Fatalln("failed to connect database")
	}
	err = db.AutoMigrate(target)
	if err != nil {
		log.WithError(err).Fatalln("failed to migrate schema")
	}
	return &dbAccess{
		db: db,
	}
}
