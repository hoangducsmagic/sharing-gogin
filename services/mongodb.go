package services

import (
	"context"
	"sharing-gogin/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongodb() *mongo.Client{
	var uri=utils.GetEnvVariable("CONNECTION_STRING");
	db,err:=mongo.Connect(context.TODO(),options.Client().ApplyURI(uri));
	if (err!=nil){
		panic(err)
	}

	return db;
}

func DisconnectMongodb(db *mongo.Client){
	if err:=db.Disconnect(context.TODO());err !=nil{
		panic(err);
	}
}