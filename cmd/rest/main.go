package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:rootpassword@tcp(mysql-container:3306)/test-db")

	if err != nil {
		fmt.Println("mysql接続エラー")
	}

	fmt.Println(db)

	fmt.Println("Hello world!")
}