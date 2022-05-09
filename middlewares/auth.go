package middlewares

import (
	"sharing-gogin/models"
	"sharing-gogin/services"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func Auth() gin.HandlerFunc{
	return func(ctx *gin.Context){
		var token=ctx.GetHeader("authToken");
		if (token==""){
			ctx.JSON(400,"Not found token");
			ctx.Abort();
		}

		res,err:=services.ValidateToken(token);
		if (err!=nil){
			ctx.JSON(400,err.Error());
			ctx.Abort();
		}

		if (res.Valid){
			var claim=services.DecodingToken(token);
			var user models.User;
			mapstructure.Decode(claim["User"],&user);
			if (user.Level!="VIP"){
				ctx.JSON(400,"You are too weak");
				ctx.Abort();
			}
		} else {
			ctx.JSON(400,"Token is invalid");
			ctx.Abort();
		}
		
		ctx.Next();
	}
}