package models

type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Level    string `json:"level" bson:"level"`
}