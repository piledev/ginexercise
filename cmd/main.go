package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yokoyamada/ginexercise/controller"
)

func main() {
	// Engineインスタンスを取得する。
	router := gin.Default()

	// Glob：パターンマッチング
	// Globパターンで取得したHTMLファイルをHTMLファイルをレンダラーに関連付ける。
	router.LoadHTMLGlob("../templates/*.html")

	// GET is shortcut for router.Handle("GET",path,handle).
	// Handle registers a new request handle and middleware with the given path and method.
	router.GET("/", controller.IndexGET)
	router.Run(":8080")
}
