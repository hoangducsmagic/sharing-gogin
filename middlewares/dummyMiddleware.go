package middlewares

import "github.com/gin-gonic/gin"

func DummyMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context){
		print("GO THOURGH MIDDLEWARE")
	}
}