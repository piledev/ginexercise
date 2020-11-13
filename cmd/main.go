package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yokoyamada/ginexercise/controller"
)

func main() {
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
