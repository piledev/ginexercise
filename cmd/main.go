package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yokoyamada/ginexercise/controller"
)

// Todo is structure
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

func getAll(db *gorm.DB) []Todo {
	todos := []Todo{}
	db.Find(&todos)
	return todos
}

func create(db *gorm.DB, title string, detail string) {
	todo := Todo{}
	todo.Done = false
	todo.Title = title
	todo.Detail = detail
	db.Create(&todo)
}

func main() {

	db := gormConnect()
	defer db.Close()

	// gormの練習（select, insert, update, delete を一通り）
	// gormExercise(db)

	// Engineインスタンスを取得する。
	engin := gin.Default()

	// Glob：パターンマッチング
	// Globパターンで取得したHTMLファイルをHTMLファイルをレンダラーに関連付ける。
	engin.LoadHTMLGlob("../templates/*.tmpl")

	// GET is shortcut for router.Handle("GET",path,handle).
	// Handle registers a new request handle and middleware with the given path and method.
	engin.GET("/", func(c *gin.Context) {
		todos := getAll(db)
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"todos": todos,
		})
	})

	// engin.POST("/new", func(c *gin))
	engin.Run(":8080")
}

func hoge() {

	db := gormConnect()
	defer db.Close()

	// gormの練習（select, insert, update, delete を一通り）
	// gormExercise(db)

	// Engineインスタンスを取得する。
	engin := gin.Default()

	// Glob：パターンマッチング
	// Globパターンで取得したHTMLファイルをHTMLファイルをレンダラーに関連付ける。
	engin.LoadHTMLGlob("../templates/*.html")

	// 静的ファイルの置き場所を指定する。
	// URLで直接指定が可能になる。
	engin.Static("/templates", "../templates")

	// GET is shortcut for router.Handle("GET",path,handle).
	// Handle registers a new request handle and middleware with the given path and method.
	engin.GET("/", controller.IndexGET)
	engin.Run(":8080")
}

func gormExercise(db *gorm.DB) {

	// DB初期化（マイグレーション）→ gooseでやるから不要
	// db.AutoMigrate(&todo)

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
	minid := 0
	todos := []Todo{}
	db.Find(&todos, "title like ?", "Insert%")

	// id の最大値の取得
	for _, r := range todos {
		if maxid < r.ID {
			maxid = r.ID
		}
		if minid == 0 || r.ID < minid {
			minid = r.ID
		}
	}

	// update records
	// こんなことはできそうでできない。
	// todosAfter := []Todo{}
	// for _, r := range todos {
	// 	r.Detail = "update_records"
	// 	todosAfter = append(todosAfter, r)
	// }
	// db.Model(&todos).Updates(&todosAfter)
	// やるならこんな方法
	db.Model(&todo).Where("1=1").Update("detail", "update_all_records")

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
	db.First(&todo, "title like ?", "%2")

	// update a record
	todoAfter = todo
	todoAfter.Detail = "update_like_1"
	db.Model(&todo).Update(&todoAfter)

	// delete a record
	todo = Todo{}
	todo.ID = minid
	db.Delete(&todo)
}
