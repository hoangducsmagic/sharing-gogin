package productcontrollers

import (
	"context"
	"net/http"
	"sharing-gogin/models"
	"sharing-gogin/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func SearchProduct(ctx *gin.Context){	//keyword
	var keyword=ctx.DefaultQuery("keyword","");

	var db=services.ConnectMongodb();
	coll:=db.Database("ginGonicPractice").Collection("product");
	filter:=bson.M{"name":bson.M{"$regex":keyword,"$options":"ig"}}

	cursor,err:=coll.Find(context.TODO(),filter);
	if (err!=nil){
		ctx.JSON(http.StatusBadRequest,err.Error());
		return
	}

	var products []models.Product
	if err:=cursor.All(ctx,&products);err!=nil{
		ctx.JSON(http.StatusBadRequest,err.Error());
		return
	}

	ctx.JSON(200,products);
}