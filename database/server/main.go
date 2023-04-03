package server

import (
	"log"
	"os"
)

const SSLMode = "disable"

func DsnMysql() string {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	log.Print("dbname", dbname)
	port := os.Getenv("DB_PORT_MYSQL")
	return user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname
}

func DsnPostgres() string {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT_POSTGRES")
	sslmode := SSLMode
	timezone := "Asia/Jakarta"
	return "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=" + sslmode + " TimeZone=" + timezone
}
