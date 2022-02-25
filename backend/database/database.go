package database

import (
	"backend/additionalFunc"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)
const (
	dbHost = "192.168.29.174"
	dbPort = "5434"
	dbUser = "darshita"
	dbpass = "darshita"
	dbname = "darshita"
)
 
func SetupDB() *sql.DB {
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbpass, dbname)
	db, err := sql.Open("postgres", dbInfo)
	additionalFunc.CheckErr(err)
	return db
}