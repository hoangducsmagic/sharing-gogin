package userControllers

import (
	"context"
	"log"
	"sharing-gogin/models"
	"sharing-gogin/services"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var user models.User
	if err:=ctx.ShouldBindJSON(&user);err!=nil{
		log.Fatal(err);
	}

	var db=services.ConnectMongodb();
	coll := db.Database("ginGonicPractice").Collection("user")

	_, err := coll.InsertOne(context.TODO(), user)
	if (err!=nil){
		log.Fatal(err);
	}	
	ctx.JSON(200,user);
}