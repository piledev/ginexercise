package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IndexGET displays application index page
func IndexGET(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"message": "Hello go/gin/gorm!",
	})
}
