package main

import (
	"github.com/gin-gonic/gin"
	"github.com/harisheoran/postsync/controllers"
	"github.com/harisheoran/postsync/inits"
)

func init() {
	inits.LoadEnv()
	inits.DBinit()
}

func main() {
	mainRouter := gin.Default()

	mainRouter.GET("/", controllers.GetPosts)
	mainRouter.GET("/:id", controllers.GetPost)
	mainRouter.POST("/create", controllers.CreatePost)
	mainRouter.PUT("/:id", controllers.UpdatePost)
	mainRouter.DELETE("/:id", controllers.DeletePost)
	mainRouter.Run()

}

func mainHandler(context *gin.Context) {
	context.JSON(
		200,
		gin.H{
			"message": "Post Sync",
		},
	)
}
