package infrastructure

import (
	"net/http"

	"github.com/labstack/echo/v4"
	plan "github.com/shiki-tak/planed-future-mountain/interfaces/controllers"
)

func Init() {
	e := echo.New()
	planController := plan.NewPlanController()

	api := e.Group("/api")
	api.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, Mountain")
	})

	api.POST(plan.Path, planController.Create)
	api.PUT(plan.Path+"/:"+plan.Id, planController.Edit)
	api.GET(plan.Path+"/:"+plan.Id, planController.Show)
	api.DELETE(plan.Path+"/:"+plan.Id, planController.Delete)

	e.Logger.Fatal(e.Start(":1323"))
}
