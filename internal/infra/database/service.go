package database

import (
	"native-setup/config"

	"gorm.io/gorm"
)


func NewService(env config.Env, logger *DBLogger) (*gorm.DB, error) {
	db, err := GetStandalone(env, logger)
	if err != nil {
		return nil, err
	}
	return db, nil
}
