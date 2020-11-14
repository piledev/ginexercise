package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Todo is construction
type Todo struct {
	ID     int
	Done   bool
	Title  string
	Detail string
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "password"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "tododb"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {

	db := gormConnect()
	defer db.Close()

	// DB初期化（マイグレーション）→ gooseでやるから不要
	// db.AutoMigrate(&todo)

	// 構造体のインスタンス化
	todo := Todo{}

	// Insert
	todo.ID = 1
	todo.Title = "test"
	db.Create(&todo)

	// Select Update

	// Select multi

	// Delete

	// // Engineインスタンスを取得する。
	// engin := gin.Default()

	// // Glob：パターンマッチング
	// // Globパターンで取得したHTMLファイルをHTMLファイルをレンダラーに関連付ける。
	// engin.LoadHTMLGlob("../templates/*.html")

	// // 静的ファイルの置き場所を指定する。
	// // URLで直接指定が可能になる。
	// engin.Static("/templates", "../templates")

	// // GET is shortcut for router.Handle("GET",path,handle).
	// // Handle registers a new request handle and middleware with the given path and method.
	// engin.GET("/", controller.IndexGET)
	// engin.Run(":8080")
}
