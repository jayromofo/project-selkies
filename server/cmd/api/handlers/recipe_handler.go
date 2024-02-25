package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Recipe struct {
	Id          int
	Name        string
	Description string
	Type        string
	Category    string
}

type RecipeCategory struct {
}

func AddRecipe(c echo.Context) error {

	return nil
}

func ViewRecipe(c echo.Context) error {
	r := Recipe{
		Id:          1,
		Name:        "Waffles",
		Description: "Beautiful Savory Waffles",
		Type:        "Breakfast",
		Category:    "Breakfast",
	}

	// id := c.Param("id")

	return c.JSON(http.StatusOK, r)
}

func ViewAllRecipes(c echo.Context) error {

	return nil
}

func DeleteRecipe(c echo.Context) error {
	return nil
}

func EditRecipe(c echo.Context) error {
	return nil
}
