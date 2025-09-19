package connection

import (
	"crudwebsocket/internal/config"
	"crudwebsocket/model"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func GetDatabase(conf config.DataBase) (*Database, error) {

	db, err := gorm.Open(postgres.Open(conf.Database), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.Cars{})
	return &Database{db: db}, nil

}

func (d *Database) Close() error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

func (d *Database) GetDB() *gorm.DB {
	return d.db
}