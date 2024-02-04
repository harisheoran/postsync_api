package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/harisheoran/postsync/inits"
	"github.com/harisheoran/postsync/models"
)

/*
Handler Functions
*/

/*
Create Post
*/

// *gin.Context parameter which provides access to information about the HTTP request and allows sending responses.
func CreatePost(context *gin.Context) {
	//  This struct is used to represent the
	// JSON payload that is expected in the HTTP request body
	var expectedPayload struct {
		Title   string
		Content string
	}

	// parse json according to defined payload
	context.BindJSON(&expectedPayload)

	post := models.Post{Title: expectedPayload.Title, Content: expectedPayload.Content}

	result := inits.DB.Create(&post)

	if result.Error != nil {
		context.JSON(
			500,
			gin.H{
				"error": result.Error,
			},
		)
		return
	}

	context.JSON(
		200,
		gin.H{
			"data": post,
		},
	)

}

/*
Get all Posts
*/
func GetPosts(context *gin.Context) {
	var posts []models.Post

	//Find method of the GORM used to retrieve all records from the "posts" table
	// and populate the posts slice with the results.
	result := inits.DB.Find(&posts)

	if result.Error != nil {
		context.JSON(
			500,
			gin.H{
				"error": result.Error,
			},
		)
	}

	context.JSON(
		200,
		gin.H{
			"data": posts,
		},
	)
}

/*
Get a Post
*/
func GetPost(context *gin.Context) {
	var post models.Post

	result := inits.DB.First(&post, context.Param("id"))

	if result.Error != nil {
		context.JSON(
			500,
			gin.H{
				"error": result.Error,
			},
		)
	}

	context.JSON(
		200,
		gin.H{
			"data": post,
		},
	)
}

/*
Update Post
*/
func UpdatePost(context *gin.Context) {
	var expectedPayload struct {
		Title   string
		Content string
	}

	context.BindJSON(&expectedPayload)

	var post models.Post

	result := inits.DB.First(&post, context.Param("id"))
	if result != nil {
		context.JSON(
			500,
			gin.H{
				"error": result.Error,
			},
		)
	}

	inits.DB.Model(&post).Updates(models.Post{Title: expectedPayload.Title, Content: expectedPayload.Content})

	context.JSON(
		200,
		gin.H{
			"data": post,
		},
	)
}

/*
Delete Post
*/
func DeletePost(context *gin.Context) {
	id := context.Param("id")
	inits.DB.Delete(&models.Post{}, id)
	context.JSON(
		200,
		gin.H{
			"data": "post has been deleted successfully",
		},
	)
}
