package productControllers

import (
	"context"
	"net/http"
	"sharing-gogin/models"
	"sharing-gogin/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateProduct(ctx *gin.Context) {
	var product models.Product
	if err:=ctx.ShouldBindJSON(&product);err!=nil{
		ctx.JSON(http.StatusBadRequest,err.Error());
		return
	}

	product.Id=uuid.New().String();

	var db=services.ConnectMongodb();
	coll:=db.Database("ginGonicPractice").Collection("product");

	_,err:=coll.InsertOne(context.TODO(),product);
	if (err!=nil){
		ctx.JSON(http.StatusBadRequest,err.Error());
		return
	}
	ctx.JSON(200,product);
}