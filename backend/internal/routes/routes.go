package routes

import (
	"cook_book/backend/internal/controllers/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	h := handlers.InitAllHandlers()

	r.GET("/users", handlers.GetAllUsers)
	r.POST("/login", handlers.Login)
	r.POST("/register", handlers.Register)

	cook := r.Group("/")
	// cook.Use(middleware.JWTMiddleware())

	cook.GET("/", h.GetAllRecipes)
	cook.GET("//:id", h.GetRecipeByID)
	cook.POST("/", h.CreateRecipe)
	cook.PUT("//:id", h.UpdateRecipe)
	cook.DELETE("//:id", h.DeleteRecipe)

	return r
}
