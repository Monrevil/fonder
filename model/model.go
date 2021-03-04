package model

import (
	"fmt"
	"mnt/c/Users/DELL/nix/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


//Google is a Helper struct
type Google struct {
	Code string `path:"code" query:"code" form:"code" json:"code"`
}

//PingDB establishes conn to the DB if none are available
func PingDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	if err = sqlDB.Ping(); err != nil {
		return err
	}
	return nil
}

//GormStat prints out db stats
//TODO: use Prometheus + Grafana instead
func GormStat(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	stats := sqlDB.Stats()
	fmt.Printf("[DB]:conn: \nin use: %v\nopen: %v\nidle: %v\n", stats.InUse, stats.OpenConnections, stats.Idle)
	fmt.Println("Max open conn:", stats.MaxOpenConnections)
	return nil
}

//InitDB Updates db schemas
func InitDB() *gorm.DB {
	dsn := config.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Post{}, &Comment{}, &User{})

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(100)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Minute * 10)

	return db
}
