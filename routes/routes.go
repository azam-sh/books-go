package routes

import (
	"log"
	"net/http"

	"books/controllers"
	"books/middleware"

	"github.com/gin-gonic/gin"
)

func StartRoutes() {
	r := gin.Default()

	r.GET("/ping", ping)

	private := r.Group("/v1")
	private.Use(middleware.JWTAuthMiddleware())

	private.GET("/users", controllers.GetAllUsers)
	private.GET("/users/:id", controllers.GetUserByID)
	private.PUT("/users/:id", controllers.UpdateUser)

	private.GET("/books", controllers.GetAllBooks)
	private.GET("/books/:id", controllers.GetBookByID)
	private.POST("/books", controllers.AddBook)
	private.DELETE("/books/:id", controllers.DeleteBook)
	private.PUT("/books/:id", controllers.UpdateBook)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "page not found"})
	})

	err := r.Run()
	if err != nil {
		log.Panic("failed to start router")
	}
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, "Connection established!")
}
