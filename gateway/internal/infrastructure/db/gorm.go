package database

import (
	"github.com/superdev/ecommerce/gateway/internal/config"
	gormrepo "github.com/superdev/ecommerce/gateway/internal/data/repository/gorm_repo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(conf *config.Config) (*gorm.DB, error) {

	dsn := conf.DB.Dsn()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Auto-migrate User model
	if conf.AutoMigrate {
		//TODO: Create a migrator service with injectors
		if err = db.AutoMigrate(&gormrepo.GormUser{}); err != nil {
			return nil, err
		}
	}

	return db, err
}
