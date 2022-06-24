package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//r.LoadHTMLGlob("templates/*")
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin")
	})

	r.Run(":8000")

}
