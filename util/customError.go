package util

import "github.com/gin-gonic/gin"

func CustomError(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
