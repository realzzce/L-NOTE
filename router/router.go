package router

import (
	"log"
	"other/L-NOTE/middleware"
	"other/L-NOTE/note"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

// InitRouter .
func InitRouter() {
	router := gin.New()
	router.Use(gin.Logger())

	router.Use(middleware.CorsHandler())

	router.Use(gin.Recovery())

	var authMiddleware = middleware.GinJWTMiddlewareInit()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "123ok",
		})
	})
	router.POST("/login", authMiddleware.LoginHandler)
	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("No Route claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	authRouter := router.Group("/auth")
	{
		authRouter.GET("/refresh_token", authMiddleware.RefreshHandler)
	}

	// NOTE note
	apiRouter := router.Group("/api/v1")
	apiRouter.Use(authMiddleware.MiddlewareFunc())
	{
		apiRouter.GET("/note/:noteid", note.GetNote)
		apiRouter.POST("/note", note.AddNote)
		apiRouter.PUT("/note/:noteid", note.UpdateNote)
		apiRouter.DELETE("/note/:noteid", note.DeleteNote)
	}

	router.Run(":8080")
}
