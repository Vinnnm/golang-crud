package main

import (
	"github.com/Vinnnm/golang-crud/controllers"
	"github.com/Vinnnm/golang-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

// Reference Page : https://gin-gonic.com/docs/introduction/
func main() {

	r := gin.Default()

	// CREATE
	r.POST("/posts", controllers.PostsCreate)

	// READ ALL
	r.GET("/posts", controllers.PostsIndex)

	// READ BY ID
	r.GET("/posts/:id", controllers.PostsShow)

	// UPDATE
	r.PUT("/posts/:id", controllers.PostsUpdate)

	//DELETE
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run()
}
