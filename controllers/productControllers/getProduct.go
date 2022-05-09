package productControllers

import (
	"context"
	"net/http"
	"sharing-gogin/models"
	"sharing-gogin/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetProduct(ctx *gin.Context){
	var productId=ctx.Param("id");

	var product models.Product
	var db=services.ConnectMongodb();
	coll:=db.Database("ginGonicPractice").Collection("product");
	if err:=coll.FindOne(context.TODO(),bson.M{"id":productId}).Decode(&product);err!=nil{
		ctx.JSON(http.StatusBadRequest,err.Error());
		return
	}

	ctx.JSON(200,product);
}