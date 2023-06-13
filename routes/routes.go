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
	private.Use(middleware.RequireJWT())

	private.GET("/users", middleware.CheckRole(1), controllers.GetAllUsers)
	private.GET("/users/:id", middleware.CheckRole(2), controllers.GetUserByID)
	private.PUT("/users/:id", middleware.CheckRole(1), controllers.UpdateUser)

	private.GET("/books", middleware.CheckRole(2), controllers.GetAllBooks)
	private.GET("/books/:id", middleware.CheckRole(2), controllers.GetBookByID)
	private.POST("/books", middleware.CheckRole(1), controllers.AddBook)
	private.DELETE("/books/:id", middleware.CheckRole(1), controllers.DeleteBook)
	private.PUT("/books/:id", middleware.CheckRole(1), controllers.UpdateBook)

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
