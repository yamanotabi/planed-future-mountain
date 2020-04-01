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

	api.POST(plan.Path, func(c echo.Context) error { return planController.Create(c) })
	api.PUT(plan.Path+"/:"+plan.Id, func(c echo.Context) error { return planController.Edit(c) })
	api.GET(plan.Path+"/:"+plan.Id, func(c echo.Context) error { return planController.Show(c) })
	api.DELETE(plan.Path+"/:"+plan.Id, func(c echo.Context) error { return planController.Delete(c) })

	e.Logger.Fatal(e.Start(":1323"))
}
