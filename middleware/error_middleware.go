package middleware

import (
	"go-crud-demo/model"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			if appErr, ok := err.(*model.AppError); ok {
				c.JSON(appErr.Code, gin.H{
					"error": appErr.Message,
				})
				return
			}

			c.JSON(500, gin.H{
				"error": "internal server error",
			})
		}
	}
}
