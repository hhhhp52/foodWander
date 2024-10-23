package routes

import (
	"fmt"
	"foodWander/src/handlers"
	"foodWander/src/middlleware"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	_ "strings"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func Router() *gin.Engine {
	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.New()

	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	router.Use(Logger())
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.Use(gin.Recovery())

	api := router.Group("/api")

	api.POST("/register", handlers.Register)
	api.POST("/login", handlers.Login)
	api.POST("/verify-email", handlers.VerifyEmail)
	api.POST("/logout", handlers.Logout)

	restaurant := api.Group("/restaurant")
	restaurant.GET("/all", handlers.GetRestaurants)
	restaurant.GET("/detail", handlers.GetRestaurantDetail)

	restaurant.Use(middlleware.AuthRequired)
	{
		restaurant.POST("/create", handlers.CreateRestaurant)
		restaurant.PUT("/update", handlers.UpdateRestaurant)
		restaurant.DELETE("/delete", handlers.DeleteRestaurant)
	}

	comment := api.Group("/comment")
	comment.GET("/all", handlers.GetComments)

	comment.Use(middlleware.AuthRequired)
	{
		comment.POST("/create", handlers.CreateComment)
		comment.PUT("/update", handlers.UpdateComment)
		comment.DELETE("/delete", handlers.DeleteComment)
	}

	user := api.Group("/user")
	user.Use(middlleware.AuthRequired)
	{
		user.GET("/me", handlers.GetUser)
		user.GET("/profile/:user_id", handlers.GetProfile)
		user.PUT("/update_profile", handlers.UpdateProfile)
		user.POST("/update_image", handlers.UpdateProfileImage)
		user.DELETE("/delete", handlers.DeleteUser)
	}

	admin := api.Group("/admin", middlleware.AuthRequired)
	admin.GET("/users", handlers.GetUser)

	return router
}
