package pzn_golang_db

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
)

func TestExecSql(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customers(name) VALUES('fajar setiawan siagian')"
	_, err := db.ExecContext(ctx, query)
	if err != nil{
		panic(err.Error())
	}

	fmt.Println("success insert new customer")
}

func TestQuerySql(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT * FROM customers"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	for rows.Next(){
		var id int
		var name string

		err := rows.Scan(&id, &name)
		if err != nil{
			panic(err)
		}

		fmt.Println("ID: ", id)
		fmt.Println("Name: ", name)
	}

	fmt.Println("success get customers")
}

func TestQuerySqlComplex(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id,name,email,balance,rating,created_at,birth_date,married FROM customers"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	for rows.Next(){
		// var id int
		// var name, email string
		// var balance int
		// var rating float64
		// var createdAt, birthDate time.Time
		// var married bool

		// ini agar bisa null tapi jadinya struct
		var id int
		var name, email sql.NullString
		var balance sql.NullInt64
		var rating sql.NullFloat64
		var createdAt, birthDate sql.NullTime
		var married bool
	
		err := rows.Scan(&id, &name, &email, &balance, &rating, &createdAt, &birthDate, &married)
		if err != nil{
			panic(err)
		}

		fmt.Println("=================")
		fmt.Println("ID: ", id)
		fmt.Println("email: ", email)
		fmt.Println("balance: ", balance)
		fmt.Println("rating: ", rating)
		fmt.Println("created_at: ", createdAt)
		fmt.Println("birth_date: ", birthDate)
		fmt.Println("married: ", married)
	}

	// bagaimana melakukan null ke string, jika ada field yang null

	fmt.Println("success get customers")
}

func TestSqlInjection(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"

	query := "SELECT id, username FROM users WHERE username = '"+username+"' LIMIT 1"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	if rows.Next(){
		var id int
		var username string
		err := rows.Scan(&id, &username)
		if err != nil{
			panic(err)
		}
		fmt.Println("success login: ",username)
		}else{
			fmt.Println("gagal login")
		}
}

func TestSqlInjectionSafe(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"

	query := "SELECT id, username FROM users WHERE username = ? LIMIT 1"
	fmt.Println(query)
	rows, err := db.QueryContext(ctx, query, username)
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	if rows.Next(){
		var id int
		var username string
		err := rows.Scan(&id, &username)
		if err != nil{
			panic(err)
		}
		fmt.Println("success login: ",username)
		}else{
			fmt.Println("gagal login")
		}
}

func TestExecSqlParameter(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "fajarsiagian1; DROP TABLE users; #"
	pass := "password"

	query := "INSERT INTO users(username, password) VALUES(?,?)"
	_, err := db.ExecContext(ctx, query, username, pass)
	if err != nil{
		panic(err.Error())
	}

	fmt.Println("success insert new customer")
}

func TestGetIdFromInsert(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "fajarsiagian@mail.test"
	comment := "comment terbaru"

	query := "INSERT INTO comments(email, comment) VALUES(?,?)"
	result, err := db.ExecContext(ctx, query, email, comment)
	if err != nil{
		panic(err.Error())
	}

	insertId, err := result.LastInsertId()
	if err != nil{
		panic(err)
	}

	fmt.Println("success insert new comment with id ", insertId)
}

func TestPrepareStatement(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "INSERT INTO comments(email, comment) VALUES(?,?)"
	statement, err := db.PrepareContext(ctx, query)
	if err != nil{
		panic(err)
	}
	defer statement.Close()

	for i:=0; i<10; i++{
		email := "fajar" + strconv.Itoa(i) + "@mail.local"
		comment := "sudah saya bilang ini komen ke" + strconv.Itoa(i)
		
		// tidak perlu buat script query karena sudah di binding di atas
		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil{
			panic(err)
		}

		id, err := result.LastInsertId()
		if  err != nil{
			panic(err)
		}

		fmt.Println("komen id ", id)
	}

}

func TestTransactionDB(t *testing.T){
	db := GetConnection()
	defer db.Close()

	// transactin db

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil{
		panic(err)
	}

	// do transaction
	query := "INSERT INTO comments(email, comment)VALUES(?,?)"
	for i:=0; i<10; i++{
		email := "ariaja" + strconv.Itoa(i) + "@mail.test"
		comment := "aku komen ini sudah sebanyak " + strconv.Itoa(i) + " kali"

		result, err := tx.ExecContext(ctx, query, email, comment)
		if err != nil{
			panic(err)
		}

		id,err := result.LastInsertId()
		if err != nil{
			panic(err)
		}
		
		fmt.Println("Comment id ", id)
	}


	// err = tx.Commit()
	// if err != nil {
	// 	panic(err)
	// }

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}