package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Todo is construction
type Todo struct {
	ID     int `gorm:"primary_key"`
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

func gormExercise(db *gorm.DB) {

	// delete all records
	todo := Todo{}
	db.Delete(&todo)

	// insert 4 records
	for i := 0; i < 4; i++ {
		todofor := Todo{}
		todofor.Title = fmt.Sprintf("insert_%v", i)
		db.Create(&todofor)
	}

	// select records
	maxid := 0
	todos := []Todo{}
	db.Find(&todos, "title like ?", "Insert%")

	// update records
	todosAfter := todos
	for _, r := range &todosAfter {
		r.Detail = "update_records"
		if maxid < r.ID {
			maxid = r.ID
		}
		fmt.Println(r.ID, r.Detail)
	}
	db.Model(&todos).Update(&todosAfter)

	// select a record pattern 1
	todo = Todo{}
	todo.ID = maxid
	db.First(&todo)

	// update a record
	todoAfter := todo
	todoAfter.Detail = "update_maxid"
	db.Model(&todo).Update(&todoAfter)

	// select a record pattern 2
	todo = Todo{}
	db.First(&todo, "title like ?", "%1")

	// update a record
	todoAfter = todo
	todoAfter.Detail = "update_like_1"
	db.Model(&todo).Update(&todoAfter)

	// delete a record
	// todo = Todo{}

}

func main() {

	db := gormConnect()
	defer db.Close()

	// DB初期化（マイグレーション）→ gooseでやるから不要
	// db.AutoMigrate(&todo)

	gormExercise(db)

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
