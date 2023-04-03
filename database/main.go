package database

import (
	"tugas_golang_alfi_ahsani/database/migrations"
	"tugas_golang_alfi_ahsani/database/server"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// "gorm.io/driver/postgres"

var DB *gorm.DB

func Connect() {
	// if os.Getenv("DB_TYPE") == "mysql" {
	dsn := server.DsnMysql()
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed enpoint connect mysql database")
	}
	DB = connection
	err = migrations.Migration(connection)
	if err != nil {
		return
	}
	// }

	// if os.Getenv("DB_TYPE") == "postgres" {
	// 	dsn := server.DsnPostgres()
	// 	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// 	if err != nil {
	// 		panic("failed enpoint connect postgres database")
	// 	}
	// 	DB = connection
	// 	err = migrations.Migration(connection)
	// 	if err != nil {
	// 		return
	// 	}
	// }
}
