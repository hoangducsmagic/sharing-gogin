package userControllers

import (
	"context"
	"net/http"
	"sharing-gogin/models"
	"sharing-gogin/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUser(ctx *gin.Context) {
	var user models.User
	
	var db=services.ConnectMongodb();
	coll := db.Database("ginGonicPractice").Collection("user")
	if err := coll.FindOne(context.TODO(), bson.M{"username":ctx.Param("username")}).Decode(&user);err!=nil{
		println(err);
		ctx.JSON(http.StatusBadRequest,"Not found")
		return
	}
	
	ctx.JSON(200,user);
}

