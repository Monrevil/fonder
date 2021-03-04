package main

import (
	"fmt"

	"gorm.io/gorm"
)

func pingDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	if err = sqlDB.Ping(); err != nil {
		return err
	}
	return nil
}

func gormStat(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	stats := sqlDB.Stats()
	fmt.Printf("[DB]:conn: \nin use: %v\nopen: %v\nidle: %v\n", stats.InUse, stats.OpenConnections, stats.Idle)
	fmt.Println("Max open conn:", stats.MaxOpenConnections)
	return nil
}

// HTTPError error struct
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}