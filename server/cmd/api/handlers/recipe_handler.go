package handlers

import (
	"fmt"
	"net/http"

	"github.com/jayromofo/project-selkies/server/cmd/api/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RecipeRepository struct {
	repo service.Repository
}

type Recipe struct {
	gorm.Model
	Name         string              `json:"name" gorm:"primaryIndex`
	Description  string              `json:"description"`
	Type         string              `json:"type"`
	Category     string              `json:"category"`
	Instructions []RecipeInstruction `gorm:"foreignKey:RecipeId"`
	MetaData     RecipeMetaData      `gorm:"foreignKey:RecipeId"`
}

type RecipeCategory struct {
	gorm.Model
	Id          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RecipeMetaData struct {
	gorm.Model
	// Id           uint   `json:"id" gorm:"primaryKey"`
	RecipeId     uint   `json:"recipe_id" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Foreign key referring to Recipe.ID
	Servings     int    `json:"servings"`
	CookTime     int    `json:"cook_time"`
	IsKeto       bool   `json:"is_keto"`
	IsVegetarian bool   `json:"is_vegetarian"`
	Tags         string `json:"tags"`
	ImagePath    string `json:"image_page"`
}

type RecipeInstruction struct {
	gorm.Model
	// Id          uint   `json:"id" gorm:"primaryKey"`
	RecipeId    uint   `json:"recipe_id" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Foreign key referring to Recipe.ID
	LineNum     int    `json:"line_num"`
	Instruction string `json:"instruction"`
}

/* Global */
var recRepo *RecipeRepository
var recipes []Recipe

var recipeSamples = make([]Recipe, 6)

/*
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
*/

func (r *RecipeRepository) InitDb() {
	info := service.InitializeDatabase()
	err := r.repo.Connect(*info)

	recRepo = r
	if err != nil {
		panic("Unable to initialize database")
	}
	dbMigrate(r)
	dbPreload(r)

	// CreateFullRecipe(r)

	// Print out the recipes

	for _, recipe := range recipes {
		println("Recipe:", recipe.Name)
		println("Description:", recipe.Description)
		println("Type:", recipe.Type)
		println("Category:", recipe.Category)
		println("Metadata: Servings:", recipe.MetaData.Servings, "CookTime:", recipe.MetaData.CookTime, "Tags:", recipe.MetaData.Tags)
		for _, instruction := range recipe.Instructions {
			println("Instruction:", instruction.LineNum, instruction.Instruction)
		}
	}

}

func dbMigrate(r *RecipeRepository) {
	r.repo.DB.AutoMigrate(&Recipe{}, &RecipeCategory{}, &RecipeMetaData{}, &RecipeInstruction{})
	fmt.Println("Recipe Repos Migrated")
}

func dbPreload(r *RecipeRepository) {
	r.repo.DB.Preload("Instructions").Preload("MetaData").Find(&recipes)
	fmt.Println("Recipes, Instructions, and MetaData preloaded")
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

	type Payload struct {
		recipePayload      Recipe
		metadataPayload    RecipeMetaData
		instructionPayload RecipeInstruction
	}

	id := c.Param("id")
	payload := Payload{}

	recipeModel := &Recipe{}
	// metadataModel := &RecipeMetaData{}
	// instructionsModel := &RecipeInstruction{}

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

	// err = r.repo.DB.Where("recipe_id = ?", id).First(metadataModel).Error
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, map[string]any{
	// 		"status":  "fail",
	// 		"message": "book unavailable",
	// 	})
	// 	return err
	// }

	/*
		err = r.repo.DB.Where("recipe_id = ?", id).(metadataModel).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]any{
				"status":  "fail",
				"message": "book unavailable",
			})
			return err
		}
	*/

	return c.JSON(http.StatusOK, payload)

}

func ViewAllRecipes(c echo.Context) error {
	// r := SampleData()

	return c.JSON(http.StatusOK, &recipes)
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

func GetInstructionById(c echo.Context) error {
	return nil
}

func GetMetadataById(c echo.Context) error {
	return nil
}

// func CreateFullRecipe(c echo.Context) error {
// 	return recRepo.CreateFullRecipe(c)
// }

// This works to create a recipe transaction
// Now find a way to create it via form
func CreateFullRecipe(r *RecipeRepository) {
	r.repo.DB.Transaction(func(tx *gorm.DB) error {
		newRecipe := Recipe{
			Name:        "Food Sample 1",
			Description: "This is a description of Food Sample 1",
			Type:        "Supper",
			Category:    "Supper",
			MetaData: RecipeMetaData{
				Servings:     4,
				CookTime:     20,
				IsKeto:       false,
				IsVegetarian: false,
				Tags:         "Italian, Pasta",
				ImagePath:    "/images/food-sample-1.jpeg",
			},
			Instructions: []RecipeInstruction{
				{LineNum: 1, Instruction: "Boil water in a large pot."},
				{LineNum: 2, Instruction: "Cook spaghetti until al dente."},
				{LineNum: 3, Instruction: "Fry pancetta in a pan."},
				{LineNum: 4, Instruction: "Mix eggs and cheese in a bowl."},
				{LineNum: 5, Instruction: "Combine spaghetti, pancetta, and egg mixture."},
			},
		}

		if err := tx.Clauses(clause.OnConflict{
			Columns:      []clause.Column{{Name: "name"}},
			DoUpdates:    clause.AssignmentColumns([]string{"description", "type", "category"}),
			OnConstraint: "recipes_pkey",
		}).Create(&newRecipe).Error; err != nil {
			return err
		}
		return nil
	})

}
