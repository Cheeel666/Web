package server

import (
	"web/config"
	"web/internal/handlers"

	"github.com/gin-gonic/gin"
)

// SetupServer - create server and endpoints
func SetupServer(cfg *config.Config, service handlers.Service) *gin.Engine {
	r := gin.Default()
	r.Use(
		CORSMiddleware(),
		gin.Logger(),
		gin.Recovery(),
	)

	r.GET("/ping", service.Ping)

	api := r.Group("/api/v1")
	{
		user := api.Group("/user")
		{
			user.POST("/add", service.AddUser)
			user.DELETE("/delete/:email", service.DeleteUser)
			user.PUT("/make_mod/:email", service.MakeMod)
			user.GET("/get_all", service.GetUsers)
			user.POST("/login", service.Login)
		}

		courort := api.Group("/courort")
		{
			courort.GET("/courorts", service.GetCourorts)                   //get_courorts
			courort.GET("/roads_and_courorts", service.GetRoadsAndCourorts) //get_roads_and_courorts
			courort.GET("/:cour", service.GetCour)
		}

		comment := api.Group("/comment")
		{
			comment.POST("/add", service.AddComment)                               // add_comment
			comment.DELETE("/delete/:text/:email/:id_cour", service.DeleteComment) //delete_comment
			comment.GET("/get_by_courort/:cour", service.GetComments)              //get_comments
		}

	}

	return r
}

// CORSMiddleware for logging
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
