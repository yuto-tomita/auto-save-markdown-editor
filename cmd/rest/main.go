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
	
	db, err := sql.Open(
		os.Getenv("DB_CONNECTION"), os.Getenv("DB_NAME") + ":" + os.Getenv("DB_PASS") + "@tcp(mysql-container:3306)/" + os.Getenv("DB_NAME"))
	if (err != nil) {
		fmt.Println("DB接続エラー")
	}

	fmt.Println(db)
}

func main() {
	connectDB()
}