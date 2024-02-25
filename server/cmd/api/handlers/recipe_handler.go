package handlers

import (
	"fmt"
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

// var recipeSamples [5]Recipe

func SampleData() []Recipe {
	var recipeSamples = make([]Recipe, 5)
	recipeSamples[0] = Recipe{
		Id:          1,
		Name:        "Waffles",
		Description: "Beautiful Savory Waffles",
		Type:        "Breakfast",
		Category:    "Breakfast",
	}

	recipeSamples[1] = Recipe{
		Id:          2,
		Name:        "Steak",
		Description: "Beautiful Medium Rare Steak",
		Type:        "Supper",
		Category:    "Meat",
	}

	recipeSamples[2] = Recipe{
		Id:          3,
		Name:        "Chicken Noodle Soup",
		Description: "Homemade Chicken Noodle Soup",
		Type:        "Supper",
		Category:    "Soup",
	}

	recipeSamples[3] = Recipe{
		Id:          4,
		Name:        "Jalepeno Poppers",
		Description: "Spicy Jalepeno Poppers",
		Type:        "Snack",
		Category:    "Snack",
	}

	recipeSamples[4] = Recipe{
		Id:          5,
		Name:        "Smoked Salmon",
		Description: "Beautiful Smoked Salmon",
		Type:        "Snack",
		Category:    "Snack",
	}
	return recipeSamples

}

func AddRecipe(c echo.Context) error {
	response := c.Param("id")
	fmt.Println("Maybe it worked?")
	fmt.Printf("%s", response)
	return c.JSON(http.StatusOK, response)

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
	r := SampleData()

	return c.JSON(http.StatusOK, r)
}

func DeleteRecipe(c echo.Context) error {
	return nil
}

func EditRecipe(c echo.Context) error {
	return nil
}
