package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repositories"

	"github.com/labstack/echo/v4"
)

func TripRouter(e *echo.Group) {
	TripRepository := repositories.RepositoryTrip(mysql.DB)
	h := handlers.HandlerTrip(TripRepository)

	e.GET("/trip", h.GetAllTrip)
	e.GET("/trip/:id", h.GetTripById)
	e.POST("/trip", middleware.UploadFile(h.CreateNewTrip))
	e.PATCH("/trip/:id", middleware.UploadFile(h.UpdateDataTrip))
	e.DELETE("/trip/:id", h.DeleteDataTrip)
}
