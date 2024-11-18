package connection

import (
	"github.com/andremartinsds/go_admin/configs"
	"github.com/andremartinsds/go_admin/internal/infra/db/automigrate"
	"gorm.io/gorm"
)

var DataSource *gorm.DB

func DatabaseStart() (*gorm.DB, error) {
	db, err := configs.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	err = automigrate.AutoMigrate(db)
	if err != nil {
		return nil, err
	}
	DataSource = db
	return db, nil
}
