package routes

import (
	"github.com/jayromofo/project-selkies/server/cmd/api/handlers"
	"github.com/labstack/echo/v4"
)

func GetRecipeRoutes(e *echo.Echo) {
	e.POST("/api/recipe/create/", handlers.AddRecipe2)
	// e.POST("/api/recipe/createSample/", handlers.CreateFullRecipe)
	e.GET("/api/recipe/", handlers.ViewAllRecipes)
	// e.GET("/api/recipe/:id", handlers.ViewRecipe)
	e.GET("/api/recipe/:id", handlers.GetRecipeById)
	e.GET("/api/recipe/:id/instructions", handlers.GetInstructionById)
	e.GET("/api/recipe/:id/metadata", handlers.GetMetadataById)
	e.DELETE("/api/recipe/:id", handlers.DeleteRecipe)
	e.PATCH("/api/recipe/:id", handlers.EditRecipe)
}
