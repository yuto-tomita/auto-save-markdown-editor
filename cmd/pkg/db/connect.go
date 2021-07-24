package db

import (
	"fmt"
	"os"

	domain "myapp/pkg/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func ConnectDB() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("環境変数読み込みエラー")
	}

	connectTemplate := "%s:%s@%s/%s"

	connect := fmt.Sprintf(connectTemplate, os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB_NAME"))

	db, err := gorm.Open(os.Getenv("DB_CONNECTION"), connect)
	
	if (err != nil) {
		fmt.Println("DB接続エラー")
	}
	
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&domain.Draft{})
	// db.AutoMigrate(&domain.Draft{})

	return db
}
