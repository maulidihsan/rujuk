package database;
import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

type database struct {
	DB *gorm.DB
}

func Init(uri string) (*database, error) {
	db, err := gorm.Open("mysql", uri)
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxIdleConns(10)

	return &database{DB: db}, nil
}

func (d *database) GetDB() *gorm.DB {
	return d.DB
}