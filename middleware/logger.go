package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("start")
		c.Next()
		fmt.Println("end")
	}
}
