package main

import (
	productcontrollers "sharing-gogin/controllers/productControllers"
	"sharing-gogin/controllers/userControllers"
	"sharing-gogin/middlewares"
	"sharing-gogin/services"
	"sharing-gogin/validators"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main(){
	router:=gin.Default();
	var address=":3000"

	var db=services.ConnectMongodb();
	defer func(){
		services.DisconnectMongodb(db);
	}()

	router.GET("/ping",func(ctx *gin.Context){
		ctx.String(200,"pong");
	})

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("discountvalidator", validators.DiscountValidator)
	}

	var productRouter=router.Group("/product");
	var userRouter=router.Group("/user")

	productRouter.Use(middlewares.Auth());

	productRouter.POST("/",productcontrollers.CreateProduct);
	productRouter.GET("/",productcontrollers.SearchProduct);
	productRouter.GET("/:id",productcontrollers.GetProduct);

	userRouter.POST("/",userControllers.CreateUser);
	userRouter.GET("/:username",userControllers.GetUser);
	userRouter.POST("/login",userControllers.Login);


	router.Run(address);

	

}