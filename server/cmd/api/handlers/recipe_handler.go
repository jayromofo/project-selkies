package handlers

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/jayromofo/project-selkies/server/cmd/api/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RecipeRepository struct {
	repo service.Repository
}

type Recipe struct {
	gorm.Model
	Id          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Category    string `json:"category"`
}

type RecipeCategory struct {
	gorm.Model
	Id          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MetaData struct {
	gorm.Model
	Id           int    `json:"id" gorm:"primaryKey"`
	RecipeId     int    `json:"recipe_id" gorm:"foreignKey"`
	Servings     int    `json:"servings"`
	CookTime     int    `json:"cook_time"`
	IsKeto       bool   `json:"is_keto"`
	IsVegetarian bool   `json:"is_vegetarian"`
	Tags         string `json:"tags"`
	ImagePath    string `json:"image_page"`
}

type RecipeInstruction struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	RecipeId    int    `json:"recipe_id" gorm:"foreignKey"`
	LineNum     int    `json:"line_num"`
	Instruction string `json:"instruction"`
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

	recRepo = r
	if err != nil {
		panic("Unable to initialize database")
	}
	dbMigrate(r)

	fmt.Println("Recipe Repos Migrated")

}

func dbMigrate(r *RecipeRepository) {
	r.repo.DB.AutoMigrate(&Recipe{}, &RecipeCategory{}, &MetaData{}, &RecipeInstruction{})
}

// Add recipe via createRecipe
// http://localhost:4444/api/recipe/create
func AddRecipe2(c echo.Context) error {
	var recipe Recipe
	if err := c.Bind(&recipe); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := recRepo.repo.DB.Create(&recipe).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, recipe)

}

func AddRecipe(c echo.Context) error {

	fmt.Println("Adding Recipe: ")

	newRecipe := Recipe{
		Id:          rand.Intn(1000000-1) + 1,
		Name:        c.Param("name"),
		Description: c.Param("description"),
		Type:        c.Param("type"),
		Category:    c.Param("category"),
	}

	fmt.Println("Maybe it worked?")
	fmt.Printf("Name: %s, Category: %s", newRecipe.Name, newRecipe.Category)
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
	if id == "test" {
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
	// Get the id parameter
	id := c.Param("id")

	fmt.Printf("Deleting id: %s \n", id)

	//
	recRepo.repo.DB.Delete(&Recipe{}, id)
	fmt.Print("Deleting Recipe")
	return nil
}

func EditRecipe(c echo.Context) error {
	// Get the id parameter
	id := c.Param("id")

	fmt.Print("Editing Recipe")

	// Find the record data and fill it with the model

	// Overwrite the data and send it back to the database

	if id == "666" {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "Editing Recipe",
			"message": "NomNOm",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "updated",
	})

}
