package post

import "github.com/gin-gonic/gin"

func Setup(router *gin.RouterGroup) {
	post := router.Group("/post")
	{
		post.GET("/", getPosts)
		post.POST("/", createPost)
		post.GET("/:id", getPostByID)
		post.PUT("/:id", updatePost)
		post.DELETE("/:id", deletePost)
	}
}