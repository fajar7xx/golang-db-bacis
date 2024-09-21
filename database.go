package pzn_golang_db

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB{
	// tambahkan parseTime=true agar bisa membaca timestamp
	db, err := sql.Open("mysql", "fajarsiagian:tekanenter8x@tcp(localhost:3306)/pzn_golang_db?parseTime=true")
	if err != nil{
		panic(err.Error())
	}
	// defer db.Close()

	// open doesnt open a connection, validate dsn data:
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	// set connection pooling
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}