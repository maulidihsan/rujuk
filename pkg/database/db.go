package database;
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

type database struct {
	db *gorm.DB
}

func Init(uri string) (*database, error) {
	db, err := gorm.Open("mysql", uri)
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxIdleConns(10)
	defer db.Close()

	return &database{db: db}, nil
}

func (d *database) GetDB() *gorm.DB {
	return d.db
}