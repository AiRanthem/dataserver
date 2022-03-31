package utils

import "github.com/gin-gonic/gin"

func Response(status int, data any) gin.H {
	return gin.H{
		"error": status,
		"data":  data,
	}
}
