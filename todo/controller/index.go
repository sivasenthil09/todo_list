package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"signup/models"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var Coll *mongo.Collection
var Name string

func SetCollection(coll *mongo.Collection) {
	Coll = coll
}
func Signup(c *gin.Context) {
	var signupRequest models.SignupRequest
	if err := c.ShouldBindJSON(&signupRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	name := signupRequest.Username

	filter := bson.D{{"username", name}}
	var result models.SignupRequest
	err := Coll.FindOne(c, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		_, err1 := Coll.InsertOne(c, signupRequest)
		fmt.Println(err1)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": true})
}
func Loginn(c *gin.Context) {
	var loginrequest models.SignupRequest
	if err := c.ShouldBindJSON(&loginrequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	name := loginrequest.Username

	filter := bson.D{{"username", name}}
	var result models.SignupRequest
	err := Coll.FindOne(c, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusBadRequest, gin.H{"message": false})
		return
	}
	if result.Password == loginrequest.Password {
		c.JSON(http.StatusOK, gin.H{"message": true})
		Name = result.Username
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": false})
	}
}
func Add(c *gin.Context) {
	var newtodo models.Todo
	err1 := c.ShouldBindJSON(&newtodo)
	fmt.Println(err1)
	filter := bson.D{{"username", Name}}
	var user models.User
	err := Coll.FindOne(c, filter).Decode(&user)
	if err != nil {
		log.Println("Error finding user:", err)
		return
	}
	
	user.Todos = append(user.Todos, newtodo)
	update := bson.M{
		"$set": bson.M{"todos": user.Todos},
	}
	_, err = Coll.UpdateOne(c, filter, update)
	if err != nil {
		log.Println("Error updating user with new todo:", err)
		return
	}
	err2:= Coll.FindOne(c, filter).Decode(&user)
	fmt.Println(user.Todos)
	fmt.Println(err2)
	c.JSON(http.StatusOK,gin.H{"todos":user.Todos})
     
}
func Deletetodo(c *gin.Context){
	var a int32
	err1:=c.ShouldBindJSON(&a)
	fmt.Println(err1)
	filter := bson.D{{"username", Name}}
	var user models.User
	err := Coll.FindOne(c, filter).Decode(&user)
	if err != nil {
		log.Println("Error finding user:", err)
		return
	}
   user.Todos= append(user.Todos[:a],user.Todos[a+1:]...)
   update := bson.M{
	"$set": bson.M{"todos": user.Todos},
}
_, err = Coll.UpdateOne(c, filter, update)
if err != nil {
	log.Println("Error updating user with new todo:", err)
	return
}
err2:= Coll.FindOne(c, filter).Decode(&user)
fmt.Println(user.Todos)
fmt.Println(err2)
c.JSON(http.StatusOK,gin.H{"todos":user.Todos})

}


func GetAllTODO(c *gin.Context){
	var user models.User
	filter := bson.M{"username":Name}
	err := Coll.FindOne(context.Background(),filter).Decode(&user)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"Error":"No result in Mongo"})
	}
	c.JSON(http.StatusOK,gin.H{"todos":user.Todos})
}
