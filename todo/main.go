package main

import (
	"context"

	"signup/constants"
	"signup/controller"
	"signup/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignupRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type SignupResponse struct {
	Message  string `json:"message"`
	Username string `json:"username"`
}

func main() {
	ctx := context.Background()
	connection := options.Client().ApplyURI(constants.ConnectionString)
	Client, _ := mongo.Connect(ctx, connection)
	Collections := Client.Database("management").Collection("SignUpDetails")
	controller.SetCollection(Collections)

	r := gin.Default()
	routes.AppRoutes(r)

	r.Run(constants.Port)
	
}
