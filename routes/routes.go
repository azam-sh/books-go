package routes

import (
	"log"
	"net/http"

	"github.com/azam-sh/books/controllers"
	"github.com/azam-sh/books/middleware"
	"github.com/gin-gonic/gin"
)

func StartRoutes() {
	r := gin.Default()

	public := r.Group("/v1")

	public.GET("/ping", ping)
	public.POST("/signup", controllers.Signup)
	public.POST("/login", controllers.Login)

	private := r.Group("/v2")
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
