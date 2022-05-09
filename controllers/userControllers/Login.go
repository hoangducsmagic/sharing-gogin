package userControllers

import (
	"context"
	"log"
	"net/http"
	"sharing-gogin/models"
	"sharing-gogin/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Login(ctx *gin.Context) {
	var reqUser models.User
	var dbUser models.User
	
	if err:=ctx.ShouldBindJSON(&reqUser);err!=nil{
		log.Fatal(err);
	}

	var db=services.ConnectMongodb();
	coll := db.Database("ginGonicPractice").Collection("user")
	if err := coll.FindOne(context.TODO(), bson.M{"username":reqUser.Username}).Decode(&dbUser);err!=nil{
		println(err);
		ctx.JSON(http.StatusBadRequest,"User does not existed!")
		return
	}

	if (dbUser.Password!=reqUser.Password){
		ctx.JSON(http.StatusBadRequest,"Wrong password!")
		return
	}

	var token=services.GenerateToken(dbUser);
	
	ctx.JSON(200,gin.H{
		"authToken":token,
	});
}

