package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"signup/models"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var Coll *mongo.Collection


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
	var claims *models.Claims
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
		// c.JSON(http.StatusOK, gin.H{"message": true})
		
		exprationTime := time.Now().UTC().Add(time.Minute *5)
		claims = &models.Claims{
			Username: loginrequest.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: exprationTime.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte("secret-key"))
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": false})
	}
}
func Add(c *gin.Context) {
	token := c.GetHeader("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	claims := &models.Claims{}
	tkn, err3 := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret-key"), nil
	})
	fmt.Println(err3)
	if !tkn.Valid {
	
		c.JSON(http.StatusBadRequest, gin.H{"message": true})
		return
	}
	var newtodo models.Todo
	err1 := c.ShouldBindJSON(&newtodo)
	fmt.Println(newtodo.Duedate)
	fmt.Println(newtodo.Duetime)
	fmt.Println(newtodo.Name)
	if(newtodo.Duedate==""||newtodo.Duetime==""||newtodo.Name==""){
		c.JSON(http.StatusBadRequest,gin.H{"messages":"no value"})
		return
	}
	fmt.Println(err1)
	filter := bson.D{{"username", claims.Username}}
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
	err2 := Coll.FindOne(c, filter).Decode(&user)
	fmt.Println(user.Todos)
	fmt.Println(err2)
	c.JSON(http.StatusOK, gin.H{"todos": user.Todos})

}
func Deletetodo(c *gin.Context) {
	token := c.GetHeader("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	claims := &models.Claims{}
	tkn, err3 := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret-key"), nil
	})
	fmt.Println(err3)
	if !tkn.Valid {
	
		c.JSON(http.StatusBadRequest, gin.H{"message": true})
		return
	}
	var a int32
	err1 := c.ShouldBindJSON(&a)
	fmt.Println(err1)
	filter := bson.D{{"username", claims.Username}}
	var user models.User
	err := Coll.FindOne(c, filter).Decode(&user)
	if err != nil {
		log.Println("Error finding user:", err)
		return
	}
	user.Todos = append(user.Todos[:a], user.Todos[a+1:]...)
	update := bson.M{
		"$set": bson.M{"todos": user.Todos},
	}
	_, err = Coll.UpdateOne(c, filter, update)
	if err != nil {
		log.Println("Error updating user with new todo:", err)
		return
	}
	err2 := Coll.FindOne(c, filter).Decode(&user)
	fmt.Println(user.Todos)
	fmt.Println(err2)
	c.JSON(http.StatusOK, gin.H{"todos": user.Todos})

}

func GetAllTODO(c *gin.Context) {
	token := c.GetHeader("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	claims := &models.Claims{}
	tkn, err3 := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret-key"), nil
	})
	fmt.Println(err3)
	if !tkn.Valid {
	
		c.JSON(http.StatusBadRequest, gin.H{"message": true})
		return
	}
	var user models.User
	filter := bson.M{"username": claims.Username}
	err := Coll.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "No result in Mongo"})
	}
	c.JSON(http.StatusOK, gin.H{"todos": user.Todos})
}
