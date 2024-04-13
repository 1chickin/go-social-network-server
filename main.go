package main

import (
	"log"

	"github.com/1chickin/go-social-network-server/config"
	"github.com/1chickin/go-social-network-server/internal/controller"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
}

func main() {
	db := config.ConnectDatabase()
	if db == nil {
		log.Fatal("Failed to connect to database")
	}
	config.MigrateDB()
	redisClient := config.ConnectRedis()
	router := gin.Default()
	router.POST("/signup", controller.Signup)
	router.POST("/login", func(c *gin.Context) {
		controller.Login(c, redisClient)
	})
	/*
		router.POST("/v1/sessions", controllers.Login)
		router.POST("/v1/users", controllers.SignUp)
		router.PUT("/v1/users", controllers.EditProfile)
		router.GET("/v1/friends/:user_id", controllers.SeeFollowList)
		router.POST("/v1/friends/:user_id", controllers.Follow)
		router.DELETE("/v1/friends/:user_id", controllers.Unfollow)
		router.GET("/v1/friends/:user_id/posts", controllers.SeeUserPosts)
		router.GET("/v1/posts/:post_id", controllers.SeePost)
		router.POST("/v1/posts", controllers.CreatePost)
		router.PUT("/v1/posts/:post_id", controllers.EditPost)
		router.DELETE("/v1/posts/:post_id", controllers.DeletePost)
		router.POST("/v1/posts/:post_id/comments", controllers.CommentPost)
		router.POST("/v1/posts/:post_id/likes", controllers.LikePost)
		router.GET("/v1/newsfeeds", controllers.Newsfeed)
	*/
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
