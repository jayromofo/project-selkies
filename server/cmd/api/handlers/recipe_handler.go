package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/jayromofo/project-selkies/server/cmd/api/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RecipeRepository struct {
	repo service.Repository
}

type Recipe struct {
	gorm.Model
	Id          int    `json:'id' gorm:"primaryKey"`
	Name        string `json:'name' gorm:"size:64"`
	Description string `json:'desc' gorm:"size:255"`
	Type        string `json:'type' gorm:"size:64"`
	Category    string `json:'category' gorm:"size:64"`
}

type RecipeCategory struct {
	gorm.Model
	Id          int    `gorm:"primaryKey"`
	Name        string `gorm:"size:255"`
	Description string `gorm:"size:255"`
}

type MetaData struct {
	gorm.Model
	Id           int
	RecipeId     int
	Servings     int
	CookTime     int
	IsKeto       bool
	IsVegetarian bool
	Tags         string
	ImagePath    string
}

type RecipeInstruction struct {
	Id          int
	RecipeId    int
	LineNum     int
	Instruction string
}

/* Global */
var recRepo *RecipeRepository

// var recipeSamples [5]Recipe
var recipeSamples = make([]Recipe, 6)

func SampleData() []Recipe {
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

func (r *RecipeRepository) InitDb() {
	info := service.InitializeDatabase()
	err := r.repo.Connect(*info)
	if err != nil {
		panic("Unable to initialize database")
	}
	dbMigrate(r)

	fmt.Println("Recipe Repos Migrated")
}

func dbMigrate(r *RecipeRepository) {
	r.repo.DB.AutoMigrate(&Recipe{}, &RecipeCategory{}, &MetaData{}, &RecipeInstruction{})
}

func AddRecipe2(c echo.Context) {

}

func AddRecipe(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		fmt.Println("Unable to find an id num")
	}
	if id == 0 {
		id = 1
	}
	fmt.Println("Adding Recipe: ", id)

	newRecipe := Recipe{
		Id:          rand.Intn(1000000-1) + 1,
		Name:        c.QueryParam("name"),
		Description: c.QueryParam("description"),
		Type:        c.QueryParam("type"),
		Category:    c.QueryParam("category"),
	}

	fmt.Println("Maybe it worked?")
	recipeSamples = append(recipeSamples, newRecipe)

	return c.JSON(http.StatusOK, map[string]any{"status": "success", "msg": "Added to recipe list"})

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

func GetRecipeById(c echo.Context) error {
	return recRepo.GetRecipeById(c)
}

func (r *RecipeRepository) GetRecipeById(c echo.Context) error {
	id := c.Param("id")
	recipeModel := &Recipe{}
	if id == "1" {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "fail",
			"message": "id cannot be empty",
		})
		return nil
	}
	fmt.Println("the ID is", id)

	err := r.repo.DB.Where("id = ?", id).First(recipeModel).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "fail",
			"message": "book unavailable",
		})
		return err
	}
	return c.JSON(http.StatusOK, recipeModel)

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
