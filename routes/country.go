package routes

import (
	"dumbmerch/handlers"
	// "dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repositories"

	"github.com/labstack/echo/v4"
)

func CountryRouter(e *echo.Group) {
	countryRepository := repositories.RepositoryCountry(mysql.DB)
	h := handlers.HandlerCountry(countryRepository)

	e.GET("/country", h.GetAllCountry)
	e.GET("/country/:id", h.GetCountryById)
	e.POST("/country", h.CreateNewCountry)
	e.PATCH("/country/:id", h.UpdateDataCountry)
	e.DELETE("/country/:id", h.DeleteDataCountry)
}
