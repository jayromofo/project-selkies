package main

import (
	"fmt"
	"net/http"

	"github.com/jayromofo/project-selkies/server/cmd/api/routes"
	"github.com/jayromofo/project-selkies/server/cmd/api/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Person struct {
	Id         int
	Name       string
	Age        int
	Email      string
	Profession string
}

func main() {
	fmt.Println("Testing gogo")

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAuthorization, echo.HeaderXRequestedWith},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if apiErr, ok := err.(APIError); ok {
			c.JSON(apiErr.Status, map[string]any{"error": apiErr.Msg})
			return
		}
		c.JSON(http.StatusInternalServerError, map[string]any{"error": "internal server error"})
	}

	routes.GetRecipeRoutes(e)
	routes.GetBudgetRoutes(e)
	r := service.Repository{}
	// rec := handlers.RecipeRepository{}
	info := service.InitializeDatabase() /* TODO: Change this called inside Connect */
	r.Connect(*info)
	// rec.InitDb()

	e.Logger.Fatal(e.Start(":4444"))

}

func handleGetSample(c echo.Context) error {
	testmode := false
	if testmode {
		sample, err := GetPerson()
		if err != nil {
			return NotFoundError("User")
		}
		return c.JSON(http.StatusOK, sample)
	}

	sample := Person{}

	return c.JSON(http.StatusOK, sample)
}

func GetPerson() (*Person, error) {

	return nil, nil
}
