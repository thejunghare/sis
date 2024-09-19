package main

import (
	"erp/api/config"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.Connectdb()
)

func main() {
	defer config.Disconnectdb(db)
}
