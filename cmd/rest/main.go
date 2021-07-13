package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func connectDB() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("環境変数読み込みエラー")
	}
	fmt.Println(os.Getenv("DB_NAME"))
}

func main() {
	connectDB()
	db, err := sql.Open("mysql", "root:rootpassword@tcp(mysql-container:3306)/test-db")

	if err != nil {
		fmt.Println("mysql接続エラー")
	}

	fmt.Println(db)

	fmt.Println("Hello world!")
}