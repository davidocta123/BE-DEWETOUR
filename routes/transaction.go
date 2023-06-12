package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	TransactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(TransactionRepository)
	e.GET("/transaction", h.GetAllTransaction)
	e.GET("/transactions/:id", h.GetTransByUser)
	e.DELETE("/transaction/:id", h.DeleteTransaction)
	e.POST("/transaction", middleware.Auth(h.CreateTransaction))
	// e.PATCH("/transaction/:id", middleware.Auth(h.UpdateTransaction))
	e.POST("/notification", h.Notification)
}
