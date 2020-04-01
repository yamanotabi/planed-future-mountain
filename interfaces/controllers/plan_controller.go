package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	Path = "/plan"
	Id   = "id"
)

type PlanController struct {
}

func NewPlanController() PlanController {
	return PlanController{}
}

func (p *PlanController) Create(c echo.Context) error {
	return c.JSON(http.StatusOK, "create api")
}

func (p *PlanController) Edit(c echo.Context) error {
	return c.JSON(http.StatusOK, "edit api")
}

func (p *PlanController) Show(c echo.Context) error {
	return c.JSON(http.StatusOK, "sho api")
}

func (p *PlanController) Delete(c echo.Context) error {
	return c.JSON(http.StatusOK, "delete")
}
