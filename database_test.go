package pzn_golang_db

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T){
	// db, err := sql.Open("mysql", "fajarsiagian:tekanetner8x@tcp(localhost:3306)/gdc_local_payment_gateway")
	// if err != nil{
	// 	panic(err)
	// }
	// defer db.Close()

	db := GetConnection()
	defer db.Close()

	// gunakan DB
}